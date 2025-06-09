package handler

import (
	"github.com/billykore/go-service-tmpl/internal/adapter/http/dto"
	"github.com/billykore/go-service-tmpl/internal/adapter/http/response"
	"github.com/billykore/go-service-tmpl/internal/pkg/validation"
	"github.com/billykore/go-service-tmpl/internal/service/example"
	"github.com/labstack/echo/v4"
)

type ExampleHandler struct {
	va  *validation.Validator
	svc *example.Service
}

func NewExampleHandler(va *validation.Validator, svc *example.Service) *ExampleHandler {
	return &ExampleHandler{
		va:  va,
		svc: svc,
	}
}

// Get swaggo annotation.
//
//	@Summary		Get example
//	@Description	Get example by ID.
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Param			HelloRequest	body		dto.GetRequest	true	"Get Request"
//	@Success		200				{object}	response.Response
//	@Failure		400				{object}	response.Response
//	@Failure		404				{object}	response.Response
//	@Failure		500				{object}	response.Response
//	@Router			/example [get]
func (h *ExampleHandler) Get(ctx echo.Context) error {
	var req dto.GetRequest
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(response.BadRequest(err))
	}
	err = h.va.Validate(req)
	if err != nil {
		return ctx.JSON(response.BadRequest(err))
	}
	e, err := h.svc.Get(req.ID)
	if err != nil {
		return ctx.JSON(response.Error(err))
	}
	resp := dto.NewGetResponse(e)
	return ctx.JSON(response.Success(resp))
}
