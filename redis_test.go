package main

import (
	"fmt"
	"testing"

	"github.com/go-redis/redis"
)

func TestRedist(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		// IdleTimeout: 5 * time.Minute,
		// MaxRetries:  5,
	})
	defer func() {
		_ = client.Close()
	}()
	pong, err := client.Ping().Result()
	fmt.Println("ping result: ", pong, err)

	err = client.Set("key", "value", 0).Err()
	if err != nil {
		t.Fatalf("cannot set key in redis, err: %+v\n", err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	// Output: PONG <nil>

}
