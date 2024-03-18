package core

import (
	"fmt"

	pb "github.com/converged-computing/distributed-fractal/pkg/api/v1"

	"context"
)

type NodeServiceGrpcServer struct {
	pb.UnimplementedNodeServiceServer
	ResultChannel chan IterationResult
	WorkChannel   chan MandelIteration
}

// ReportStatus reports a worker status
// Assume it's at the start of a run and we set reporting of metrics back to false
func (n NodeServiceGrpcServer) ReportStatus(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	fmt.Printf("received worker status: %s\n", request.Action)
	return &pb.Response{Data: "ok"}, nil
}

// Report a result to the leader
func (n NodeServiceGrpcServer) ReportResult(ctx context.Context, result *pb.WorkResponse) (*pb.Response, error) {

	// I hate that this is the "best way" to do this
	its := make([]int, len(result.It))
	for ix, it := range result.It {
		its[ix] = int(it)
	}
	// Prepare into a result
	res := IterationResult{
		Norm:         result.Norm,
		IndexY:       int(result.Iy),
		ResultIt:     its,
		MaxIteration: result.Iters,
	}
	n.ResultChannel <- res
	return &pb.Response{Data: "ok"}, nil
}

func (n NodeServiceGrpcServer) AssignTask(request *pb.Request, server pb.NodeService_AssignTaskServer) error {
	for {
		select {
		case work := <-n.WorkChannel:

			// Convert the channel Mandelbrot work into the WorkRequest
			req := pb.WorkRequest{
				Xmin:  work.Xmin,
				Ymin:  work.Ymin,
				Xmax:  work.Xmax,
				Ymax:  work.Ymax,
				Iy:    int32(work.IndexY),
				Iters: int32(work.MaxIteration),
				Width: int32(work.Width),
			}
			if err := server.Send(&req); err != nil {
				return err
			}
		}
	}
}

var server *NodeServiceGrpcServer

// GetNodeServiceGrpcServer singleton service
func GetNodeServiceGrpcServer() *NodeServiceGrpcServer {
	if server == nil {
		server = &NodeServiceGrpcServer{
			WorkChannel:   make(chan MandelIteration),
			ResultChannel: make(chan IterationResult),
		}
	}
	return server
}
