package formatter

import (
	"student/models"
	"time"
)

type EnrollmentDeleteResponse struct {
	Courses        CourseResponse `json:"Courses"`
	EnrollmentDate time.Time      `json:"EnrollmentDate"`
	CreatedAt      time.Time      `json:"CreatedAt"`
	UpdatedAt      time.Time      `json:"UpdatedAt"`
}

type CourseResponse struct {
	CourseID     int    `json:"CourseID"`
	Name         string `json:"Name"`
	Description  string `json:"Description"`
	Credits      string `json:"Credits"`
	DepartmentId int    `json:"DepartmentId"`
}

func FormatEnrollmentDelete(enrollment *models.Enrollments) EnrollmentDeleteResponse {
	return EnrollmentDeleteResponse{
		Courses: CourseResponse{
			CourseID:     enrollment.Courses.CourseID,
			Name:         enrollment.Courses.Name,
			Description:  enrollment.Courses.Description,
			Credits:      enrollment.Courses.Credits,
			DepartmentId: enrollment.Courses.DepartmentId,
		},
		EnrollmentDate: enrollment.EnrollmentDate,
		CreatedAt:      time.Time{}, // Default value "0001-01-01T00:00:00Z"
		UpdatedAt:      time.Time{}, // Default value "0001-01-01T00:00:00Z"
	}
}

type EnrollmentResponse struct {
	StudentId      int       `json:"StudentId"`
	CourseId       int       `json:"CourseId"`
	EnrollmentDate time.Time `json:"EnrollmentDate"`
}

func FormatEnrollmentResponse(enrollments []*models.Enrollments) []EnrollmentResponse {
	var response []EnrollmentResponse
	for _, enrollment := range enrollments {
		response = append(response, EnrollmentResponse{
			StudentId:      enrollment.StudentId,
			CourseId:       enrollment.CourseId,
			EnrollmentDate: enrollment.EnrollmentDate,
		})
	}
	return response
}
