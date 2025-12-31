package echo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userServiceInterface UserServiceInterface
}

func newUserController(userServiceInterface UserServiceInterface) *UserController {
	return &UserController{userServiceInterface: userServiceInterface}
}

func (controller *UserController) test(c echo.Context) error {
	log.Println("Start:UserController:test")

	user := controller.userServiceInterface.Test()

	json, _ := json.Marshal(user)
	fmt.Println(string(json))

	log.Println("End:UserController:test")

	return c.JSON(http.StatusOK, user)
}
