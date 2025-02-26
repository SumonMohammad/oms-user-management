CREATE TABLE IF NOT EXISTS broker_houses (
    id SERIAL PRIMARY KEY,
    broker_house_name VARCHAR(255),
    short_name VARCHAR(255),
    description VARCHAR(255),
    address VARCHAR(255),
    phone_number VARCHAR(255),
    telephone_number VARCHAR(255),
    email_address VARCHAR(255),
    valid_currency VARCHAR(255),
    country_code VARCHAR(255),
    status VARCHAR(255),
    is_enabled BOOLEAN,
    expire_at TIMESTAMPTZ DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL
    );