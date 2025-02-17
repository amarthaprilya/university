package service

import (
	"student/models"
	"student/repository"
)

type CourseService interface {
	GetAllCourses() ([]models.Course, error)
}

type serviceCourse struct {
	repositoryCourse repository.CourseRepository
}

func NewCourseService(repositoryCourse repository.CourseRepository) *serviceCourse {
	return &serviceCourse{repositoryCourse}
}

func (s *serviceCourse) GetAllCourses() ([]models.Course, error) {
	return s.repositoryCourse.GetAllCourses()
}
