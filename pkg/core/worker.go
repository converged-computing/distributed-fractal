package core

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/converged-computing/distributed-fractal/pkg/algorithm"
	pb "github.com/converged-computing/distributed-fractal/pkg/api/v1"
	"github.com/converged-computing/distributed-fractal/pkg/metrics"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type WorkerNode struct {
	conn       *grpc.ClientConn
	client     pb.NodeServiceClient
	leaderHost string
	retries    int
	quiet      bool
	counts     map[string]int32
}

func (n *WorkerNode) Init() (err error) {
	if !n.quiet {
		fmt.Println(n.leaderHost)
	}
	n.conn, err = grpc.Dial(n.leaderHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	n.client = pb.NewNodeServiceClient(n.conn)
	return nil
}

// Connect to the stream with some backoff
func (n *WorkerNode) ConnectStream() (pb.NodeService_AssignTaskClient, error) {

	// Retry 5 times
	retry := n.retries
	sleeptime := 2
	for {
		// assign task - done with retry to allow worker starting after
		stream, err := n.client.AssignTask(context.Background(), &pb.Request{Action: "Started"})

		// If we are successful OR we've run out of retries, return
		if err == nil || retry <= 0 {
			return stream, nil
		}
		fmt.Printf("Issue connecting to %s, will retry in %d seconds.\n", n.leaderHost, sleeptime)
		retry -= 1
		time.Sleep(time.Duration(sleeptime) * time.Second)
		sleeptime *= 2
	}
}

// recordMetrics takes a count of calculations, etc.
func (n *WorkerNode) recordMetrics() {
	_, ok := n.counts["tasks"]
	if !ok {
		n.counts["tasks"] = 0
	}
	n.counts["tasks"] += 1
}

func (n *WorkerNode) Start() error {

	if !n.quiet {
		fmt.Println("worker node started")
	}

	stream, err := n.ConnectStream()
	if err != nil {
		return err
	}
	for {

		// receive work from leader - a pixel of mandelbrot to parse
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		if !n.quiet {
			fmt.Printf("Received work: %s\n", req)
		}
		n.recordMetrics()

		// Do calculations across the width, save vector of norms and it values
		norms := make([]float64, req.Width)
		its := make([]int32, req.Width)

		for ix := 0; ix < int(req.Width); ix++ {
			xRange := req.Xmax - req.Xmin
			yRange := req.Ymax - req.Ymin
			x := req.Xmin + (xRange)*float64(ix)/float64(req.Width-1)
			y := req.Ymin + (yRange)*float64(req.Iy)/float64(req.Width-1)
			norm, it := algorithm.MandelIteration(x, y, int(req.Iters))
			norms[ix] = norm
			its[ix] = int32(it)
		}

		// send response back
		_, err = n.client.ReportResult(context.Background(), &pb.WorkResponse{
			Norm:  norms,
			It:    its,
			Iy:    req.Iy,
			Iters: req.Iters,
		})
		if err != nil {
			return err
		}
	}
}

// Report metrics for the worker
// This is currently not used.
func (n *WorkerNode) reportMetrics() {

	// TODO some trigger here to know they are requested
	// Add the worker hostname to the prefix
	prefix := "WORKER"
	hostname, err := os.Hostname()
	if err == nil {
		prefix = fmt.Sprintf("WORKER %s", hostname)
	}
	for key, value := range n.counts {
		fmt.Printf("METRICS WORKER %s %s: %d\n", prefix, key, value)
	}

	// If we report metrics, do based on hostname
	metrics.ReportMetrics(prefix)
}

var workerNode *WorkerNode

func GetWorkerNode(host string, retries int, quiet bool) *WorkerNode {
	if retries == 0 {
		retries = 10
	}
	if workerNode == nil {
		workerNode = &WorkerNode{
			leaderHost: host,
			retries:    retries,
			quiet:      quiet,
			counts:     map[string]int32{},
		}
		if err := workerNode.Init(); err != nil {
			panic(err)
		}
	}

	return workerNode
}
