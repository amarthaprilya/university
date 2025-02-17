package formatter

import (
	"student/models"
	"time"
)

type TeachingResponse struct {
	TeachingId  int               `json:"TeachingId"`
	ProfessorId int               `json:"ProfessorId"`
	Professors  ProfessorResponse `json:"Professors"`
	CourseId    int               `json:"CourseId"`
	Courses     CourseResponse    `json:"Courses"`
	CreatedAt   time.Time         `json:"CreatedAt"`
	UpdatedAt   time.Time         `json:"UpdatedAt"`
}

type ProfessorResponse struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Address   string `json:"Address"`
}

func FormatTeaching(teaching *models.Teachings) TeachingResponse {
	return TeachingResponse{
		TeachingId:  teaching.TeachingId,
		ProfessorId: teaching.ProfessorId,
		Professors: ProfessorResponse{
			FirstName: teaching.Professors.FirstName,
			LastName:  teaching.Professors.LastName,
			Address:   teaching.Professors.Address,
		},
		CourseId: teaching.CourseId,
		Courses: CourseResponse{
			CourseID:     teaching.Courses.CourseID,
			Name:         teaching.Courses.Name,
			Description:  teaching.Courses.Description,
			Credits:      teaching.Courses.Credits,
			DepartmentId: teaching.Courses.DepartmentId,
		},
		CreatedAt: teaching.CreatedAt,
		UpdatedAt: teaching.UpdatedAt,
	}
}

func FormatTeachings(teachings []*models.Teachings) []TeachingResponse {
	var response []TeachingResponse
	for _, teaching := range teachings {
		response = append(response, FormatTeaching(teaching))
	}
	return response
}
