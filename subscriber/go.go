package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	cs "github.com/garenwen/go-micro-demo/proto/call"
)

type Call struct{}

func (e *Call) Handle(ctx context.Context, msg *cs.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *cs.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
