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

func TestHandler_createItem(t *testing.T) {
	type mockBehavior func(s *mock_service.MockItem, userId int, listId int, item structure.Item)

	tests := []struct {
		name               string
		userId             int
		listId             int
		inputBody          string
		setUserId          func(c *gin.Context)
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
			setUserId: func(context *gin.Context) {
				context.Set(userCt, 1)
				context.Params = []gin.Param{
					{
						Key:   "id",
						Value: "1",
					},
				}
			},
			mockBehavior: func(s *mock_service.MockItem, userId int, listId int, item structure.Item) {
				s.EXPECT().CreateItem(userId, listId, item).Return(nil)
			},
			expectedStatusCode: 200,
			expectedResponse:   "{\"Status\":\"success\",\"Message\":\"item created successfully\"}\n",
		},
		{
			name:      "wrong user Id",
			inputBody: `{"Title":"test","Description":"test"}`,
			inputItem: structure.Item{
				Title:       "test",
				Description: "test",
			},
			setUserId:          func(context *gin.Context) {},
			mockBehavior:       func(s *mock_service.MockItem, userId int, listId int, item structure.Item) {},
			expectedStatusCode: 500,
			expectedResponse:   "{\"Status\":\"fail\",\"Message\":\"user id not found\"}\n",
		},
		{
			name:      "wrong list Id",
			userId:    1,
			listId:    1,
			inputBody: `{"Title":"test","Description":"test"}`,
			inputItem: structure.Item{
				Title:       "test",
				Description: "test",
			},
			setUserId: func(context *gin.Context) {
				context.Set(userCt, 1)
				context.Params = []gin.Param{
					{
						Key:   "id",
						Value: "qqq",
					},
				}
			},
			mockBehavior:       func(s *mock_service.MockItem, userId int, listId int, item structure.Item) {},
			expectedStatusCode: 500,
			expectedResponse:   "{\"Status\":\"fail\",\"Message\":\"invalid id param\"}\n",
		},
		{
			name:      "wrong JSON",
			userId:    1,
			listId:    1,
			inputBody: `{"test":"test","Description":"test"}`,
			inputItem: structure.Item{
				Title:       "test",
				Description: "test",
			},
			setUserId: func(context *gin.Context) {
				context.Set(userCt, 1)
				context.Params = []gin.Param{
					{
						Key:   "id",
						Value: "1",
					},
				}
			},
			mockBehavior:       func(s *mock_service.MockItem, userId int, listId int, item structure.Item) {},
			expectedStatusCode: 400,
			expectedResponse:   "{\"Status\":\"fail\",\"Message\":\"Key: 'Item.Title' Error:Field validation for 'Title' failed on the 'required' tag\"}\n",
		},
		{
			name:      "some error",
			userId:    1,
			listId:    1,
			inputBody: `{"Title":"test","Description":"test"}`,
			inputItem: structure.Item{
				Title:       "test",
				Description: "test",
			},
			setUserId: func(context *gin.Context) {
				context.Set(userCt, 1)
				context.Params = []gin.Param{
					{
						Key:   "id",
						Value: "1",
					},
				}
			},
			mockBehavior: func(s *mock_service.MockItem, userId int, listId int, item structure.Item) {
				s.EXPECT().CreateItem(userId, listId, item).Return(errors.New("some error"))
			},
			expectedStatusCode: 500,
			expectedResponse:   "{\"Status\":\"fail\",\"Message\":\"some error\"}\n",
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
			router.POST("/item", testCase.setUserId, handler.createItem)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/item", bytes.NewBufferString(testCase.inputBody))
			router.ServeHTTP(rec, req)

			assert.Equal(t, rec.Code, testCase.expectedStatusCode)
			assert.Equal(t, rec.Body.String(), testCase.expectedResponse)

		})
	}
}
