package main

import (
	"log"
	"os"

	"fmt"

	"github.com/akamensky/argparse"
	"github.com/converged-computing/distributed-fractal/pkg/types"

	"github.com/converged-computing/distributed-fractal/pkg/core"
)

var (
	Header = `
	┏        ┓
	╋┏┓┏┓┏╋┏┓┃
	┛┛ ┗┻┗┗┗┻┗			  
`
)

func RunVersion() {
	fmt.Printf("🌀️ distributed-fractal version %s\n", types.Version)
}

func main() {

	parser := argparse.NewParser("fractal", "Distributed fractal generator")
	versionCmd := parser.NewCommand("version", "See the version of rainbow")
	leaderCmd := parser.NewCommand("leader", "Start the leader")
	workerCmd := parser.NewCommand("worker", "Start a worker")

	// Shared values
	host := parser.String("", "host", &argparse.Options{Default: "localhost:50051", Help: "Leader address (host:port)"})

	// Leader arguments (for image generation)
	colorStep := leaderCmd.Int("", "step", &argparse.Options{Default: 6000, Help: "Color smooth step (greater than iteration count, defaults to 6000)"})
	width := leaderCmd.Int("", "width", &argparse.Options{Default: 1024, Help: "Image width"})
	height := leaderCmd.Int("", "height", &argparse.Options{Default: 768, Help: "Image height"})
	xPos := leaderCmd.Float("x", "xpos", &argparse.Options{Default: -0.00275, Help: "Position on the real axis, x"})
	yPos := leaderCmd.Float("y", "ypos", &argparse.Options{Default: 0.78912, Help: "Position on the imaginary axis, y"})
	escapeRadius := leaderCmd.Float("", "escape-radius", &argparse.Options{Default: 0.125689, Help: "Escape radius"})
	smoothness := leaderCmd.Int("", "smoothness", &argparse.Options{Default: 8, Help: "Rendered mandelbrot smoothness, higher is more detailed"})
	iters := leaderCmd.Int("", "iters", &argparse.Options{Default: 800, Help: "Iteration count"})
	palette := leaderCmd.String("", "palette", &argparse.Options{Default: "Hippi", Help: "Color palette, Hippi | Plan9 | AfternoonBlue | SummerBeach | Biochimist | Fiesta"})
	outfile := leaderCmd.String("", "outfile", &argparse.Options{Default: "mandelbrot.png", Help: "Output png file"})

	fmt.Println(*host)

	// Now parse the arguments
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(Header)
		fmt.Println(parser.Usage(err))
		return
	}

	// TODO add error handling here
	if workerCmd.Happened() {
		worker := core.GetWorkerNode(*host)
		err := worker.Start()
		if err != nil {
			log.Fatalf("Issue with starting worker: %s", err)
		}

	} else if leaderCmd.Happened() {
		leader, err := core.GetLeader(*host, *colorStep, *width, *height, *xPos, *yPos, *escapeRadius, *smoothness, *iters, *palette, *outfile)
		if err != nil {
			log.Fatalf("Issue with getting leader: %s", err)
		}
		leader.Start()
	} else if versionCmd.Happened() {
		RunVersion()
	} else {
		fmt.Println(Header)
		fmt.Println(parser.Usage(nil))
	}
}
