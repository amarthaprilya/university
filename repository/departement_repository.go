package repository

import (
	"database/sql"
	"fmt"
	"student/models"
)

type DepartmentRepository interface {
	GetAllDepartment() ([]*models.Department, error)
}

type departmentRepository struct {
	db *sql.DB
}

func NewDepartmentRepository(db *sql.DB) *departmentRepository {
	return &departmentRepository{db}
}

func (r *departmentRepository) GetAllDepartment() ([]*models.Department, error) {
	query := `SELECT department_id, name, description, created_at, updated_at FROM departments`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get departments: %w", err)
	}
	defer rows.Close()

	var departments []*models.Department
	for rows.Next() {
		var department models.Department
		err := rows.Scan(&department.DepartmentID, &department.Name, &department.Description, &department.CreatedAt, &department.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan department row: %w", err)
		}
		departments = append(departments, &department)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return departments, nil
}
