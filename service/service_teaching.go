package service

import (
	"student/models"
	"student/repository"
)

type TeachingService interface {
	GetAllTeaching() ([]*models.Teachings, error)
}

type teachingService struct {
	repositoryTeaching repository.TeachingRepository
}

func NewTeachingService(repositoryTeaching repository.TeachingRepository) TeachingService {
	return &teachingService{repositoryTeaching: repositoryTeaching}
}

func (s *teachingService) GetAllTeaching() ([]*models.Teachings, error) {
	get, err := s.repositoryTeaching.GetAllTeaching()
	if err != nil {
		return nil, err
	}
	return get, nil
}
