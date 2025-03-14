CREATE TABLE IF NOT EXISTS map_trader_tws (
    id SERIAL PRIMARY KEY,
    tws_id INT UNIQUE NOT NULL,
    trader_id INT UNIQUE, -- A trader can only have one TWS ID
    status VARCHAR(50) NOT NULL DEFAULT 'UNASSIGNED', -- e.g., ASSIGNED, UNASSIGNED
    is_enabled BOOLEAN DEFAULT FALSE,
    is_deleted BOOLEAN DEFAULT FALSE,
    expire_at TIMESTAMPTZ DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL
    );