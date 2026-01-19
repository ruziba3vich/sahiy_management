-- USERS

CREATE TABLE IF NOT EXISTS "users" (
    "id" SERIAL PRIMARY KEY,
    "department_id" REFERENCES "departments"("id"),
    "section_id" REFERENCES "sections"("id"),
    "schedule_id" REFERENCES "schedules"("id"),
    "role" INT,
    "phone" VARCHAR(16),
    "full_name" VARCHAR(64),
    "joined_at" BIGINT,
    "created_at" BIGINT,
    "updated_at" BIGINT,
    "tg_chat_id" BIGINT,
    "password_hash" VARCHAR(256)
);

-- users-table
