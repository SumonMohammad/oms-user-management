CREATE TABLE IF NOT EXISTS traders (
    id SERIAL PRIMARY KEY,
    user_id BIGINT,
    status VARCHAR(30) NOT NULL,
    branch_id BIGINT NOT NULL,
    is_active BOOLEAN DEFAULT FALSE,
    read_only BOOLEAN DEFAULT FALSE,
    licence_number VARCHAR NOT NULL,
    can_trade BOOLEAN DEFAULT TRUE,
    is_deleted BOOLEAN DEFAULT FALSE,
    expire_at TIMESTAMPTZ DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL
    );