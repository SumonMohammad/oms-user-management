CREATE TABLE IF NOT EXISTS health (
    id SERIAL PRIMARY KEY,
    status VARCHAR(50) NOT NULL,

    created_at  TIMESTAMPTZ DEFAULT NOW(),
    updated_at  TIMESTAMPTZ DEFAULT NULL,
    deleted_at  TIMESTAMPTZ DEFAULT NULL
);
