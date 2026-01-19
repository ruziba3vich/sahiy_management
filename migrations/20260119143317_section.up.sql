
-- SECTION

CREATE TABLE IF NOT EXISTS "sections" (
    "id"            SERIAL PRIMARY KEY,
    "department_id" INT REFERENCES "departments"("id"),
    "name"          VARCHAR(256),
    "created_at"    BIGINT,
    "updated_at"    BIGINT
);

-- sections-table
