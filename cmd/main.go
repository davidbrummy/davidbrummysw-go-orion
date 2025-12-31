package main

import (
	"log"

	"github.com/davidbrummysw/davidbrummysw-go-orion/internal/adapter/echo"
	"github.com/davidbrummysw/davidbrummysw-go-orion/internal/core/ports"
)

func main() {
	log.Println("Start")

	var httpAdapter ports.HTTPPort
	httpAdapter = echo.NewEchoAdapter()
	httpAdapter.Run()
	log.Println("End")
}
