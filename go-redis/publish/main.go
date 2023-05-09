package main

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"log"
)

type publishMessage struct {
	NormalString string `json:"normal_string"`
	JSONString   string `json:"json_string"`
}

func (p publishMessage) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

const jsonString = "[{\"name\":\"John\",\"age\":30,\"car\":null},{\"name\":\"John\",\"age\":30,\"car\":null}]"

func main() {
	conn := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	message := publishMessage{
		NormalString: "normal string",
		JSONString:   jsonString,
	}

	err := conn.Publish(context.Background(), "test", message).Err()
	if err != nil {
		log.Println("err: ", err)
	}
}
