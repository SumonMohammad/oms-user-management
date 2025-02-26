CREATE TABLE IF NOT EXISTS map_role_permission (
    id SERIAL PRIMARY KEY,
    permission_id INT NOT NULL,
    role_id INT NOT NULL,
    is_enabled BOOLEAN,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL
    );