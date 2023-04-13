package goresponse_test

import (
	"testing"

	goresponse "github.com/douglasdennys45/go-response"
	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	r := goresponse.NewResponse("test", "test", "test", 200, "test", "test")
	response := r.Response()
	assert.Equal(t, response.Data, r.Response().Data)
	assert.Equal(t, response.Error, r.Response().Error)
	assert.Equal(t, response.Message, r.Response().Message)
	assert.Equal(t, response.Path, r.Response().Path)
	assert.Equal(t, response.RequestId, r.Response().RequestId)
	assert.Equal(t, response.Status, r.Response().Status)
}
