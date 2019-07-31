package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	cs "github.com/garenwen/go-micro-demo/proto/call"
)

type Call struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Call) Call(ctx context.Context, req *cs.Request, rsp *cs.Response) error {
	log.Log("Received Call.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Call) Stream(ctx context.Context, req *cs.StreamingRequest, stream cs.Call_StreamStream) error {
	log.Logf("Received Call.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&cs.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Call) PingPong(ctx context.Context, stream cs.Call_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Callt ping %v", req.Stroke)
		if err := stream.Send(&cs.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
