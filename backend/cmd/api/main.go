package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"backend/app/interface/api"

	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
	return c.String(
		http.StatusOK, "Hello, World! 2",
	)
}

var PORT = os.Getenv("PORT")

func main() {

	parsedPort, err := strconv.Atoi(PORT)
	if err != nil {
		log.Panic(err)
	}

	server := api.NewServer()

	if err := server.Init(); err != nil {
		log.Fatal(err)
	}

	server.Run(parsedPort)
}
