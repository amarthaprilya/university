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
// @host localhost:8080/
// @securitydefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
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
	e.POST("students/register", studentHandler.RegisterStudent)
	e.POST("students/login", studentHandler.LoginStudent)
	e.GET("students/me", studentHandler.GetStudentByToken, mdw.AuthMiddleware(authService, studentService))

	courseRepo := repository.NewCourseRepository(database)
	courseService := service.NewCourseService(courseRepo)
	courseHandler := handler.NewCourseHandler(courseService)

	e.GET("/courses", courseHandler.GetAllCourses)

	enrollmentRepo := repository.NewEnrollmentRepository(database)
	enrollmentService := service.NewEnrollmentService(enrollmentRepo)
	enrollmentHandler := handler.NewEnrollmentHandler(enrollmentService)

	e.POST("/enrollments", enrollmentHandler.CreateEnrollmentHandler)
	e.DELETE("/enrollments/:id", enrollmentHandler.DeleteEnrollmentHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := e.Start(":" + port)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
