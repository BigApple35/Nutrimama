CREATE TABLE educational_tools (
    edu_tools_id SERIAL PRIMARY KEY,
    publisher VARCHAR(255),
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);