package models

import "time"

type Teachings struct {
	TeachingId  int
	ProfessorId int
	Professors  Professor
	CourseId    int
	Courses     Course
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
