-- Add user_id (assigned user) to tasks table
ALTER TABLE "tasks" ADD COLUMN "user_id" INT REFERENCES "users"("id");

-- Create index for faster filtering by user_id
CREATE INDEX IF NOT EXISTS "idx_tasks_user_id" ON "tasks"("user_id");

-- Create index for faster filtering by status
CREATE INDEX IF NOT EXISTS "idx_tasks_status" ON "tasks"("status");

-- Create index for faster filtering by section_id
CREATE INDEX IF NOT EXISTS "idx_tasks_section_id" ON "tasks"("section_id");

-- Create index for faster filtering by priority
CREATE INDEX IF NOT EXISTS "idx_tasks_priority" ON "tasks"("priority");
