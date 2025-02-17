package repository

import (
	"database/sql"
	"student/models"
)

type CourseRepository interface {
	GetAllCourses() ([]models.Course, error)
}

type courseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) *courseRepository {
	return &courseRepository{db}
}

func (r *courseRepository) GetAllCourses() ([]models.Course, error) {
	query := `
		SELECT 
			c.course_id, c.name, c.description, c.credits, c.department_id, c.created_at, c.updated_at,
			d.department_id, d.name, d.description, d.created_at, d.updated_at
		FROM courses c
		JOIN departments d ON c.department_id = d.department_id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []models.Course
	for rows.Next() {
		var course models.Course
		var department models.Department

		err := rows.Scan(
			&course.CourseID, &course.Name, &course.Description, &course.Credits, &course.DepartmentId, &course.CreatedAt, &course.UpdatedAt,
			&department.DepartmentID, &department.Name, &department.Description, &department.CreatedAt, &department.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		course.Departments = department
		courses = append(courses, course)
	}

	return courses, nil
}
