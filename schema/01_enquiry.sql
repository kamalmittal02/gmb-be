CREATE TABLE enquiries (
                           id BIGSERIAL PRIMARY KEY,
                           name VARCHAR(255) NOT NULL,
                           phone VARCHAR(50) NOT NULL,
                           email VARCHAR(255),
                           message JSONB NOT NULL,
                           is_active BOOLEAN DEFAULT true,
                           created_at TIMESTAMPTZ DEFAULT now()
);