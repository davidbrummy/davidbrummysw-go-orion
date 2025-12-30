package main

import (
	"log"

	"github.com/davidbrummysw/davidbrummysw-go-orion/internal/adapters/echo"
	"github.com/davidbrummysw/davidbrummysw-go-orion/ports"
)

func main() {
	log.Println("Start")
	var httpAdapter ports.HTTPPort
	httpAdapter = echo.NewEchoAdapter()
	httpAdapter.Run()
	log.Println("End")
}
