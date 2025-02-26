CREATE TABLE IF NOT EXISTS audit_log (
 id SERIAL PRIMARY KEY, -- Auto-incrementing primary key
 action_type VARCHAR, -- Action type (e.g., login, password change)
 attempt_by_id INT, -- ID of trader, investor, or broker_admin
 attempt_by_type VARCHAR, -- Specifies "trader", "investor", or "broker_admin"
 ip_address VARCHAR, -- IP address of the action
 http_method VARCHAR, -- HTTP method (e.g., GET, POST)
 endpoint VARCHAR, -- The API endpoint accessed
 is_success BOOLEAN NOT NULL, -- Whether the action was successful
 platform VARCHAR, -- Platform (e.g., MAC, Windows)
 device_name VARCHAR, -- Device name or MAC address
 device_type VARCHAR, -- Type of device (e.g., mobile, desktop)
 description JSONB, -- JSON field for additional details (e.g., remarks, session ID)
 request_body JSONB, -- JSON field to store the request body
 response_body JSONB, -- JSON field to store the response body
 created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP, -- Auto-set timestamp when record is created
 updated_at TIMESTAMPTZ DEFAULT NULL -- Timestamp when the record is last updated
);