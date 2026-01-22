-- Update task_statuses table: replace type with section_id

ALTER TABLE task_statuses DROP COLUMN IF EXISTS type;
ALTER TABLE task_statuses ADD COLUMN section_id INTEGER NOT NULL REFERENCES sections(id) ON DELETE CASCADE;
