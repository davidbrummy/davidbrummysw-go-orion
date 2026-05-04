package echo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	echoframework "github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLoginController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LoginController Suite")
}

type stubLoginService struct {
	token            string
	callCount        int
	receivedUserName string
	receivedPassword string
}

func (service *stubLoginService) Login(userName string, password string) string {
	service.callCount++
	service.receivedUserName = userName
	service.receivedPassword = password
	return service.token
}

var _ = Describe("LoginController", func() {
	var (
		echoInstance *echoframework.Echo
		loginService *stubLoginService
		controller   *LoginController
	)

	BeforeEach(func() {
		echoInstance = echoframework.New()
		loginService = &stubLoginService{token: "jwt-token"}
		controller = newLoginController(loginService)
	})

	It("binds credentials, delegates to the login service, and returns the token", func() {
		request := httptest.NewRequest(
			http.MethodPost,
			"/v1/login",
			strings.NewReader(`{"email":"test@example.com","password":"secret"}`),
		)
		request.Header.Set(echoframework.HeaderContentType, echoframework.MIMEApplicationJSON)
		responseRecorder := httptest.NewRecorder()
		context := echoInstance.NewContext(request, responseRecorder)

		err := controller.login(context)

		Expect(err).NotTo(HaveOccurred())
		Expect(responseRecorder.Code).To(Equal(http.StatusOK))
		Expect(loginService.callCount).To(Equal(1))
		Expect(loginService.receivedUserName).To(Equal("test@example.com"))
		Expect(loginService.receivedPassword).To(Equal("secret"))

		var response StringResponse
		Expect(json.Unmarshal(responseRecorder.Body.Bytes(), &response)).To(Succeed())
		Expect(response.Response).NotTo(BeNil())
		Expect(*response.Response).To(Equal("jwt-token"))
	})

	It("returns an error response when the request body cannot be bound", func() {
		request := httptest.NewRequest(http.MethodPost, "/v1/login", strings.NewReader(`{`))
		request.Header.Set(echoframework.HeaderContentType, echoframework.MIMEApplicationJSON)
		responseRecorder := httptest.NewRecorder()
		context := echoInstance.NewContext(request, responseRecorder)

		err := controller.login(context)

		Expect(err).NotTo(HaveOccurred())
		Expect(responseRecorder.Code).To(Equal(http.StatusOK))
		Expect(strings.TrimSpace(responseRecorder.Body.String())).To(Equal("error"))
		Expect(loginService.callCount).To(Equal(0))
	})
})
