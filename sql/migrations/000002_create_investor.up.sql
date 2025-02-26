-- CREATE TABLE IF NOT EXISTS investors (
--     id SERIAL PRIMARY KEY,
--     user_id INT,
--     primary_tws_id INT NOT NULL,
--     secondary_tws_id INT,
--     bo_account_number VARCHAR,
--     read_only BOOLEAN DEFAULT TRUE,
--     status VARCHAR(50) NOT NULL,
--     is_enabled BOOLEAN DEFAULT FALSE,
--     can_trade BOOLEAN DEFAULT TRUE,
--     is_deleted BOOLEAN DEFAULT FALSE,
--     expire_at TIMESTAMPTZ DEFAULT NULL,
--     created_at TIMESTAMPTZ DEFAULT NOW(),
--     updated_at TIMESTAMPTZ DEFAULT NULL,
--     deleted_at TIMESTAMPTZ DEFAULT NULL
--     );


CREATE TABLE IF NOT EXISTS investors (
    id SERIAL PRIMARY KEY, -- Primary key for the investor table
    user_id INT NOT NULL, -- References oms_user(id)
    primary_tws_id INT NOT NULL, -- References tws(id) (must not be null)
    secondary_tws_id INT DEFAULT NULL, -- References tws(id) (optional)
    bo_account_number VARCHAR(255), -- BO account number
    status VARCHAR(255), -- Investor status
    is_enabled BOOLEAN, -- Indicates if the investor is enabled
    can_trade BOOLEAN DEFAULT TRUE, -- Indicates if the investor can trade (default true)
    read_only BOOLEAN, -- Indicates if the investor has read-only access
    is_deleted BOOLEAN, -- Soft delete flag
    expire_at TIMESTAMPTZ DEFAULT NULL, -- Expiry timestamp
    created_at TIMESTAMPTZ DEFAULT NOW(), -- Auto-inserted timestamp
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL,-- Auto-updated timestamp
    CONSTRAINT fk_investor_user FOREIGN KEY (user_id) REFERENCES oms_user(id) ON DELETE CASCADE,
    CONSTRAINT fk_investor_primary_tws FOREIGN KEY (primary_tws_id) REFERENCES tws(id) ON DELETE CASCADE,
    CONSTRAINT fk_investor_secondary_tws FOREIGN KEY (secondary_tws_id) REFERENCES tws(id) ON DELETE SET NULL
    );