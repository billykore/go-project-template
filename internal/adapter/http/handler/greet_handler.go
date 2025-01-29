package handler

import (
	"github.com/billykore/go-service-tmpl/internal/core"
	"github.com/billykore/go-service-tmpl/internal/dto"
	"github.com/billykore/go-service-tmpl/pkg/entity"
	"github.com/billykore/go-service-tmpl/pkg/validation"
	"github.com/labstack/echo/v4"
)

type GreetHandler struct {
	va  *validation.Validator
	svc *core.GreetService
}

func NewGreetHandler(va *validation.Validator, svc *core.GreetService) *GreetHandler {
	return &GreetHandler{
		va:  va,
		svc: svc,
	}
}

// History swaggo annotation.
//
//	@Summary		Get history
//	@Description	Get all message history
//	@Tags			greet
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	entity.Response
//	@Failure		400	{object}	entity.Response
//	@Failure		404	{object}	entity.Response
//	@Failure		500	{object}	entity.Response
//	@Router			/hello [get]
func (h *GreetHandler) History(ctx echo.Context) error {
	resp, err := h.svc.History(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(resp))
}

// SayHello swaggo annotation.
//
//	@Summary		Say hello
//	@Description	Say hello to the name
//	@Tags			greet
//	@Accept			json
//	@Produce		json
//	@Param			HelloRequest	body		greet.HelloRequest	true	"Say Hello Request"
//	@Success		200				{object}	entity.Response
//	@Failure		400				{object}	entity.Response
//	@Failure		404				{object}	entity.Response
//	@Failure		500				{object}	entity.Response
//	@Router			/hello [post]
func (h *GreetHandler) SayHello(ctx echo.Context) error {
	var req dto.HelloRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	err = h.va.Validate(req)
	if err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	resp, err := h.svc.SayHello(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(resp))
}
