package service

import "log"

// Adapter implements the DbPort interface
type LoginService struct {
}

func NewLoginService() *LoginService {
	return &LoginService{}
}

func (service *LoginService) Login(userName string, password string) string {
	log.Println("LoginService:login:userName", userName)
	log.Println("LoginService:login:password", password)
	return "token123"
}
