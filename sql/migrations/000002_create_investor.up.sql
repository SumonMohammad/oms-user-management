
CREATE TABLE IF NOT EXISTS investors (
    id SERIAL PRIMARY KEY, -- Primary key for the investor table
    user_id INT NOT NULL, -- References oms_user(id)
    primary_tws_id INT NOT NULL, -- References tws(id) (must not be null)
    secondary_tws_id INT DEFAULT NULL, -- References tws(id) (optional)
    bo_account_number VARCHAR(255), -- BO account number
    status VARCHAR(255), -- Investor status
    client_code VARCHAR(255),
    can_trade BOOLEAN DEFAULT TRUE, -- Indicates if the investor can trade (default true)
    read_only BOOLEAN, -- Indicates if the investor has read-only access
    is_deleted BOOLEAN, -- Soft delete flag
    expire_at TIMESTAMPTZ DEFAULT NULL, -- Expiry timestamp
    created_at TIMESTAMPTZ DEFAULT NOW(), -- Auto-inserted timestamp
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL-- Auto-updated timestamp

    );