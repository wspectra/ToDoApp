package handler

import (
	"bytes"
	"fmt"
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

func TestHandler_createItem(t *testing.T) {
	type mockBehavior func(s *mock_service.MockItem, userId int, listId int, item structure.Item)

	tests := []struct {
		name               string
		userId             int
		listId             int
		inputBody          string
		setUserId          func()
		inputItem          structure.Item
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name:      "ok",
			userId:    1,
			listId:    1,
			inputBody: `{"Title":"test","Description":"test"}`,
			inputItem: structure.Item{
				Title:       "test",
				Description: "test",
			},
			setUserId: func() {

			},
			mockBehavior: func(s *mock_service.MockItem, userId int, listId int, item structure.Item) {
				s.EXPECT().CreateItem(userId, listId, item).Return(nil)
			},
			expectedStatusCode: 200,
			expectedResponse:   "{\"Status\":\"success\",\"Message\":\"item created successfully\"}\n",
		},
		{
			name:      "wrong userId",
			listId:    1,
			inputBody: `{"Title":"test","Description":"test"}`,
			inputItem: structure.Item{
				Title:       "test",
				Description: "test",
			},
			mockBehavior: func(s *mock_service.MockItem, userId int, listId int, item structure.Item) {
				s.EXPECT().CreateItem(userId, listId, item).Return(nil)
			},
			expectedStatusCode: 200,
			expectedResponse:   "{\"Status\":\"success\",\"Message\":\"item created successfully\"}\n",
		},
	}

	zerolog.SetGlobalLevel(zerolog.Disabled)
	gin.SetMode(gin.ReleaseMode)

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			itemMock := mock_service.NewMockItem(c)
			testCase.mockBehavior(itemMock, testCase.userId, testCase.listId, testCase.inputItem)

			serv := &service.Service{Item: itemMock}
			handler := NewHandler(serv)

			router := gin.New()
			router.POST("/item", func(context *gin.Context) {
				context.Set(userCt, testCase.userId)
				context.Params = []gin.Param{
					{
						Key:   "id",
						Value: fmt.Sprintf("%d", testCase.listId),
					},
				}
			}, handler.createItem)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/item", bytes.NewBufferString(testCase.inputBody))
			router.ServeHTTP(rec, req)

			assert.Equal(t, rec.Code, testCase.expectedStatusCode)
			assert.Equal(t, rec.Body.String(), testCase.expectedResponse)

		})
	}
}
