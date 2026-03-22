CREATE TABLE children (
    child_id SERIAL PRIMARY KEY,
    pregnancy_id INTEGER NOT NULL REFERENCES pregnancies(pregnancy_id) ON DELETE CASCADE,
    weight DECIMAL(6,3),
    length DECIMAL(6,2),
    heartrate INTEGER,
    notes TEXT,
    nutrition_status VARCHAR(100)
);

