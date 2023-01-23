package handler

import (
	"encoding/xml"
	mock_service "github.com/wspectra/ToDoApp/internal/service/mocks"
	"github.com/wspectra/ToDoApp/internal/structure"
	"testing"
)

func TestHandler_signUP(t *testing.T) {
	//Arrange
	type mockBehavior func(s *mock_service.MockAuthorization, user structure.User)

	 tests := []struct {
		name               string
		user               structure.User
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedResponse   string
	}{
		 {
			 name: "ok",
			 user: structure.User{
				 Name: "name",
				 Password: "password",
				Username: "username",
			 },
			 mockBehavior: func(s *mock_service.MockAuthorization, user structure.User) {
				 s.EXPECT().AddNewUser(user).Return(nil)
			 },
			 expectedStatusCode: 200,
			 expectedResponse:
		 },
	 }
}