package handler

import (
	"net/http"
	"student/formatter"
	"student/helper"
	"student/service"

	"github.com/labstack/echo/v4"
)

type TeachingHandler struct {
	serviceTeaching service.TeachingService
}

func NewTeachingHandler(serviceTeaching service.TeachingService) *TeachingHandler {
	return &TeachingHandler{serviceTeaching}
}

// GetAllTeaching godoc
// @Summary Get all teaching
// @Description Retrieves a list of all available teaching
// @Tags Teaching
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /teaching [get]
func (h *TeachingHandler) GetAllTeaching(c echo.Context) error {
	professor, err := h.serviceTeaching.GetAllTeaching()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	response := helper.APIresponse(http.StatusOK, formatter.FormatTeachings(professor))
	return c.JSON(http.StatusOK, response)
}
