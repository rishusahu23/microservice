CREATE EXTENSION IF NOT EXISTS pgcrypto; -- Ensure pgcrypto is enabled for UUID generation

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- UUID as the primary key with auto-generation
    username VARCHAR(50) NOT NULL UNIQUE,          -- Unique username
    email VARCHAR(255) NOT NULL UNIQUE,            -- Unique email address
    password_hash TEXT NOT NULL,                   -- Hashed password for security
    created_at TIMESTAMP DEFAULT NOW(),            -- Record creation timestamp
    updated_at TIMESTAMP DEFAULT NOW()             -- Last updated timestamp
    );
