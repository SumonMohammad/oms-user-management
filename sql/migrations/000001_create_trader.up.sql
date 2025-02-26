CREATE TABLE IF NOT EXISTS traders (
    id SERIAL PRIMARY KEY,
    user_id INT,
    status VARCHAR(20) NOT NULL,
    office_id INT,
    is_active BOOLEAN DEFAULT FALSE,
    read_only BOOLEAN DEFAULT FALSE,
    licence_number VARCHAR NOT NULL,
    tws_id VARCHAR(255) NOT NULL,
    is_enabled BOOLEAN DEFAULT TRUE,
    can_trade BOOLEAN DEFAULT TRUE,
    is_deleted BOOLEAN DEFAULT FALSE,
    expire_at TIMESTAMPTZ DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL
    CONSTRAINT fk_traders FOREIGN KEY (user_id) REFERENCES oms_user(id) ON DELETE SET NULL,

    );