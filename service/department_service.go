package service

import (
	"student/models"
	"student/repository"
)

type DepartmentService interface {
	GetAllDepartments() ([]*models.Department, error)
}

type departmentService struct {
	repo repository.DepartmentRepository
}

func NewDepartmentService(repo repository.DepartmentRepository) *departmentService {
	return &departmentService{repo}
}

func (s *departmentService) GetAllDepartments() ([]*models.Department, error) {
	return s.repo.GetAllDepartment()
}
