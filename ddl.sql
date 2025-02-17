-- CREATE DATABASE university;
-- \c university;

-- Table for departments
CREATE TABLE IF NOT EXISTS departments (
    department_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);


-- Table for courses
CREATE TABLE IF NOT EXISTS courses (
    course_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    credits VARCHAR(50),
    department_id INT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (department_id) REFERENCES departments(department_id) ON DELETE SET NULL
);

-- Table for students
CREATE TABLE IF NOT EXISTS students (
    student_id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    address TEXT,
    date_of_birth DATE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);


-- Table for professors
CREATE TABLE IF NOT EXISTS professors (
    professor_id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    address TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);


-- Table for enrollments
CREATE TABLE IF NOT EXISTS enrollments (
    enrolment_id SERIAL PRIMARY KEY,
    student_id INT NOT NULL,
    course_id INT NOT NULL,
    enrollment_date DATE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (student_id) REFERENCES students(student_id) ON DELETE CASCADE,
    FOREIGN KEY (course_id) REFERENCES courses(course_id) ON DELETE CASCADE
);

-- Table for teachings
CREATE TABLE IF NOT EXISTS teachings (
    teaching_id SERIAL PRIMARY KEY,
    professor_id INT NOT NULL,
    course_id INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (professor_id) REFERENCES professors(professor_id) ON DELETE CASCADE,
    FOREIGN KEY (course_id) REFERENCES courses(course_id) ON DELETE CASCADE
);

INSERT INTO departments (name, description)
VALUES ('Computer Science', 'Department for computer science studies');
INSERT INTO courses (name, description, credits, department_id)
VALUES ('Introduction to Programming', 'Basic programming concepts.', '3', 1);
INSERT INTO students (first_name, last_name, email, password, address, date_of_birth)
VALUES 
('Amartha', 'Student', 'amartha.student@example.com', '$2a$04$9FRlG4umpd1Q/7h7/znfsuoqgkX1TLkG5aFCLthj2dQldZNycF2hu', 'Jl. Jakarta No.1', '2001-05-15');
INSERT INTO professors (first_name, last_name, email, password, address)
VALUES 
('Amartha', 'Professor', 'amartha.professor@example.com', '$2a$04$9FRlG4umpd1Q/7h7/znfsuoqgkX1TLkG5aFCLthj2dQldZNycF2hu', 'Jl. Surabaya No.2');
INSERT INTO enrollments (student_id, course_id, enrollment_date)
VALUES (1, 1, CURRENT_DATE);
INSERT INTO teachings (professor_id, course_id)
VALUES (1, 1);


-- -- Trigger function to update `updated_at` column
-- CREATE OR REPLACE FUNCTION update_timestamp()
-- RETURNS TRIGGER AS $$
-- BEGIN
--    NEW.updated_at = CURRENT_TIMESTAMP;
--    RETURN NEW;
-- END;
-- $$ LANGUAGE plpgsql;

-- -- Create triggers for updating `updated_at` for each table
-- CREATE TRIGGER update_department_timestamp
-- BEFORE UPDATE ON departments
-- FOR EACH ROW
-- EXECUTE FUNCTION update_timestamp();

-- CREATE TRIGGER update_course_timestamp
-- BEFORE UPDATE ON courses
-- FOR EACH ROW
-- EXECUTE FUNCTION update_timestamp();

-- CREATE TRIGGER update_student_timestamp
-- BEFORE UPDATE ON students
-- FOR EACH ROW
-- EXECUTE FUNCTION update_timestamp();

-- CREATE TRIGGER update_professor_timestamp
-- BEFORE UPDATE ON professors
-- FOR EACH ROW
-- EXECUTE FUNCTION update_timestamp();

-- CREATE TRIGGER update_enrollment_timestamp
-- BEFORE UPDATE ON enrollments
-- FOR EACH ROW
-- EXECUTE FUNCTION update_timestamp();

-- CREATE TRIGGER update_teaching_timestamp
-- BEFORE UPDATE ON teachings
-- FOR EACH ROW
-- EXECUTE FUNCTION update_timestamp();

