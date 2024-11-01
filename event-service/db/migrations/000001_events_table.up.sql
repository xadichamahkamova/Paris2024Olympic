CREATE TABLE events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    sport_type VARCHAR(255),
    location VARCHAR(255),
    date DATE NOT NULL,                -- Use DATE type for the event date
    start_time TIME NOT NULL,         -- Use TIME type for the start time
    end_time TIME NOT NULL,           -- Use TIME type for the end time
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);
