package api

import (
	"net/http"

	"github.com/codehand/echo-swaggo-example/model"
	"github.com/labstack/echo/v4"
)

// GetExample godoc
// @Summary Get example
// @Description Get example detail
// @Accept  json
// @Produce  json
// @Tags examples
// @Param example_id path int true "example id"
// @Success 200 {object} model.Example
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Failure 500 {object} error
// @Router /example/{example_id} [get]
func GetExample(c echo.Context) error {
	return c.JSON(http.StatusOK, &model.Example{})
}

// PostExample godoc
// @Summary Create example
// @Description Create new example
// @Accept  json
// @Produce  json
// @Tags examples
// @Param input body model.ExampleRequest true "Payload"
// @Success 201 {object} model.Example
// @Failure 404 {object} error
// @Failure 406 {object} error
// @Failure 500 {object} error
// @Router /example [post]
func PostExample(c echo.Context) error {
	var input model.ExampleRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := c.Validate(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusCreated, &model.Example{})
}

// UpdateManufacturer godoc
// @Summary Update example
// @Description Update example
// @Accept  json
// @Produce  json
// @Tags examples
// @Param input body model.ExampleRequest true "Payload"
// @Param example_id path int true "Example id"
// @Success 202 {object} model.Example
// @Failure 404 {object} error
// @Failure 406 {object} error
// @Failure 500 {object} error "Server error"
// @Router /example/{example_id} [put]
func PutExample(c echo.Context) error {
	var input model.ExampleRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := c.Validate(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusAccepted, &model.Example{})
}

// GetExamples godoc
// @Summary List example
// @Description Get list example with pagination
// @Accept  json
// @Produce  json
// @Tags examples
// @Param page query int false "Page number" mininum(0) default(0)
// @Param page_size query int false "Page size" default(20)
// @Param sort query string false "Field sort"
// @Param get_total query bool false "Flag check get total record" default(false)
// @Success 200 {array} model.Example
// @Failure 404 {object} error
// @Failure 500 {object} error
// @Router /examples [get]
func GetExamples(c echo.Context) error {
	return c.JSON(http.StatusOK, []*model.Example{})
}
