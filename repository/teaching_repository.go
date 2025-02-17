package repository

import (
	"database/sql"
	"fmt"
	"student/models"
)

type TeachingRepository interface {
	GetAllTeaching() ([]*models.Teachings, error)
}

type teachingRepository struct {
	db *sql.DB
}

func NewTeachingRepository(db *sql.DB) *teachingRepository {
	return &teachingRepository{db}
}

func (r *teachingRepository) GetAllTeaching() ([]*models.Teachings, error) {
	query := `
		SELECT t.teaching_id, t.professor_id, t.course_id, t.created_at, t.updated_at,
			p.professor_id, p.first_name, p.last_name, p.email, p.password, p.address, p.created_at, p.updated_at,
			c.course_id, c.name, c.description, c.credits, c.department_id, c.created_at, c.updated_at
		FROM teachings t
		JOIN professors p ON t.professor_id = p.professor_id
		JOIN courses c ON t.course_id = c.course_id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching teachings: %v", err)
	}
	defer rows.Close()

	var teachings []*models.Teachings
	for rows.Next() {
		var teaching models.Teachings
		var professor models.Professor
		var course models.Course

		err := rows.Scan(&teaching.TeachingId, &teaching.ProfessorId, &teaching.CourseId, &teaching.CreatedAt, &teaching.UpdatedAt,
			&professor.ProfessorId, &professor.FirstName, &professor.LastName, &professor.Email, &professor.Password, &professor.Address,
			&professor.CreatedAt, &professor.UpdatedAt,
			&course.CourseID, &course.Name, &course.Description, &course.Credits, &course.DepartmentId, &course.CreatedAt, &course.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning teaching row: %v", err)
		}

		teaching.Professors = professor
		teaching.Courses = course

		teachings = append(teachings, &teaching)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over teachings rows: %v", err)
	}

	return teachings, nil
}
