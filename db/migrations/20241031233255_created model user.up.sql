BEGIN;
CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    login      TEXT NOT NULL UNIQUE,
    email      TEXT NOT NULL UNIQUE,
    password   TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT
                                            NOW()
);
COMMIT;