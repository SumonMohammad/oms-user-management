CREATE TABLE IF NOT EXISTS tws (
    id SERIAL PRIMARY KEY,
    tws_code VARCHAR(255) UNIQUE NOT NULL,
    is_enabled BOOLEAN DEFAULT FALSE,
    is_active  BOOLEAN DEFAULT FALSE,
    is_deleted BOOLEAN DEFAULT FALSE,
    status VARCHAR(50) NOT NULL DEFAULT 'UNASSIGNED', -- e.g., ASSIGNED, UNASSIGNED
    expire_at TIMESTAMPTZ DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL
    );