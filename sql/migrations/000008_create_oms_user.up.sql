CREATE TABLE IF NOT EXISTS oms_user (
    id SERIAL PRIMARY KEY, -- Primary key for the table
    user_name VARCHAR(255) NOT NULL UNIQUE, -- Unique username
    password_hash VARCHAR(255), -- Password hash for authentication
    reference_type VARCHAR(50) NOT NULL, -- Enum for reference type
    reference_id INT, -- ID of the related trader or investor
    can_login BOOLEAN,
    nid VARCHAR(255), -- National ID
    status VARCHAR(255), -- Status of the user
    is_verified BOOLEAN,
    is_enabled BOOLEAN,
    last_login TIMESTAMPTZ DEFAULT NULL,
    expire_at TIMESTAMPTZ DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    CONSTRAINT fk_reference_id_traders FOREIGN KEY (reference_id) REFERENCES traders(id),
    CONSTRAINT fk_reference_id_investors FOREIGN KEY (reference_id) REFERENCES investors(id)
);
