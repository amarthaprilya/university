package main

import (
	"log"
	"os"
	"student/auth"
	"student/db"
	_ "student/docs"
	"student/handler"
	mdw "student/middleware"
	"student/repository"
	"student/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Sweager Service API
// @description Sweager service API in Go using Gin framework
// @host university-51cbe47018ea.herokuapp.com/
// @securitydefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			"Origin", "Accept", "X-Requested-With", "Content-Type",
			"Access-Control-Request-Method", "Access-Control-Request-Headers", "Authorization",
		},
		AllowMethods: []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"},
	}))

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Koneksi ke database
	database := db.ConnectDB()

	// Setup auth service
	secretKey := "SECRET_KEY"
	authService := auth.NewUserAuthService()
	authService.SetSecretKey(secretKey)

	// Setup repository & service
	studentRepository := repository.NewStudentRepository(database)
	studentService := service.NewStudentService(studentRepository)

	// Setup handler
	studentHandler := handler.NewStudentHandler(studentService, authService)

	// Middleware logging & recovery (optional)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route grup untuk user
	e.POST("students/register/", studentHandler.RegisterStudent)
	e.POST("students/login/", studentHandler.LoginStudent)
	e.GET("students/me/", studentHandler.GetStudentByToken, mdw.AuthMiddleware(authService, studentService))

	courseRepo := repository.NewCourseRepository(database)
	courseService := service.NewCourseService(courseRepo)
	courseHandler := handler.NewCourseHandler(courseService)

	e.GET("/courses/", courseHandler.GetAllCourses)

	enrollmentRepo := repository.NewEnrollmentRepository(database)
	enrollmentService := service.NewEnrollmentService(enrollmentRepo)
	enrollmentHandler := handler.NewEnrollmentHandler(enrollmentService)

	e.POST("/enrollments/", enrollmentHandler.CreateEnrollmentHandler)
	e.DELETE("/enrollments/:id/", enrollmentHandler.DeleteEnrollmentHandler)
	e.GET("/enrollments/", enrollmentHandler.GetAllEnrollment)

	ProfessorRepo := repository.NewProfessorRepository(database)
	professorService := service.NewProfessorService(ProfessorRepo)
	professorHandler := handler.NewProfessorHandler(professorService)

	e.GET("/professor/", professorHandler.GetAllProfessor)

	teachingRepo := repository.NewTeachingRepository(database)
	teachingService := service.NewTeachingService(teachingRepo)
	teachingHandler := handler.NewTeachingHandler(teachingService)

	e.GET("/teaching/", teachingHandler.GetAllTeaching)

	departmentRepo := repository.NewDepartmentRepository(database)
	departmentService := service.NewDepartmentService(departmentRepo)
	departmentHandler := handler.NewDepartmentHandler(departmentService)

	e.GET("/department/", departmentHandler.GetAllDepartments)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := e.Start(":" + port)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
