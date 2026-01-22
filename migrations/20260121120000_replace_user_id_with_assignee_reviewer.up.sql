-- Replace user_id with assignee_id and reviewer_id in tasks table

-- Drop the old user_id index
DROP INDEX IF EXISTS "idx_tasks_user_id";

-- Rename user_id to assignee_id
ALTER TABLE "tasks" RENAME COLUMN "user_id" TO "assignee_id";

-- Add reviewer_id column
ALTER TABLE "tasks" ADD COLUMN "reviewer_id" INT REFERENCES "users"("id");

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS "idx_tasks_assignee_id" ON "tasks"("assignee_id");
CREATE INDEX IF NOT EXISTS "idx_tasks_reviewer_id" ON "tasks"("reviewer_id");
