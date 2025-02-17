package models

import "time"

type Department struct {
	DepartmentID int
	Name         string
	Description  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
