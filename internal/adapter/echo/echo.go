package echo

import (
	"net/http"

	"github.com/davidbrummysw/davidbrummysw-go-orion/internal/core/service"
	"github.com/labstack/echo/v4"
)

type EchoAdapter struct {
}

func NewEchoAdapter() *EchoAdapter {
	return &EchoAdapter{}
}

// Run implements ports.HTTPPort.
func (echoAdpater *EchoAdapter) Run() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/v", func(c echo.Context) error {
		return c.String(http.StatusOK, "0.38")
	})

	var userServiceInterface = service.NewUserService()

	userController := newUserController(userServiceInterface)
	e.GET("/user/test", userController.test)

	e.Logger.Fatal(e.Start(":8080"))

}
