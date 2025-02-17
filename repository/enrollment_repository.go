package repository

import (
	"database/sql"
	"fmt"
	"student/models"
)

type EnrollmentRepository interface {
	Create(input *models.Enrollments) (*models.Enrollments, error)
	Delete(id int) (*models.Enrollments, error)
	GetAll() ([]*models.Enrollments, error)
}

type enrollmentRepository struct {
	db *sql.DB
}

func NewEnrollmentRepository(db *sql.DB) EnrollmentRepository {
	return &enrollmentRepository{db}
}

func (r *enrollmentRepository) Create(input *models.Enrollments) (*models.Enrollments, error) {
	checkQuery := `SELECT COUNT(*) FROM enrollments WHERE student_id = $1 AND course_id = $2`
	var count int
	err := r.db.QueryRow(checkQuery, input.StudentId, input.CourseId).Scan(&count)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, fmt.Errorf("student with ID %d is already enrolled in course with ID %d", input.StudentId, input.CourseId)
	}

	query := `
		INSERT INTO enrollments (student_id, course_id, enrollment_date, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING enrolment_id`
	err = r.db.QueryRow(query, input.StudentId, input.CourseId, input.EnrollmentDate).Scan(&input.EnrolmentId)
	if err != nil {
		return nil, err
	}

	studentQuery := `SELECT student_id, first_name, last_name, email, address, date_of_birth, created_at, updated_at FROM students WHERE student_id = $1`
	err = r.db.QueryRow(studentQuery, input.StudentId).
		Scan(&input.Students.StudentId, &input.Students.FirstName, &input.Students.LastName, &input.Students.Email, &input.Students.Address, &input.Students.DateOfBirth, &input.Students.CreatedAt, &input.Students.UpdatedAt)
	if err != nil {
		return nil, err
	}

	courseQuery := `SELECT course_id, name, description, credits, department_id, created_at, updated_at FROM courses WHERE course_id = $1`
	err = r.db.QueryRow(courseQuery, input.CourseId).
		Scan(&input.Courses.CourseID, &input.Courses.Name, &input.Courses.Description, &input.Courses.Credits, &input.Courses.DepartmentId, &input.Courses.CreatedAt, &input.Courses.UpdatedAt)
	if err != nil {
		return nil, err
	}

	departmentQuery := `SELECT department_id, name, description, created_at, updated_at FROM departments WHERE department_id = $1`
	err = r.db.QueryRow(departmentQuery, input.Courses.DepartmentId).
		Scan(&input.Courses.Departments.DepartmentID, &input.Courses.Departments.Name, &input.Courses.Departments.Description, &input.Courses.Departments.CreatedAt, &input.Courses.Departments.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return input, nil
}

func (r *enrollmentRepository) Delete(id int) (*models.Enrollments, error) {
	var enrollment models.Enrollments
	getQuery := `
		SELECT e.course_id, e.enrollment_date, c.course_id, c.name, c.description, c.credits, c.department_id, c.created_at, c.updated_at
		FROM enrollments e
		JOIN courses c ON e.course_id = c.course_id
		WHERE e.enrolment_id = $1
	`
	err := r.db.QueryRow(getQuery, id).
		Scan(&enrollment.CourseId, &enrollment.EnrollmentDate,
			&enrollment.Courses.CourseID, &enrollment.Courses.Name, &enrollment.Courses.Description, &enrollment.Courses.Credits,
			&enrollment.Courses.DepartmentId, &enrollment.Courses.CreatedAt, &enrollment.Courses.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("enrollment with ID %d not found", id)
		}
		return nil, err
	}

	deleteQuery := `DELETE FROM enrollments WHERE enrolment_id = $1`
	_, err = r.db.Exec(deleteQuery, id)
	if err != nil {
		return nil, err
	}

	enrollment.Courses.Departments = models.Department{}

	return &models.Enrollments{
		Courses:        enrollment.Courses,
		EnrollmentDate: enrollment.EnrollmentDate,
	}, nil
}

func (r *enrollmentRepository) GetAll() ([]*models.Enrollments, error) {
	// Query untuk mengambil semua data enrollment
	query := `
		SELECT 
			e.enrolment_id, e.student_id, e.course_id, e.enrollment_date, 
			s.student_id, s.first_name, s.last_name, s.email, s.address, s.date_of_birth, s.created_at, s.updated_at,
			c.course_id, c.name, c.description, c.credits, c.department_id, c.created_at, c.updated_at,
			d.department_id, d.name, d.description, d.created_at, d.updated_at
		FROM enrollments e
		JOIN students s ON e.student_id = s.student_id
		JOIN courses c ON e.course_id = c.course_id
		JOIN departments d ON c.department_id = d.department_id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enrollments []*models.Enrollments

	for rows.Next() {
		var enrollment models.Enrollments
		var student models.Student
		var course models.Course
		var department models.Department

		// Scan the result into the enrollment and related fields
		err := rows.Scan(
			&enrollment.EnrolmentId, &enrollment.StudentId, &enrollment.CourseId, &enrollment.EnrollmentDate,
			&student.StudentId, &student.FirstName, &student.LastName, &student.Email, &student.Address, &student.DateOfBirth, &student.CreatedAt, &student.UpdatedAt,
			&course.CourseID, &course.Name, &course.Description, &course.Credits, &course.DepartmentId, &course.CreatedAt, &course.UpdatedAt,
			&department.DepartmentID, &department.Name, &department.Description, &department.CreatedAt, &department.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Assign the student, course, and department to the enrollment
		enrollment.Students = student
		enrollment.Courses = course
		enrollment.Courses.Departments = department

		// Append the enrollment to the list
		enrollments = append(enrollments, &enrollment)
	}

	// Return the list of enrollments
	return enrollments, nil
}
