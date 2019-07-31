package main

import (
	"github.com/garenwen/go-micro-demo/handler"
	"github.com/garenwen/go-micro-demo/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	cs "github.com/garenwen/go-micro-demo/proto/call"
	"github.com/micro/go-plugins/registry/kubernetes"
)

func main() {

	registry := kubernetes.NewRegistry() //a default to using env vars for master API
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.go"),
		micro.Version("latest"),
		micro.Registry(registry),
	)

	// Initialise service
	service.Init()

	// Register Handler
	cs.RegisterCallHandler(service.Server(), new(handler.Call))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.go", service.Server(), new(subscriber.Call))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.go", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
