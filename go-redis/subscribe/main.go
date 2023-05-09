package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

type publishMessage struct {
	NormalString string `json:"normal_string"`
	JSONString   string `json:"json_string"`
}

func main() {
	conn := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	pubsub := conn.Subscribe(context.Background(), "test")
	ch := pubsub.Channel()

	for msg := range ch {
		var message publishMessage
		log.Println("msg: ", msg.Payload)
		err := json.Unmarshal([]byte(msg.Payload), &message)
		if err != nil {
			log.Println("err: ", err)
		}

		fmt.Printf("%s\n", message.NormalString)
		fmt.Printf("%#v\n", message.JSONString)
	}
}
