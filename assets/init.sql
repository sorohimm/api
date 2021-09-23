CREATE TABLE books IF NOT EXISTS (
    uuid UUID primary key DEFAULT gen_random_uuid(),
    name text NOT NULL,
    year text,
    author text,
    category text,
    price integer,
    descriptions text
);