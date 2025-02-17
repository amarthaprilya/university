package service

import (
	"errors"
	"fmt"
	"student/models"
	"student/repository"
)

type EnrollmentService interface {
	CreateEnrollment(req models.EnrollmentsParam) (*models.Enrollments, error)
	DeleteEnrollment(id int) (*models.Enrollments, error)
	GetAllEnrollment() ([]*models.Enrollments, error)
}

type serviceEnrollment struct {
	repositoryEnrollment repository.EnrollmentRepository
}

func NewEnrollmentService(repositoryEnrollment repository.EnrollmentRepository) EnrollmentService {
	return &serviceEnrollment{repositoryEnrollment}
}

// CreateEnrollment creates a new enrollment
func (s *serviceEnrollment) CreateEnrollment(req models.EnrollmentsParam) (*models.Enrollments, error) {
	fmt.Printf("Service received request: %+v\n", req)

	// Validate request
	if req.StudentId == 0 || req.CourseId == 0 {
		return nil, errors.New("student ID and course ID are required")
	}

	enrollment := &models.Enrollments{
		StudentId:      req.StudentId,
		CourseId:       req.CourseId,
		EnrollmentDate: req.EnrollmentDate,
	}

	createdEnrollment, err := s.repositoryEnrollment.Create(enrollment)
	if err != nil {
		return nil, err
	}

	return createdEnrollment, nil
}

// DeleteEnrollment deletes an enrollment by ID
func (s *serviceEnrollment) DeleteEnrollment(id int) (*models.Enrollments, error) {
	if id == 0 {
		return nil, errors.New("invalid enrollment ID")
	}

	deletedEnrollment, err := s.repositoryEnrollment.Delete(id)
	if err != nil {
		return nil, err
	}

	return deletedEnrollment, nil
}

func (s *serviceEnrollment) GetAllEnrollment() ([]*models.Enrollments, error) {
	get, err := s.repositoryEnrollment.GetAll()
	if err != nil {
		return nil, err
	}
	return get, nil
}
