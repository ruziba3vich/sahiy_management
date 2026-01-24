
-- TASKS

CREATE TABLE IF NOT EXISTS "tasks" (
    "id" SERIAL PRIMARY KEY,
    "parent_id" INT REFERENCES "tasks"("id"),
    "section_id" INT REFERENCES "sections"("id") ON DELETE CASCADE,
    "title" TEXT,
    "description" TEXT,
    "priority" INT,
    "deadline" BIGINT,
    "status" INT,
    "created_at" BIGINT,
    "updated_at" BIGINT
);

-- tasks-table
