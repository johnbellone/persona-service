CREATE TABLE IF NOT EXISTS groups(
       id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
       name varchar(30) NOT NULL UNIQUE CHECK(name <> ''),
       description text,
       meta jsonb NOT NULL DEFAULT '{}'::jsonb,
       created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
       updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
