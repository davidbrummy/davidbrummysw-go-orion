package echo

import (
	"net/http"

	"github.com/davidbrummysw/davidbrummysw-go-orion/internal/core/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoAdapter struct {
}

func NewEchoAdapter() *EchoAdapter {
	return &EchoAdapter{}
}

// Run implements ports.HTTPPort.
func (echoAdpater *EchoAdapter) Run() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},                                                                // Allow only your frontend origin
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions}, // Include all necessary methods
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization}, // Include standard and custom headers like Authorization
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/v", func(c echo.Context) error {
		return c.String(http.StatusOK, "0.38")
	})

	var userServiceInterface = service.NewUserService()
	var loginServiceInterface = service.NewLoginService()

	userController := newUserController(userServiceInterface)
	loginController := newLoginController(loginServiceInterface)

	e.GET("/user/test", userController.test)
	e.POST("/v1/login", loginController.login)

	e.Logger.Fatal(e.Start(":8080"))

}
