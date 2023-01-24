package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/rs/zerolog"
	"github.com/wspectra/ToDoApp/internal/service"
	mock_service "github.com/wspectra/ToDoApp/internal/service/mocks"
	"net/http/httptest"
	"testing"
)

func TestHandler_userIdentity(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, token string)

	tests := []struct {
		name               string
		headerName         string
		headerValue        string
		token              string
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name:        "ok",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			expectedStatusCode: 200,
			expectedResponse:   "1",
		},
		{
			name:               "empty auth header",
			headerName:         "Authorization",
			headerValue:        "",
			token:              "",
			mockBehavior:       func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode: 401,
			expectedResponse:   "{\"Status\":\"fail\",\"Message\":\"empty auth header\"}\n",
		},
		{
			name:               "invalid auth header",
			headerName:         "Authorization",
			headerValue:        "123",
			token:              "token",
			mockBehavior:       func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode: 401,
			expectedResponse:   "{\"Status\":\"fail\",\"Message\":\"invalid auth header\"}\n",
		},
		{
			name:               "invalid auth header",
			headerName:         "Authorization",
			headerValue:        "123 123",
			token:              "token",
			mockBehavior:       func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode: 401,
			expectedResponse:   "{\"Status\":\"fail\",\"Message\":\"invalid auth header\"}\n",
		},
		{
			name:               "invalid auth header",
			headerName:         "Authorization",
			headerValue:        "Bearer ",
			mockBehavior:       func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode: 401,
			expectedResponse:   "{\"Status\":\"fail\",\"Message\":\"token is empty\"}\n",
		},
		{
			name:        "token expired",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(-1, errors.New("token expired"))
			},
			expectedStatusCode: 401,
			expectedResponse:   "{\"Status\":\"fail\",\"Message\":\"token expired\"}\n",
		},
	}

	zerolog.SetGlobalLevel(zerolog.Disabled)
	gin.SetMode(gin.ReleaseMode)

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.token)

			serv := &service.Service{Authorization: auth}
			handler := NewHandler(serv)

			router := gin.New()
			router.POST("/protected", handler.userIdentity, func(context *gin.Context) {
				id, _ := context.Get(userCt)
				context.String(200, fmt.Sprintf("%d", id.(int)))
			})

			rec := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/protected", nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)
			router.ServeHTTP(rec, req)

			assert.Equal(t, rec.Code, testCase.expectedStatusCode)
			assert.Equal(t, rec.Body.String(), testCase.expectedResponse)

		})
	}
}
