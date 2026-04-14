-- +goose Up
CREATE TABLE vpa_map (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vpa_id VARCHAR(255) UNIQUE NOT NULL,    
    account_id VARCHAR(255) NOT NULL,      
    bank_code VARCHAR(50) NOT NULL,       
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_vpa_lookup ON vpa_map(vpa_id);

-- +goose Down
DROP TABLE IF EXISTS vpa_map;
