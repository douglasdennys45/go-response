package gopostgres

import (
	"time"
)

type ResponseInterface interface {
	Response() response
}

type response struct {
	Data      interface{} `json:"data,omitempty"`
	RequestId string      `json:"requestId,omitempty"`
	Message   string      `json:"message,omitempty"`
	Status    int         `json:"status,omitempty"`
	Error     string      `json:"error,omitempty"`
	Path      string      `json:"path,omitempty"`
	Timestamp time.Time   `json:"timestamp,omitempty"`
}

func NewResponse(Data interface{}, RequestId string, Message string, Status int, Error string, Path string, Timestamp time.Time) ResponseInterface {
	return &response{Data, RequestId, Message, Status, Error, Path, Timestamp}
}

func (r *response) Response() response {
	return response{
		Data:      r.Data,
		RequestId: r.RequestId,
		Message:   r.Message,
		Status:    r.Status,
		Error:     r.Error,
		Path:      r.Path,
		Timestamp: r.Timestamp,
	}
}
