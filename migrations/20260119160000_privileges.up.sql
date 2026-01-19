-- PRIVILEGES

CREATE TABLE IF NOT EXISTS "privileges" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(64) UNIQUE NOT NULL,
    "resource" VARCHAR(32) NOT NULL,
    "action" VARCHAR(16) NOT NULL,
    "description" VARCHAR(256),
    "created_at" BIGINT NOT NULL,
    "updated_at" BIGINT NOT NULL,
    UNIQUE("resource", "action")
);

-- Seed privileges for all resources
INSERT INTO "privileges" ("name", "resource", "action", "description", "created_at", "updated_at") VALUES
    ('departments:create', 'departments', 'create', 'Create new departments', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('departments:read', 'departments', 'read', 'View departments', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('departments:update', 'departments', 'update', 'Update departments', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('departments:delete', 'departments', 'delete', 'Delete departments', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('sections:create', 'sections', 'create', 'Create new sections', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('sections:read', 'sections', 'read', 'View sections', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('sections:update', 'sections', 'update', 'Update sections', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('sections:delete', 'sections', 'delete', 'Delete sections', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('users:create', 'users', 'create', 'Create new users', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('users:read', 'users', 'read', 'View users', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('users:update', 'users', 'update', 'Update users', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('users:delete', 'users', 'delete', 'Delete users', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('task-statuses:create', 'task-statuses', 'create', 'Create new task statuses', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('task-statuses:read', 'task-statuses', 'read', 'View task statuses', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('task-statuses:update', 'task-statuses', 'update', 'Update task statuses', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('task-statuses:delete', 'task-statuses', 'delete', 'Delete task statuses', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('tasks:create', 'tasks', 'create', 'Create new tasks', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('tasks:read', 'tasks', 'read', 'View tasks', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('tasks:update', 'tasks', 'update', 'Update tasks', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('tasks:delete', 'tasks', 'delete', 'Delete tasks', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('task-histories:create', 'task-histories', 'create', 'Create task history entries', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('task-histories:read', 'task-histories', 'read', 'View task histories', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('task-histories:update', 'task-histories', 'update', 'Update task histories', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('task-histories:delete', 'task-histories', 'delete', 'Delete task histories', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('branches:create', 'branches', 'create', 'Create new branches', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('branches:read', 'branches', 'read', 'View branches', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('branches:update', 'branches', 'update', 'Update branches', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('branches:delete', 'branches', 'delete', 'Delete branches', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('schedules:create', 'schedules', 'create', 'Create new schedules', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('schedules:read', 'schedules', 'read', 'View schedules', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('schedules:update', 'schedules', 'update', 'Update schedules', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('schedules:delete', 'schedules', 'delete', 'Delete schedules', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('user-actions:create', 'user-actions', 'create', 'Create user actions', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('user-actions:read', 'user-actions', 'read', 'View user actions', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('user-actions:update', 'user-actions', 'update', 'Update user actions', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('user-actions:delete', 'user-actions', 'delete', 'Delete user actions', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('privileges:create', 'privileges', 'create', 'Create new privileges', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('privileges:read', 'privileges', 'read', 'View privileges', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('privileges:update', 'privileges', 'update', 'Update privileges', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('privileges:delete', 'privileges', 'delete', 'Delete privileges', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('privileges:assign', 'privileges', 'assign', 'Assign privileges to users', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT),
    ('privileges:revoke', 'privileges', 'revoke', 'Revoke privileges from users', EXTRACT(EPOCH FROM NOW())::BIGINT, EXTRACT(EPOCH FROM NOW())::BIGINT)
ON CONFLICT ("name") DO NOTHING;

-- privileges-table
