-- Revert task_statuses table: replace section_id with type

ALTER TABLE task_statuses DROP COLUMN IF EXISTS section_id;
ALTER TABLE task_statuses ADD COLUMN type INTEGER;
