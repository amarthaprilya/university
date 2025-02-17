package models

import "time"

type Course struct {
	CourseID     int
	Name         string
	Description  string
	Credits      string
	DepartmentId int
	Departments  Department
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
