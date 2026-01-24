
-- SECTION BRANCHES

CREATE TABLE IF NOT EXISTS "department_branches" (
    "id" SERIAL PRIMARY KEY,
    "branch_id" INT REFERENCES "branchs"("id") ON DELETE CASCADE,
    "department_id" INT REFERENCES "departments"("id") ON DELETE CASCADE,
    "status" INT,
    "created_at" BIGINT
);

-- section_branches-table
