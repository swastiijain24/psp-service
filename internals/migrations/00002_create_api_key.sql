-- +goose Up
CREATE TABLE psp_registrations (
    psp_id         VARCHAR(250) PRIMARY KEY,
    hashed_api_key VARCHAR(255) NOT NULL,
    psp_name       VARCHAR(100) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS psp_registrations;

