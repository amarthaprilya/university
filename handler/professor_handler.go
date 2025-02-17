package handler

import (
	"net/http"
	"student/helper"
	"student/service"

	"github.com/labstack/echo/v4"
)

type ProfessorHandler struct {
	serviceProfessor service.ProfessorService
}

func NewProfessorHandler(serviceProfessor service.ProfessorService) *ProfessorHandler {
	return &ProfessorHandler{serviceProfessor}
}

// GetAllProfessor godoc
// @Summary Get all professor
// @Description Retrieves a list of all available professor
// @Tags Professor
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /professor [get]
func (h *ProfessorHandler) GetAllProfessor(c echo.Context) error {
	professor, err := h.serviceProfessor.GetAllProfessor()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	response := helper.APIresponse(http.StatusOK, professor)
	return c.JSON(http.StatusOK, response)
}
