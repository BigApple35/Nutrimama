CREATE TABLE meal_plans (
    meal_plan_id SERIAL PRIMARY KEY,
    mother_id INTEGER NOT NULL REFERENCES mothers(mother_id) ON DELETE CASCADE,
    week INTEGER NOT NULL,
    phase VARCHAR(50) NOT NULL
);