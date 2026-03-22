CREATE TABLE nutrition_tracking (
    track_id SERIAL PRIMARY KEY,
    mother_id INTEGER REFERENCES mothers(mother_id) ON DELETE CASCADE,
    child_id INTEGER REFERENCES children(child_id) ON DELETE CASCADE,
    protein DECIMAL(6,2),
    calories DECIMAL(7,2),
    hydration DECIMAL(6,2),
    tracking_type VARCHAR(20) CHECK (tracking_type IN ('daily', 'weekly', 'monthly')),
    omega_3 DECIMAL(6,2),
    iron DECIMAL(6,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CHECK (mother_id IS NOT NULL OR child_id IS NOT NULL)
);