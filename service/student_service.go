package service

import (
	"errors"
	"fmt"
	"student/models"
	"student/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type ServiceStudent interface {
	Register(student models.StudentParam) (*models.Student, error)
	Login(inputStudent models.StudentLoginParam) (*models.Student, error)
	GetStudentById(studentId int) (*models.Student, error)
	IsEmailAvailable(email string) (bool, error)
}

type serviceStudent struct {
	repositoryStudent repository.StudentRepository
}

func NewStudentService(repositoryStudent repository.StudentRepository) *serviceStudent {
	return &serviceStudent{repositoryStudent}
}

func (s *serviceStudent) Register(student models.StudentParam) (*models.Student, error) {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	students := &models.Student{
		FirstName:   student.FirstName,
		LastName:    student.LastName,
		Email:       student.Email,
		Password:    string(passwordHash),
		Address:     student.Address,
		DateOfBirth: student.DateOfBirth,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	createUser, err := s.repositoryStudent.Register(students)
	if err != nil {
		return createUser, err
	}
	return createUser, nil
}

func (s *serviceStudent) Login(inputStudent models.StudentLoginParam) (*models.Student, error) {
	email := inputStudent.Email
	password := inputStudent.Password

	checkUser, err := s.repositoryStudent.FindStudentByEmail(email)
	if err != nil {
		return checkUser, err
	}
	if checkUser.StudentId == 0 {
		return checkUser, errors.New("user not found that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return checkUser, nil
}

func (s *serviceStudent) GetStudentById(studentID int) (*models.Student, error) {
	student, err := s.repositoryStudent.GetStudentById(studentID)
	if err != nil {
		fmt.Println("Database error while fetching student:", err)
		return nil, err
	}

	if student.StudentId == 0 {
		fmt.Println("Student not found with ID:", studentID)
		return nil, errors.New("student not found")
	}

	return student, nil
}

func (s *serviceStudent) IsEmailAvailable(email string) (bool, error) {
	student, err := s.repositoryStudent.FindStudentByEmail(email)
	if err != nil {
		if err.Error() == "student not found" {
			return true, nil
		}
		return false, err
	}
	if student != nil {
		return false, nil
	}
	return true, nil
}
