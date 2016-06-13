package main

import (
	"log"

	"github.com/kazouphq/crawler/srv/handler"
	"github.com/kazouphq/crawler/srv/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/cmd"
)

func main() {
	// Initate cmd flags
	cmd.Init()

	service := micro.NewService(
		micro.Name("go.micro.srv.crawler"),
		micro.Version("latest"),
	)
	//Register handler
	go_micro_srv_crawler.RegisterCrawlerHandler(service.Server(), new(handler.Crawler))
	//Init service
	service.Init()
	//Start Service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
