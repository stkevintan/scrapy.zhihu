package main

import (
	"io/ioutil"
	"log"

	api "./api"
	browser "./browser"
)

func main() {
	config := Parser()
	cookies := browser.Launch(config.Account)
	api.Init(cookies)
	travelTopic := api.NewTopicDescr("19551556", 20, "")
	data, err := api.TopicList(travelTopic)
	if err != nil {
		log.Fatalf("api failed. %s", err.Error())
	}

	err = ioutil.WriteFile("request.json", data, 0644)
	if err != nil {
		log.Fatalf("write file failed. %s", err.Error())
	}
}
