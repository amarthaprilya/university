package models

import "time"

type Student struct {
	StudentId   int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	Address     string
	DateOfBirth time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type StudentParam struct {
	FirstName   string
	LastName    string
	Email       string
	Password    string
	Address     string
	DateOfBirth time.Time
}

type StudentLoginParam struct {
	Email    string
	Password string
}
