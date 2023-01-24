package handler

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/rs/zerolog"
	"github.com/wspectra/ToDoApp/internal/service"
	mock_service "github.com/wspectra/ToDoApp/internal/service/mocks"
	"github.com/wspectra/ToDoApp/internal/structure"
	"net/http/httptest"
	"testing"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, user structure.User)

	tests := []struct {
		name               string
		inputBody          string
		inputUser          structure.User
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name:      "ok",
			inputBody: `{"name":"name","password":"123","username":"qqq"}`,
			inputUser: structure.User{
				Name:     "name",
				Password: "123",
				Username: "qqq",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user structure.User) {
				s.EXPECT().AddNewUser(user).Return(nil)
			},
			expectedStatusCode: 200,
			expectedResponse:   "{\"Status\":\"success\",\"Message\":\"new user created successfully\"}\n",
		},
		{
			name:      "invalid JSON",
			inputBody: `{"name":"name","username":"qqq"}`,
			inputUser: structure.User{},
			mockBehavior: func(s *mock_service.MockAuthorization, user structure.User) {
			},
			expectedStatusCode: 400,
			expectedResponse:   "{\"Status\":\"fail\",\"Message\":\"Key: 'User.Password' Error:Field validation for 'Password' failed on the 'required' tag\"}\n",
		},
		{
			name:      "internal error",
			inputBody: `{"name":"name","password":"123","username":"qqq"}`,
			inputUser: structure.User{
				Name:     "name",
				Password: "123",
				Username: "qqq",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user structure.User) {
				s.EXPECT().AddNewUser(user).Return(errors.New("some error"))
			},
			expectedStatusCode: 500,
			expectedResponse:   "{\"Status\":\"fail\",\"Message\":\"some error\"}\n",
		},
	}

	zerolog.SetGlobalLevel(zerolog.Disabled)

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			gin.SetMode(gin.ReleaseMode)
			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			serv := &service.Service{Authorization: auth}
			handler := NewHandler(serv)

			router := gin.New()
			router.POST("/sign-up", handler.signUp)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(testCase.inputBody))
			router.ServeHTTP(rec, req)

			assert.Equal(t, rec.Code, testCase.expectedStatusCode)
			assert.Equal(t, rec.Body.String(), testCase.expectedResponse)

		})
	}
}

func TestHandler_signIn(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, user structure.SignInUser)

	tests := []struct {
		name               string
		inputBody          string
		inputUser          structure.SignInUser
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name:      "ok",
			inputBody: `{"password":"123","username":"qqq"}`,
			inputUser: structure.SignInUser{
				Password: "123",
				Username: "qqq",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user structure.SignInUser) {
				s.EXPECT().GetToken(user).Return("token", nil)
			},
			expectedStatusCode: 200,
			expectedResponse:   "{\"Status\":\"success\",\"Message\":\"token\"}\n",
		},
		{
			name:      "invalid JSON",
			inputBody: `{"name":"name","username":"qqq"}`,
			inputUser: structure.SignInUser{},
			mockBehavior: func(s *mock_service.MockAuthorization, user structure.SignInUser) {
			},
			expectedStatusCode: 400,
			expectedResponse:   "{\"Status\":\"fail\",\"Message\":\"Key: 'SignInUser.Password' Error:Field validation for 'Password' failed on the 'required' tag\"}\n",
		},
		{
			name:      "internal error",
			inputBody: `{"password":"123","username":"qqq"}`,
			inputUser: structure.SignInUser{
				Password: "123",
				Username: "qqq",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user structure.SignInUser) {
				s.EXPECT().GetToken(user).Return("", errors.New("wrong password"))
			},
			expectedStatusCode: 500,
			expectedResponse:   "{\"Status\":\"fail\",\"Message\":\"wrong password\"}\n",
		},
	}

	zerolog.SetGlobalLevel(zerolog.Disabled)

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			gin.SetMode(gin.ReleaseMode)
			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			serv := &service.Service{Authorization: auth}
			handler := NewHandler(serv)

			router := gin.New()
			router.POST("/sign-in", handler.signIn)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-in", bytes.NewBufferString(testCase.inputBody))
			router.ServeHTTP(rec, req)

			assert.Equal(t, rec.Code, testCase.expectedStatusCode)
			assert.Equal(t, rec.Body.String(), testCase.expectedResponse)

		})
	}
}
