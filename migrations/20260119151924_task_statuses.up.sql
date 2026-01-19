-- TASK_STATUSES

CREATE TABLE IF NOT EXISTS "task_statuses" (
    "id" SERIAL PRIMARY,
    "name" VARCHAR(64),
    "type" INT
);

-- task_statuses-table
