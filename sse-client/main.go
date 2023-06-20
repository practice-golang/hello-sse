package main // import "sse-client"

import (
	"fmt"
	"net/http"
	"strings"
)

type Client struct{}

func (c *Client) Connect(address string) error {
	req, err := http.NewRequest("POST", address, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Connection", "keep-alive")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	for {
		bytes := make([]byte, 4096)
		_, err := resp.Body.Read(bytes)
		if err != nil {
			return err
		}

		data := strings.Trim(string(bytes), "\x00")
		if data[len(data)-2:] == "\n\n" {
			data = data[:len(data)-2]
		}

		fmt.Printf("Received message: \n%s\n", data)
	}
}

func main() {
	c := new(Client)
	c.Connect("http://localhost:2918/percent")
}
