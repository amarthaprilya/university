package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"student/auth"
	"student/helper"
	"student/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// AuthMiddleware untuk validasi token JWT
func AuthMiddleware(authService auth.UserAuthService, studentService service.ServiceStudent) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if !strings.Contains(authHeader, "Bearer") {
				fmt.Println("Authorization header missing or invalid format")
				return c.JSON(http.StatusUnauthorized, helper.APIresponse(http.StatusUnauthorized, "Unauthorized"))
			}

			tokenString := strings.Split(authHeader, " ")[1]

			token, err := authService.ValidasiToken(tokenString)
			if err != nil {
				fmt.Println("Token validation error:", err)
				return c.JSON(http.StatusUnauthorized, helper.APIresponse(http.StatusUnauthorized, "Invalid Token"))
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				fmt.Println("Invalid token claims")
				return c.JSON(http.StatusUnauthorized, helper.APIresponse(http.StatusUnauthorized, "Unauthorized"))
			}

			studentID := int(claims["student_id"].(float64))
			fmt.Println("Extracted studentID:", studentID)

			user, err := studentService.GetStudentById(studentID)
			if err != nil {
				fmt.Println("Error retrieving student by ID:", err)
				return c.JSON(http.StatusUnauthorized, helper.APIresponse(http.StatusUnauthorized, "User Not Found"))
			}

			fmt.Println("Authenticated student:", user)
			c.Set("currentUser", user)
			return next(c)
		}
	}
}
