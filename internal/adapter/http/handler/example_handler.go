package handler

import (
	"github.com/billykore/go-service-tmpl/internal/adapter/http/dto"
	"github.com/billykore/go-service-tmpl/internal/adapter/http/response"
	"github.com/billykore/go-service-tmpl/internal/pkg/validation"
	"github.com/billykore/go-service-tmpl/internal/usecase/example"
	"github.com/labstack/echo/v4"
)

type ExampleHandler struct {
	va  *validation.Validator
	svc *example.Usecase
}

func NewExampleHandler(va *validation.Validator, svc *example.Usecase) *ExampleHandler {
	return &ExampleHandler{
		va:  va,
		svc: svc,
	}
}

// GetEntity swaggo annotation.
//
//	@Summary		GetEntity example
//	@Description	GetEntity example by ID.
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Param			GetRequest	body		dto.GetRequest	true	"GetEntity Request"
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	response.Response
//	@Failure		404			{object}	response.Response
//	@Failure		500			{object}	response.Response
//	@Router			/example [get]
func (h *ExampleHandler) GetEntity(ctx echo.Context) error {
	var req dto.GetRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(response.BadRequest(err))
	}
	err = h.va.Validate(req)
	if err != nil {
		return ctx.JSON(response.BadRequest(err))
	}
	e, err := h.svc.GetEntity(req.ID)
	if err != nil {
		return ctx.JSON(response.Error(err))
	}
	resp := dto.NewGetResponse(e)
	return ctx.JSON(response.Success(resp))
}

// SaveEntity swaggo annotation.
//
//	@Summary		SaveEntity example
//	@Description	SaveEntity example.
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Param			SaveRequest	body		dto.SaveRequest	true	"SaveEntity Request"
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	response.Response
//	@Failure		404			{object}	response.Response
//	@Failure		500			{object}	response.Response
//	@Router			/example [post]
func (h *ExampleHandler) SaveEntity(ctx echo.Context) error {
	var req dto.SaveRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(response.BadRequest(err))
	}
	err = h.va.Validate(req)
	if err != nil {
		return ctx.JSON(response.BadRequest(err))
	}
	err = h.svc.SaveEntity(req.ID)
	if err != nil {
		return ctx.JSON(response.Error(err))
	}
	return ctx.JSON(response.Success(nil))

}
