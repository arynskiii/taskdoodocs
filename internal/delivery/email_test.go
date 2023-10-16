package handler

import (
	"doodocs_task/internal/service"
	mock_service "doodocs_task/internal/service/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestHandler_sendfile(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockEmail)

	type fields struct {
		service *service.Service
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := gomock.NewController(t)
			defer c.Finish()

			emailService := mock_service.NewMockEmail(c)
			testService := &service.Service{
				Email: emailService,
			}
			h := &Handler{
				service: testService,
			}

			h.sendfile(tt.args.c)
		})
	}
}
