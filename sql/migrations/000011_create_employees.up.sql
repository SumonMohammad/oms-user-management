CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    branch_id BIGINT NOT NULL,
    designation VARCHAR(255),
    description VARCHAR(255),
    expire_at TIMESTAMPTZ DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL
    );