CREATE TABLE IF NOT EXISTS map_user_permission (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    permission_id INT NOT NULL,
    is_enabled BOOLEAN,
    is_revoked BOOLEAN,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL
    );