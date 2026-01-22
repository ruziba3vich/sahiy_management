-- Rollback: Replace assignee_id and reviewer_id with user_id

-- Drop new indexes
DROP INDEX IF EXISTS "idx_tasks_assignee_id";
DROP INDEX IF EXISTS "idx_tasks_reviewer_id";

-- Drop reviewer_id column
ALTER TABLE "tasks" DROP COLUMN "reviewer_id";

-- Rename assignee_id back to user_id
ALTER TABLE "tasks" RENAME COLUMN "assignee_id" TO "user_id";

-- Recreate old user_id index
CREATE INDEX IF NOT EXISTS "idx_tasks_user_id" ON "tasks"("user_id");
