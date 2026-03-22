CREATE TABLE consultants (
    consultant_id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    facility_name VARCHAR(255),
    rating DECIMAL(3,2) CHECK (rating >= 0 AND rating <= 5)
);

