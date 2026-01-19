-- SCHEDULES --

CREATE TABLE IF NOT EXISTS "schedules" (
    "id"            BIGSERIAL PRIMARY KEY,
    "name"          VARCHAR(255),
    "timezone"      SMALLINT,
    "week_days"     JSON,
    "status"        SMALLINT,
    "created_at"    BIGINT,
    "updated_at"    BIGINT
);

-- schedules-table
