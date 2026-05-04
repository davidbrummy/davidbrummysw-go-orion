package main

import (
	"context"
	"log"

	"github.com/davidbrummysw/davidbrummysw-go-orion/internal/adapter/echo"
	"github.com/davidbrummysw/davidbrummysw-go-orion/internal/adapter/postgres"
	"github.com/davidbrummysw/davidbrummysw-go-orion/internal/core/ports"
	"github.com/davidbrummysw/davidbrummysw-go-orion/internal/core/service"
)

func main() {
	log.Println("Start")

	ctx := context.Background()
	db, err := postgres.NewConnection(ctx, postgres.DatabaseURLFromEnvironment())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepository := postgres.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	loginService := service.NewLoginService()

	var httpAdapter ports.HTTPPort
	httpAdapter = echo.NewEchoAdapter(userService, loginService)
	httpAdapter.Run()
	log.Println("End")
}
