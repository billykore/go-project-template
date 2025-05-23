package handler

import (
	"github.com/billykore/go-service-tmpl/internal/http/dto"
	"github.com/billykore/go-service-tmpl/internal/http/response"
	"github.com/billykore/go-service-tmpl/pkg/service"
	"github.com/billykore/go-service-tmpl/pkg/utils/validation"
	"github.com/labstack/echo/v4"
)

type GreetHandler struct {
	va  *validation.Validator
	svc *service.GreetService
}

func NewGreetHandler(va *validation.Validator, svc *service.GreetService) *GreetHandler {
	return &GreetHandler{
		va:  va,
		svc: svc,
	}
}

// SayHello swaggo annotation.
//
//	@Summary		Say hello
//	@Description	Say hello to the name
//	@Tags			greet
//	@Accept			json
//	@Produce		json
//	@Param			HelloRequest	body		dto.HelloRequest	true	"Say Hello Request"
//	@Success		200				{object}	response.Response
//	@Failure		400				{object}	response.Response
//	@Failure		404				{object}	response.Response
//	@Failure		500				{object}	response.Response
//	@Router			/hello [post]
func (h *GreetHandler) SayHello(ctx echo.Context) error {
	var req dto.HelloRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(response.BadRequest(err))
	}
	err = h.va.Validate(req)
	if err != nil {
		return ctx.JSON(response.BadRequest(err))
	}
	greet, err := h.svc.SayHello(ctx.Request().Context(), req.Name)
	if err != nil {
		return ctx.JSON(response.Error(err))
	}
	resp := dto.NewHelloResponse(greet.Name)
	return ctx.JSON(response.Success(resp))
}
