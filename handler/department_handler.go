package handler

import (
	"fmt"
	"net/http"
	"student/service"

	"github.com/labstack/echo/v4"
)

type DepartmentHandler struct {
	service service.DepartmentService
}

func NewDepartmentHandler(service service.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{service}
}

// GetAllDepartment godoc
// @Summary Get all department
// @Description Retrieves a list of all available department
// @Tags Department
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /department [get]
func (h *DepartmentHandler) GetAllDepartments(c echo.Context) error {
	departments, err := h.service.GetAllDepartments()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error retrieving departments: %v", err))
	}
	return c.JSON(http.StatusOK, departments)
}
