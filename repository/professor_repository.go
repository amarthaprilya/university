package repository

import (
	"database/sql"
	"fmt"
	"student/models"
)

type ProfessorRepository interface {
	GetAllProfessor() ([]*models.Professor, error)
}

type professorRepository struct {
	db *sql.DB
}

func NewProfessorRepository(db *sql.DB) *professorRepository {
	return &professorRepository{db}
}

func (r *professorRepository) GetAllProfessor() ([]*models.Professor, error) {
	query := `SELECT professor_id, first_name, last_name, email, password, address, created_at, updated_at FROM professors`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching professors: %v", err)
	}
	defer rows.Close()

	var professors []*models.Professor
	for rows.Next() {
		var professor models.Professor
		err := rows.Scan(&professor.ProfessorId, &professor.FirstName, &professor.LastName, &professor.Email,
			&professor.Password, &professor.Address, &professor.CreatedAt, &professor.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning professor row: %v", err)
		}
		professors = append(professors, &professor)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over professors rows: %v", err)
	}

	return professors, nil
}
