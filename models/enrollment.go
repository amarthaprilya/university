package models

import "time"

type Enrollments struct {
	EnrolmentId    int
	StudentId      int
	Students       Student
	CourseId       int
	Courses        Course
	EnrollmentDate time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type EnrollmentsParam struct {
	StudentId      int       `json:"student_id"`
	CourseId       int       `json:"course_id"`
	EnrollmentDate time.Time `json:"enrollment_date"`
}
