-- USER ACTION --

CREATE TABLE IF NOT EXISTS "user_actions" (
    "id"                SERIAL PRIMARY KEY,
    "user_id"           INT REFERENCES "users"("id") ON DELETE CASCADE,
    "visit_branch_id"   INTEGER NOT NULL REFERENCES "branchs"("id") ON DELETE CASCADE,
    "leave_branch_id"   INTEGER REFERENCES "branchs"("id") ON DELETE CASCADE,
    "come_status"       SMALLINT,
    "out_status"        SMALLINT,
    "started_at"        BIGINT,
    "finished_at"       BIGINT
);

CREATE INDEX IF NOT EXISTS "user_actions_user_id_idx" ON "user_actions"("user_id");
CREATE INDEX IF NOT EXISTS "user_actions_visit_branch_id_idx" ON "user_actions"("visit_branch_id");
CREATE INDEX IF NOT EXISTS "user_actions_leave_branch_id_idx" ON "user_actions"("leave_branch_id");

-- user_action-table
