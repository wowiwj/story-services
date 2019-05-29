package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/wowiwj/story-services/user-service/handler"
	"github.com/wowiwj/story-services/user-service/subscriber"

	user "github.com/wowiwj/story-services/user-service/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.common.srv.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterUserServiceHandler(service.Server(), new(handler.UserService))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.common.srv.user", service.Server(), new(subscriber.User))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.common.srv.user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
