package example

import "context"

type Handler interface {
	ExampleDomain(ctx context.Context) (string, error)
}

func NewHandler(service Service) Handler {
	return &HandlerImpl{Service: service}
}

type HandlerImpl struct {
	Service Service
}

func (h *HandlerImpl) ExampleDomain(ctx context.Context) (string, error) {
	// Binding Req Body
	// var req ExampleDomainReq

	// Call Service
	// res, err := h.Service.ExampleDomain(req)

	// Return Response
	return h.Service.Process("cid")
}
