package handler 

import (
	service "api-gateway/internal/service"
)

type HandlerST struct {
	Service *service.ServiceRepositoryClient
}

func NewHandler(service *service.ServiceRepositoryClient) *HandlerST {
	return &HandlerST{
		Service: service,
	}
}

