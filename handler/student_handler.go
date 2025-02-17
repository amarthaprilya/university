package handler

import (
	"net/http"
	"student/auth"
	"student/helper"
	"student/models"
	"student/service"

	"github.com/labstack/echo/v4"
)

type studentHandler struct {
	studentService service.ServiceStudent
	authService    auth.UserAuthService
}

func NewStudentHandler(studentService service.ServiceStudent, authService auth.UserAuthService) *studentHandler {
	return &studentHandler{studentService, authService}
}

// RegisterStudent godoc
// @Summary Register a new student
// @Description Create a new student account
// @Tags Students
// @Accept json
// @Produce json
// @Param student body models.StudentParam true "Student Data"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 422 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /students/register [post]
func (h *studentHandler) RegisterStudent(c echo.Context) error {
	var input models.StudentParam

	if err := c.Bind(&input); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIresponse(http.StatusUnprocessableEntity, errors)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	isEmailAvailable, err := h.studentService.IsEmailAvailable(input.Email)
	if err != nil {
		response := helper.APIresponse(http.StatusInternalServerError, "Internal Server Error")
		return c.JSON(http.StatusInternalServerError, response)
	}
	if !isEmailAvailable {
		response := helper.APIresponse(http.StatusConflict, "Email already in use")
		return c.JSON(http.StatusConflict, response)
	}

	new, err := h.studentService.Register(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, "Failed to register student")
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := helper.APIresponse(http.StatusOK, new)
	return c.JSON(http.StatusOK, response)
}

// LoginStudent godoc
// @Summary Login a student
// @Description Authenticate student and generate token
// @Tags Students
// @Accept json
// @Produce json
// @Param student body models.StudentLoginParam true "Student Login Data"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /students/login [post]
func (h *studentHandler) LoginStudent(c echo.Context) error {
	var input models.StudentLoginParam

	if err := c.Bind(&input); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIresponse(http.StatusUnprocessableEntity, errors)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	student, err := h.studentService.Login(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnauthorized, err.Error())
		return c.JSON(http.StatusUnauthorized, response)
	}

	token, err := h.authService.GenerateToken(student.StudentId)
	if err != nil {
		response := helper.APIresponse(http.StatusInternalServerError, "Failed to generate token")
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIresponse(http.StatusOK, token)
	return c.JSON(http.StatusOK, response)
}

// GetStudentByToken godoc
// @Summary Get student by token
// @Description Retrieve student data based on the authentication token
// @Tags Students
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Router /students/me [get]
func (h *studentHandler) GetStudentByToken(c echo.Context) error {
	user, ok := c.Get("currentUser").(*models.Student)
	if !ok {
		response := helper.APIresponse(http.StatusUnauthorized, "Unauthorized")
		return c.JSON(http.StatusUnauthorized, response)
	}

	response := helper.APIresponse(http.StatusOK, user)
	return c.JSON(http.StatusOK, response)
}
