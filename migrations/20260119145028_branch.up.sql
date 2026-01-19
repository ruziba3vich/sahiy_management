-- BRANCHS --

CREATE TABLE IF NOT EXISTS "branchs" (
    "id"            SERIAL PRIMARY KEY,
    "name"          VARCHAR(64),
    "lat"           DECIMAL(8, 6),
    "long"          DECIMAL(8, 6),
    "radius"        INT,
    "status"        INT,
    "created_at"    BIGINT,
    "updated_at"    BIGINT
);

-- branchs-table
