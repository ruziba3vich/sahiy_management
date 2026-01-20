-- Remove indexes
DROP INDEX IF EXISTS "idx_tasks_priority";
DROP INDEX IF EXISTS "idx_tasks_section_id";
DROP INDEX IF EXISTS "idx_tasks_status";
DROP INDEX IF EXISTS "idx_tasks_user_id";

-- Remove user_id column
ALTER TABLE "tasks" DROP COLUMN IF EXISTS "user_id";
