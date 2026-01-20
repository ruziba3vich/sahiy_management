-- USERS

CREATE TABLE IF NOT EXISTS "users" (
    "id" SERIAL PRIMARY KEY,
    "department_id" INT REFERENCES "departments"("id"),
    "section_id" INT REFERENCES "sections"("id"),
    "schedule_id" INT REFERENCES "schedules"("id"),
    "role" INT,
    "phone" VARCHAR(16) UNIQUE,
    "full_name" VARCHAR(64),
    "joined_at" BIGINT,
    "created_at" BIGINT,
    "updated_at" BIGINT,
    "tg_chat_id" BIGINT,
    "password_hash" VARCHAR(256)
);

-- Super Admin Seed (Role 99 = SuperAdmin)
-- Credentials:
--   Phone: +998901112233
--   Password: SuperSecure2026!
INSERT INTO "users" (
    "department_id", "section_id", "schedule_id", "role", "phone", "full_name",
    "joined_at", "created_at", "updated_at", "tg_chat_id", "password_hash"
) VALUES (
    NULL, NULL, NULL, 99, '+998901112233', 'Super Admin',
    EXTRACT(EPOCH FROM NOW())::BIGINT,
    EXTRACT(EPOCH FROM NOW())::BIGINT,
    EXTRACT(EPOCH FROM NOW())::BIGINT,
    0,
    '$2a$10$K6ND7pyEIn.gUkVJ.FL0teaD1zuY8ZM8JeqD/S9hQsPFPkrN.PibW'
) ON CONFLICT ("phone") DO NOTHING;

-- users-table
