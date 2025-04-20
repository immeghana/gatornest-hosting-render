-- Remove the not-null constraint from student_id
ALTER TABLE students ALTER COLUMN student_id DROP NOT NULL;

-- Add a unique constraint only when student_id is not null
CREATE UNIQUE INDEX idx_students_student_id_not_null ON students (student_id) WHERE student_id IS NOT NULL; 