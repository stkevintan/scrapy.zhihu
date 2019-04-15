package main

import (
	"log"

	"./api"
	"./browser"
)

func main() {
	config, err := Parser()
	if err != nil {
		log.Fatal(err)
	}

	cookies := browser.Launch(config.Account)

	api.Init(cookies)

	Start(config.TopicNames, config.MysqlConfig)

	log.Printf("Scrapy stopped\n")
}
