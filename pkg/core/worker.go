package core

import (
	"context"
	"fmt"

	"github.com/converged-computing/distributed-fractal/pkg/algorithm"
	pb "github.com/converged-computing/distributed-fractal/pkg/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type WorkerNode struct {
	conn       *grpc.ClientConn
	client     pb.NodeServiceClient
	leaderHost string
}

func (n *WorkerNode) Init() (err error) {
	fmt.Println(n.leaderHost)
	n.conn, err = grpc.Dial(n.leaderHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	n.client = pb.NewNodeServiceClient(n.conn)
	return nil
}

func (n *WorkerNode) Start() error {

	fmt.Println("worker node started")

	// assign task -
	stream, err := n.client.AssignTask(context.Background(), &pb.Request{Action: "Started"})
	if err != nil {
		return err
	}
	for {

		// receive work from leader - a pixel of mandelbrot to parse
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		fmt.Printf("Received work: %s\n", req)
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

var workerNode *WorkerNode

func GetWorkerNode(host string) *WorkerNode {

	if workerNode == nil {
		workerNode = &WorkerNode{leaderHost: host}
		if err := workerNode.Init(); err != nil {
			panic(err)
		}
	}

	return workerNode
}
