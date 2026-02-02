package echo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/davidbrummysw/davidbrummysw-go-orion/model"
	"github.com/labstack/echo/v4"
)

type LoginController struct {
	loginServiceInterface LoginServiceInterface
}

func newLoginController(loginServiceInterface LoginServiceInterface) *LoginController {
	return &LoginController{loginServiceInterface: loginServiceInterface}
}

func (controller *LoginController) login(c echo.Context) error {
	log.Println("Start:LoginController:login")

	model := model.NewAuthCredentials()
	if err := c.Bind(model); err != nil {
		log.Println(err)
		return c.String(http.StatusOK, "error")
	}

	jsonLog, _ := json.Marshal(model)
	fmt.Println(string(jsonLog))

	token := controller.loginServiceInterface.Login(*model.Email, *model.Password)

	stringResponse := NewStringResponseFill(&token)

	jsonLog1, _ := json.Marshal(stringResponse)
	fmt.Println(string(jsonLog1))

	log.Println("End:LoginController:login")

	return c.JSON(http.StatusOK, stringResponse)
}
