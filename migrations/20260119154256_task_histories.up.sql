
-- TASK_HISTORIES

CREATE TABLE IF NOT EXISTS "task_histories" (
    "id" SERIAL PRIMARY KEY,
    "task_id" INT REFERENCES "tasks"("id"),
    "user_id" INT "users"("id"),
    "status" INT,
    "started_at" BIGINT,
    "finished_at" BIGINT
);

-- task_histories-table
