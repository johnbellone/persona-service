CREATE TABLE IF NOT EXISTS realms(
       id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
       name varchar(200) NOT NULL UNIQUE CHECK(name <> ''),
       description text,
       created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
       updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
