package service

import (
	"log"

	"github.com/davidbrummysw/davidbrummysw-go-orion/internal/core/auth"
)

// Adapter implements the DbPort interface
type LoginService struct {
}

func NewLoginService() *LoginService {
	return &LoginService{}
}

func (service *LoginService) Login(userName string, password string) string {
	log.Println("LoginService:login:userName", userName)
	log.Println("LoginService:login:password", password)

	token, err := auth.NewJWTToken(userName)
	if err != nil {
		log.Println("LoginService:login:tokenError", err)
		return ""
	}

	return token
}
