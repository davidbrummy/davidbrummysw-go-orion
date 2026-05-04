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
	return controller.getUser(c)
}

func (controller *UserController) getUser(c echo.Context) error {
	log.Println("Start:UserController:getUser")

	ctx := c.Request().Context()

	log.Println("UserController:getUser:requestId", ctx.Value(requestIDKey).(string))
	claims, err := jwtClaims(c)
	if err != nil {
		return err
	}
	log.Println("UserController:getUser:subject", claims["sub"])

	user := controller.userServiceInterface.Test()

	json, _ := json.Marshal(user)
	fmt.Println(string(json))

	log.Println("End:UserController:getUser")

	return c.JSON(http.StatusOK, user)
}
