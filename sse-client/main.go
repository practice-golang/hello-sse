package main // import "sse-client"

import (
	"fmt"

	"github.com/r3labs/sse/v2"
)

func main() {
	client := sse.NewClient("http://localhost:2918/percent")

	client.Subscribe("messages", func(msg *sse.Event) {
		fmt.Println(string(msg.Data))
	})
}
