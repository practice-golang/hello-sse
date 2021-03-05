package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var values = []string{"10", "30", "50", "70", "90", "100"}

func sendPercent(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	c.Response().WriteHeader(http.StatusOK)

	for _, v := range values {
		// fmt.Fprintf(c.Response().Writer, "data: %s%s\n\n", v, "%")
		// c.Response().Writer.Write([]byte("data: " + v + "%\n\n"))
		c.Response().Write([]byte("data: " + v + "%\n\n"))
		c.Response().Flush()
		time.Sleep(1 * time.Second)
	}

	c.Response().Write([]byte("data: END-OF-STREAM\n\n"))
	c.Response().Flush()

	return nil
}

func main() {
	e := echo.New()

	e.HideBanner = true

	e.Use(
		middleware.CORS(),
		middleware.Recover(),
	)

	e.GET("/percent", sendPercent)
	e.Logger.Fatal(e.Start("127.0.0.1:2918"))
}
