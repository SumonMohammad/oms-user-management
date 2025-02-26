CREATE TABLE IF NOT EXISTS map_trader_team (
    id SERIAL PRIMARY KEY, -- Primary key for map_trader_team
    team_id INT NOT NULL, -- References trader_team(id)
    trader_id INT NOT NULL, -- References traders(id)
    is_enabled BOOLEAN DEFAULT TRUE, -- Indicates if mapping is enabled
    status VARCHAR(50), -- Optional status (additional information)
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL, -- Creation timestamp
    CONSTRAINT fk_map_team_id FOREIGN KEY (team_id) REFERENCES trader_team(id) ON DELETE CASCADE,
    CONSTRAINT fk_map_trader_id FOREIGN KEY (trader_id) REFERENCES traders(id) ON DELETE CASCADE
    );