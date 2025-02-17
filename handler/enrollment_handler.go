package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"student/formatter"
	"student/helper"
	"student/models"
	"student/service"

	"github.com/labstack/echo/v4"
)

type EnrollmentHandler struct {
	enrollmentService service.EnrollmentService
}

func NewEnrollmentHandler(enrollmentService service.EnrollmentService) *EnrollmentHandler {
	return &EnrollmentHandler{enrollmentService}
}

// CreateEnrollmentHandler godoc
// @Summary Enroll a student in a course
// @Description Enrolls a student in a specific course
// @Tags Enrollments
// @Accept json
// @Produce json
// @Param enrollment body models.EnrollmentsParam true "Enrollment Data"
// @Success 201 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /enrollments [post]
func (h *EnrollmentHandler) CreateEnrollmentHandler(c echo.Context) error {
	var req models.EnrollmentsParam

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request"})
	}

	fmt.Printf("Parsed Request: %+v\n", req)

	enrollment, err := h.enrollmentService.CreateEnrollment(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	response := helper.APIresponse(http.StatusCreated, enrollment)
	return c.JSON(http.StatusCreated, response)
}

// DeleteEnrollmentHandler godoc
// @Summary Delete a student's enrollment
// @Description Removes a student's enrollment from a course
// @Tags Enrollments
// @Param id path int true "Enrollment ID"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /enrollments/{id} [delete]
func (h *EnrollmentHandler) DeleteEnrollmentHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid enrollment ID"})
	}

	deletedEnrollment, err := h.enrollmentService.DeleteEnrollment(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	response := helper.APIresponse(http.StatusOK, formatter.FormatEnrollmentDelete(deletedEnrollment))
	return c.JSON(http.StatusOK, response)
}
