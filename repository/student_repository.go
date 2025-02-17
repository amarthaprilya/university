package repository

import (
	"database/sql"
	"errors"
	"student/models"
)

type StudentRepository interface {
	Register(student *models.Student) (*models.Student, error)
	GetStudentById(studentId int) (*models.Student, error)
	FindStudentByEmail(email string) (*models.Student, error)
}

type studentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *studentRepository {
	return &studentRepository{db}
}

func (r *studentRepository) Register(student *models.Student) (*models.Student, error) {
	query := `
		INSERT INTO students (first_name, last_name, email, password, address, date_of_birth, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
		RETURNING student_id`
	err := r.db.QueryRow(query, student.FirstName, student.LastName, student.Email, student.Password, student.Address, student.DateOfBirth).
		Scan(&student.StudentId)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (r *studentRepository) GetStudentById(studentId int) (*models.Student, error) {
	var student models.Student
	query := `SELECT student_id, first_name, last_name, email, password, address, date_of_birth FROM students WHERE student_id = $1`
	err := r.db.QueryRow(query, studentId).Scan(
		&student.StudentId, &student.FirstName, &student.LastName,
		&student.Email, &student.Password, &student.Address, &student.DateOfBirth,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("student not found")
		}
		return nil, err
	}
	return &student, nil
}

func (r *studentRepository) FindStudentByEmail(email string) (*models.Student, error) {
	var student models.Student
	query := `SELECT student_id, first_name, last_name, email, password, address, date_of_birth FROM students WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(
		&student.StudentId, &student.FirstName, &student.LastName,
		&student.Email, &student.Password, &student.Address, &student.DateOfBirth,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("student not found")
		}
		return nil, err
	}
	return &student, nil
}
