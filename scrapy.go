package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"sync"

	"./api"
)

var store Store
var wg sync.WaitGroup

type paging struct {
	IsStart bool   `json:"is_start"`
	Totals  int    `json:"totals"`
	IsEnd   bool   `json:"is_end"`
	Next    string `json:"next"`
}

type author struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Url       string `json:"url"`
	UserType  string `json:"user_type"`
	AvatarUrl string `json:"avatar_url"`
}

type target struct {
	ID           int    `json:"id"`
	Url          string `json:"url"`
	Excerpt      string `json:"excerpt"`
	Type         string `json:"type"`
	Author       author `json:"author"`
	Created      int64  `json:"created"`
	Updated      int64  `json:"updated"`
	Title        string `json:"title"`
	CommentCount int    `json:"comment_count"`
	VoteUpCount  int    `json:"voteup_count"`
}

type detail struct {
	Target target `json:"target"`
}

type TopicResult struct {
	Paging  paging   `json:"paging"`
	Content []detail `json:"data"`
}

func handler(chunk []byte) (string, error) {
	var topicResult TopicResult
	err := json.Unmarshal(chunk, &topicResult)
	if err != nil {
		return "", err
	}
	nextURL := topicResult.Paging.Next
	u, err := url.Parse(nextURL)
	if err != nil {
		return "", err
	}

	if topicResult.Paging.IsEnd {
		return "", nil
	}

	afterID := u.Query().Get("after_id")

	fmt.Printf("afterid is %s\n", afterID)

	err = store.SaveTopics(topicResult)

	if err != nil {
		return "", fmt.Errorf("cannot save current topic to dataBase, topic: %v, error: %v", topicResult, err)
	}

	return afterID, nil
}

func do(ctx context.Context, topic string) {
	td := api.NewTopicDescr(topic, 20, "")
	log.Printf("Scrapy with topic:%s\n", topic)
Loop:
	for {
		select {
		case <-ctx.Done():
			log.Printf("Scrap process stopped, with topic: %s\n", topic)
			break Loop
		default:
			chunk, err := td.TopicList()
			if err != nil {
				log.Fatalf("cannot get the topicList with %v, error: %v\n", td, err)
			}
			afterID, err := handler(chunk)
			if err != nil {
				log.Fatalf("cannot store data, error: %v\n", err)
			}

			if afterID == "" {
				log.Printf("topic %s scrapy end.\n", td.ID)
				break Loop
			}

			td.AfterID = afterID
		}
	}
	wg.Done()
}

//Start the scrapy process
func Start(ctx context.Context, topics []string, mysqlConfig MysqlConfig) {

	mysqlConfig.Default()

	err := store.Init(ctx, mysqlConfig)

	if err != nil {
		log.Fatalf("cannot connect to database, error: %v", err)
	}
	for _, topic := range topics {
		wg.Add(1)
		go do(ctx, topic)
	}
	wg.Wait()
	store.Close()
	fmt.Println("done")
}
