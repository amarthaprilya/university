-- a. Daftar semua student, termasuk informasi kontak mereka dan mata kuliah yang saat ini mereka ikuti.
SELECT s.student_id, s.first_name, s.last_name, s.email, s.address, 
       c.course_id, c.name AS course_name, c.description AS course_description
FROM students s
LEFT JOIN enrollments e ON s.student_id = e.student_id
LEFT JOIN courses c ON e.course_id = c.course_id;

-- b. Daftar semua mata kuliah, termasuk departemen tempat mereka berada, dosen yang mengajarkannya, dan student yang saat ini terdaftar di dalamnya.
SELECT c.course_id, c.name AS course_name, c.description AS course_description, 
       d.department_id, d.name AS department_name, 
       p.professor_id, p.first_name AS professor_first_name, p.last_name AS professor_last_name, 
       s.student_id, s.first_name AS student_first_name, s.last_name AS student_last_name
FROM courses c
LEFT JOIN departments d ON c.department_id = d.department_id
LEFT JOIN teachings t ON c.course_id = t.course_id
LEFT JOIN professors p ON t.professor_id = p.professor_id
LEFT JOIN enrollments e ON c.course_id = e.course_id
LEFT JOIN students s ON e.student_id = s.student_id;

-- c. Daftar semua dosen, termasuk informasi kontak mereka dan mata kuliah yang mereka ajarkan.
SELECT p.professor_id, p.first_name, p.last_name, p.email, p.address, 
       c.course_id, c.name AS course_name, c.description AS course_description
FROM professors p
LEFT JOIN teachings t ON p.professor_id = t.professor_id
LEFT JOIN courses c ON t.course_id = c.course_id;

-- d. Tanggal pendaftaran dan kredit mata kuliah untuk setiap pendaftaran student di setiap mata kuliah.
SELECT e.enrollment_date, c.credits, s.student_id, s.first_name, s.last_name, c.course_id, c.name AS course_name
FROM enrollments e
JOIN students s ON e.student_id = s.student_id
JOIN courses c ON e.course_id = c.course_id;

-- e. Daftar semua departemen dan mata kuliah yang termasuk ke dalam setiap departemen.
SELECT d.department_id, d.name AS department_name, d.description AS department_description, 
       c.course_id, c.name AS course_name, c.description AS course_description
FROM departments d
LEFT JOIN courses c ON d.department_id = c.department_id;

-- f. Jumlah total student yang terdaftar di setiap mata kuliah.
SELECT c.course_id, c.name AS course_name, COUNT(e.student_id) AS total_students
FROM courses c
LEFT JOIN enrollments e ON c.course_id = e.course_id
GROUP BY c.course_id, c.name;

-- g. Rata-rata jumlah student yang terdaftar dalam mata kuliah di setiap departemen.
SELECT d.department_id, d.name AS department_name, 
       AVG(student_count) AS average_students_per_course
FROM (
    SELECT c.department_id, COUNT(e.student_id) AS student_count
    FROM courses c
    LEFT JOIN enrollments e ON c.course_id = e.course_id
    GROUP BY c.course_id, c.department_id
) course_counts
JOIN departments d ON course_counts.department_id = d.department_id
GROUP BY d.department_id, d.name;
