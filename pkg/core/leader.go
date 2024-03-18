package core

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/converged-computing/distributed-fractal/pkg/algorithm"
	pb "github.com/converged-computing/distributed-fractal/pkg/api/v1"
	"github.com/converged-computing/distributed-fractal/pkg/colors"
	"github.com/converged-computing/distributed-fractal/pkg/metrics"
	"google.golang.org/grpc"
)

// Leader node instance type
type Leader struct {
	api     *gin.Engine
	ln      net.Listener
	svr     *grpc.Server
	nodeSvr *NodeServiceGrpcServer

	waitGroup    *sync.WaitGroup
	host         string
	colorStep    int
	width        int
	height       int
	xPos         float64
	yPos         float64
	escapeRadius float64
	smoothness   int
	iters        int
	palette      string
	outfile      string
	image        *image.RGBA
	forceExit    bool
	quiet        bool
	metrics      bool
}

// Port returns the port of the leader
func (n *Leader) Port() string {
	parts := strings.Split(n.host, ":")
	return fmt.Sprintf(":%s", parts[1])
}

// Init initializes the leader
func (n *Leader) Init() (err error) {

	// leader grpc server listener with port as 50051
	n.ln, err = net.Listen("tcp", n.Port())
	if err != nil {
		return err
	}

	// grpc server
	n.svr = grpc.NewServer()

	// node service that stores image
	n.nodeSvr = GetNodeServiceGrpcServer()

	// register node service to grpc server
	pb.RegisterNodeServiceServer(n.svr, n.nodeSvr)

	done := make(chan struct{})

	// Create a new image for the leader to hold
	n.width *= n.smoothness
	n.height *= n.smoothness
	n.image = image.NewRGBA(
		image.Rectangle{Min: image.Point{}, Max: image.Point{X: n.width, Y: n.height}},
	)

	// Generate color palette
	palette := colors.InterpolateColors(&n.palette, n.colorStep)

	// Keep a ticker for the user to watch...
	ticker := time.NewTicker(time.Millisecond * 100)
	ticker.Stop()
	var start time.Time

	// require an api ping (or not, headless)
	if n.quiet {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.Discard
	}
	n.api = gin.Default()
	n.api.POST("/start", func(c *gin.Context) {

		// Start the ticker and timer when image is going
		if !n.quiet {
			ticker.Reset(time.Second * 5)
		}
		start = time.Now()

		if len(palette) > 0 {
			if !n.quiet {
				fmt.Print("Rendering image...")
			}
			n.RenderDistributed(palette, done)
		}
		time.Sleep(time.Second)
		c.AbortWithStatus(http.StatusOK)
	})

	// Keep a count of widths, I'm not sure how else to do this
	counter := 0

	// This go routine will populate the result from data sent by the workers
	go func() {
		for {
			select {

			// Done and exit from function
			case <-done:
				fmt.Println("Work is done.")
				return

			// Show ticker!
			case <-ticker.C:
				if !n.quiet {
					fmt.Print(".")
				}

			// We have a result
			case result := <-n.nodeSvr.ResultChannel:
				counter += 1
				for ix, it := range result.ResultIt {
					norm := result.Norm[ix]
					iteration := float64(n.iters-it) + math.Log(norm)

					if int(math.Abs(iteration)) < len(palette)-1 {
						color1 := palette[int(math.Abs(iteration))]
						color2 := palette[int(math.Abs(iteration))+1]
						compiledColor := algorithm.LinearInterpolation(
							algorithm.RgbaToUint(color1),
							algorithm.RgbaToUint(color2),
							uint32(iteration),
						)
						newColor := algorithm.Uint32ToRgba(compiledColor)
						n.image.Set(ix, result.IndexY, newColor)

					}
				}
				// Bad proxy for "this is the last result"
				if counter == n.height-1 {
					ticker.Stop()
					totalTime := time.Since(start)
					output, err := os.Create(n.outfile)
					png.Encode(output, n.image)
					if n.metrics {
						fmt.Printf("METRICS LEADER time: %v\n", totalTime)
						metrics.ReportMetrics("METRICS LEADER")
					}
					if err != nil {
						fmt.Printf("Warning: error creating image file: %s\n", err)
					}
					fmt.Printf("\n\nMandelbrot set rendered into `%s`\n", n.outfile)
					if n.forceExit {
						panic("Image generation complete, force exited.")
					}
				}
			}
		}
	}()
	return nil
}

// RenderDistributed sends work to the worker channel
func (n *Leader) RenderDistributed(colors []color.RGBA, done chan struct{}) {

	// Multiply to account for smoothness
	n.waitGroup = &sync.WaitGroup{}

	// ratio between height and width - TODO read and understand these lines
	ratio := float64(n.height) / float64(n.width)
	xmin, xmax := n.xPos-n.escapeRadius/2.0, math.Abs(n.xPos+n.escapeRadius/2.0)
	ymin, ymax := n.yPos-n.escapeRadius*ratio/2.0, math.Abs(n.yPos+n.escapeRadius*ratio/2.0)

	for iy := 0; iy < n.height; iy++ {
		n.waitGroup.Add(1)
		go func(iy int) {
			defer n.waitGroup.Done()

			// Each worker does an entire set of Y across X
			work := MandelIteration{
				Xmin:         xmin,
				Ymin:         ymin,
				Xmax:         xmax,
				Ymax:         ymax,
				IndexY:       iy,
				Width:        n.width,
				MaxIteration: n.iters,
			}

			// Send one mandel Iteration to a worker to do
			n.nodeSvr.WorkChannel <- work
		}(iy)
	}
	n.waitGroup.Wait()
}

func (n *Leader) Start() {
	// start grpc server
	go n.svr.Serve(n.ln)

	// start api server
	_ = n.api.Run(":9092")

	// wait for exit
	n.svr.Stop()
}

var leader *Leader

// validatePalette ensures the palette we've chosen is known
func (l *Leader) validatePalette() error {
	for _, palette := range colors.ColorPalettes {
		if palette.Keyword == l.palette {
			return nil
		}
	}
	return fmt.Errorf("palette %s is not recognized", l.palette)
}

// GetLeader returns the leader node instance
func GetLeader(
	host string,
	colorStep int,
	width int,
	height int,
	xPos float64,
	yPos float64,
	escapeRadius float64,
	smoothness int,
	iters int,
	palette string,
	outfile string,
	forceExit, quiet, metrics bool,
) (*Leader, error) {

	// Validate the host - must have a port
	if strings.Count(host, ":") != 1 {
		return nil, fmt.Errorf("the leader hostname must have a port")
	}

	// The color step needs to be greater than or = to iterations
	if colorStep < iters {
		colorStep = iters
	}

	// Validate the palette name is known to us
	if leader == nil {
		leader = &Leader{
			host:         host,
			colorStep:    colorStep,
			width:        width,
			height:       height,
			xPos:         xPos,
			yPos:         yPos,
			escapeRadius: escapeRadius,
			smoothness:   smoothness,
			iters:        iters,
			palette:      palette,
			outfile:      outfile,
			forceExit:    forceExit,
			quiet:        quiet,
			metrics:      metrics,
		}

		// Ensure color palette is okay
		err := leader.validatePalette()
		if err != nil {
			return leader, err
		}

		// initialize node
		if err := leader.Init(); err != nil {
			return nil, err
		}
	}
	return leader, nil
}
