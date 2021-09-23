CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS books (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name text NOT NULL,
    year text,
    author text,
    category text,
    price text,
    descriptions text
);