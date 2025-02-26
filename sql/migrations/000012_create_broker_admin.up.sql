CREATE TABLE IF NOT EXISTS broker_admin (
    id SERIAL PRIMARY KEY,
    employee_id INT,
    user_name VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255),
    manager_id INT,
    office_id INT,
    can_trade BOOLEAN,
    can_login BOOLEAN,
    read_only BOOLEAN,
    type VARCHAR(255),
    status VARCHAR(255),
    is_isolated_user BOOLEAN,
    is_deleted BOOLEAN,
    expire_at TIMESTAMPTZ DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    CONSTRAINT fk_broker_admin_employee_id FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE,
    CONSTRAINT fk_broker_admin_office_id FOREIGN KEY (office_id) REFERENCES offices(id) ON DELETE CASCADE

    );