CREATE TABLE IF NOT EXISTS trader_team (
    id SERIAL PRIMARY KEY, -- Primary key for trader_team
    name VARCHAR(255) NOT NULL, -- Team name
    description VARCHAR(255), -- Description of the team
    is_enabled BOOLEAN DEFAULT TRUE, -- Indicates if the team is enabled
    status VARCHAR(50), -- Optional status (can use enum-like behavior if needed)
    is_deleted BOOLEAN DEFAULT FALSE, -- Indicates if the team is deleted
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL
    );