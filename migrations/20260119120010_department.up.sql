-- DEPARTMENS --

CREATE TABLE IF NOT EXISTS "departments" (
    "id"            SERIAL PRIMARY KEY,
    "name"          VARCHAR(64),
    "status"        INT,
    "created_at"    BIGINT,
    "updated_at"    BIGINT
);

-- departments-table
