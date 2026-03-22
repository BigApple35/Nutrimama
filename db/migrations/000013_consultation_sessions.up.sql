CREATE TABLE consultation_sessions (
    consultation_id SERIAL PRIMARY KEY,
    mother_id INTEGER NOT NULL REFERENCES mothers(mother_id) ON DELETE CASCADE,
    consultant_id INTEGER NOT NULL REFERENCES consultants(consultant_id) ON DELETE CASCADE,
    session_date DATE NOT NULL,
    time_start TIME NOT NULL,
    hour_end TIME NOT NULL,
    CHECK (hour_end > time_start)
);