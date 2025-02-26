CREATE TABLE IF NOT EXISTS broker_admin (
    id SERIAL PRIMARY KEY,
    user_id BIGINT,
    branch_id BIGINT,
    can_trade BOOLEAN,
    read_only BOOLEAN,
    status VARCHAR(30),
    is_isolated_user BOOLEAN,
    is_deleted BOOLEAN,
    expire_at TIMESTAMPTZ DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL

    );