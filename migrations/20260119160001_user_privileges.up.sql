-- USER_PRIVILEGES (Many-to-Many junction table)

CREATE TABLE IF NOT EXISTS "user_privileges" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INT NOT NULL REFERENCES "users"("id") ON DELETE CASCADE,
    "privilege_id" INT NOT NULL REFERENCES "privileges"("id") ON DELETE CASCADE,
    "granted_by" INT REFERENCES "users"("id") ON DELETE SET NULL,
    "granted_at" BIGINT NOT NULL,
    UNIQUE("user_id", "privilege_id")
);

CREATE INDEX IF NOT EXISTS "idx_user_privileges_user_id" ON "user_privileges"("user_id");
CREATE INDEX IF NOT EXISTS "idx_user_privileges_privilege_id" ON "user_privileges"("privilege_id");

-- user_privileges-table
