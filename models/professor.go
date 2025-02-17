package models

import "time"

type Professor struct {
	ProfessorId int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	Address     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
