package echo

type LoginServiceInterface interface {
	Login(userName string, password string) string
}
