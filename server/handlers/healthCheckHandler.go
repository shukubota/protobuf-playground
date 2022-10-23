package handlers

import (
	"context"
	protobuf "github.com/shukubota/protobuf-playground/grpc/v1"
)

type healthCheckHandler struct {
}

func (h healthCheckHandler) GetStatus(ctx context.Context, request *protobuf.GetStatusRequest) (*protobuf.GetStatusResponse, error) {
	return &protobuf.GetStatusResponse{
		Status: "OK",
	}, nil
}

func NewHealthCheckHandler() protobuf.HealthCheckServiceServer {
	var h protobuf.HealthCheckServiceServer
	h = healthCheckHandler{}
	return h
}
