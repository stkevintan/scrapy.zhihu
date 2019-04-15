package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"./api"
	"./browser"
)

var sigs = make(chan os.Signal, 1)

func init() {
	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		<-sigs
		cancel()
	}()

	config, err := Parser()
	if err != nil {
		log.Fatal(err)
	}

	cookies := browser.Launch(ctx, config.Account)

	api.Init(cookies)

	Start(ctx, config.TopicNames, config.MysqlConfig)

	log.Printf("Scrapy stopped\n")
}
