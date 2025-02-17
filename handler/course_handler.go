package handler

import (
	"net/http"
	"student/helper"
	"student/service"

	"github.com/labstack/echo/v4"
)

type CourseHandler struct {
	service service.CourseService
}

func NewCourseHandler(service service.CourseService) *CourseHandler {
	return &CourseHandler{service}
}

// GetAllCourses godoc
// @Summary Get all courses
// @Description Retrieves a list of all available courses
// @Tags Courses
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /courses [get]
func (h *CourseHandler) GetAllCourses(c echo.Context) error {
	courses, err := h.service.GetAllCourses()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	response := helper.APIresponse(http.StatusOK, courses)
	return c.JSON(http.StatusOK, response)
}
