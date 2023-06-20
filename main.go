// https://github.com/gofiber/recipes/blob/master/sse/main.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/valyala/fasthttp"
)

var values = []string{"10", "25", "30", "45", "50", "65", "70", "80", "90", "100"}

func sendPercent(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Transfer-Encoding", "chunked")

	c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		fmt.Println("Begin to send percent")
		for _, v := range values {
			msg := fmt.Sprintf("%s%s", v, "%")
			fmt.Fprintf(w, "data: %s\n\n", msg)
			fmt.Println("Sent ", msg)

			err := w.Flush()
			if err != nil {
				fmt.Printf("Error while flushing: %v. Closing http connection.\n", err)
				break
			}
			time.Sleep(2 * time.Second)
		}
	}))

	fmt.Println("End to send percent")

	return nil
}

func main() {
	config := fiber.Config{
		AppName:               string("SSE Server"),
		DisableStartupMessage: true,
		Prefork:               false,
	}

	app := fiber.New(config)
	app.Use(cors.New(cors.Config{AllowOrigins: "*", AllowHeaders: "Cache-Control", AllowCredentials: true}))

	app.Get("/", func(c *fiber.Ctx) error { return c.SendString("Hello, World 👋!") })
	app.Get("/percent", sendPercent)

	log.Fatal(app.Listen("127.0.0.1:2918"))
}
