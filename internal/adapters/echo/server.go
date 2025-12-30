package echo

import (
	"net/http"

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

	e.Logger.Fatal(e.Start(":8080"))

}
