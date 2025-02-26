CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY, -- Primary key for the table
    user_name VARCHAR(255) NOT NULL, -- Unique username
    auth_id BIGINT,
    user_type VARCHAR(50) NOT NULL, -- Enum for reference type
    phone_number VARCHAR(50) NOT NULL UNIQUE,
    email_address VARCHAR(255) NOT NULL UNIQUE,
    country_code VARCHAR(15) NOT NULL,
    nid VARCHAR(255), -- National ID
    can_login BOOLEAN,
    is_verified BOOLEAN,
    is_enabled BOOLEAN,
    last_login TIMESTAMPTZ DEFAULT NULL,
    expire_at TIMESTAMPTZ DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL
    );
