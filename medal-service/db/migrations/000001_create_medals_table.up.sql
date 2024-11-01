CREATE TABLE IF NOT EXISTS medals (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    country_id UUID NOT NULL,
    type INT NOT NULL,
    event_id UUID NOT NULL,
    athlete_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at INT DEFAULT 0
);
