package service

import (
	"student/models"
	"student/repository"
)

type ProfessorService interface {
	GetAllProfessor() ([]*models.Professor, error)
}

type professorService struct {
	repositoryProfessor repository.ProfessorRepository
}

func NewProfessorService(repositoryProfessor repository.ProfessorRepository) ProfessorService {
	return &professorService{repositoryProfessor}
}

func (s *professorService) GetAllProfessor() ([]*models.Professor, error) {
	get, err := s.repositoryProfessor.GetAllProfessor()
	if err != nil {
		return nil, err
	}
	return get, nil
}
