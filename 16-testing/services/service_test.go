package services

import (
	"testing"
	mocks "testing-demo/mocks/services"
)

func TestMessageProcessor(t *testing.T) {
	mockMessageService := &mocks.MessageService{}
	mockMessageService.Mock.On("Send", "Hello").Return(true)

	processor := MessageProcessor{mockMessageService}
	processor.Process("Hello")

	mockMessageService.AssertExpectations(t)
}
