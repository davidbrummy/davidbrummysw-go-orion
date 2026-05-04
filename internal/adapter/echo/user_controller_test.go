package echo

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/davidbrummysw/davidbrummysw-go-orion/internal/core/auth"
	"github.com/davidbrummysw/davidbrummysw-go-orion/internal/core/domain"
	echoframework "github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type stubUserService struct {
	callCount int
	err       error
}

func (service *stubUserService) Test(ctx context.Context) (*domain.User, error) {
	service.callCount++
	if service.err != nil {
		return nil, service.err
	}

	id := uint64(1)
	uuid := "user-uuid"
	userName := "Test User"

	return domain.NewUserFill(&id, &uuid, &userName), nil
}

var _ = Describe("UserController", func() {
	var (
		echoInstance *echoframework.Echo
		userService  *stubUserService
	)

	BeforeEach(func() {
		echoInstance = echoframework.New()
		userService = &stubUserService{}
		userController := newUserController(userService)

		echoInstance.GET(
			"/v1/user",
			userController.getUser,
			jwtMiddleware(),
			wrapperMiddleware1,
			wrapperMiddleware2,
		)
	})

	It("allows a request with a valid JWT token", func() {
		token, err := auth.NewJWTToken("test@example.com")
		Expect(err).NotTo(HaveOccurred())

		request := httptest.NewRequest(http.MethodGet, "/v1/user", nil)
		request.Header.Set(echoframework.HeaderAuthorization, "Bearer "+token)
		responseRecorder := httptest.NewRecorder()

		echoInstance.ServeHTTP(responseRecorder, request)

		Expect(responseRecorder.Code).To(Equal(http.StatusOK))
		Expect(userService.callCount).To(Equal(1))
	})

	It("rejects a request without a JWT token", func() {
		request := httptest.NewRequest(http.MethodGet, "/v1/user", nil)
		responseRecorder := httptest.NewRecorder()

		echoInstance.ServeHTTP(responseRecorder, request)

		Expect(responseRecorder.Code).To(Equal(http.StatusUnauthorized))
		Expect(userService.callCount).To(Equal(0))
	})

	It("returns an error when the user service fails", func() {
		userService.err = errors.New("database error")
		token, err := auth.NewJWTToken("test@example.com")
		Expect(err).NotTo(HaveOccurred())

		request := httptest.NewRequest(http.MethodGet, "/v1/user", nil)
		request.Header.Set(echoframework.HeaderAuthorization, "Bearer "+token)
		responseRecorder := httptest.NewRecorder()

		echoInstance.ServeHTTP(responseRecorder, request)

		Expect(responseRecorder.Code).To(Equal(http.StatusInternalServerError))
		Expect(userService.callCount).To(Equal(1))
	})
})
