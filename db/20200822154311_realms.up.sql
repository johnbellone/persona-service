CREATE TABLE IF NOT EXISTS realms(
       id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
       uuid uuid NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
       created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
       updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
       deleted_at timestamp WITH TIME ZONE DEFAULT NULL,

       name varchar(200) NOT NULL UNIQUE CHECK(name <> ''),
       description text,
       attributes jsonb NOT NULL DEFAULT '{}'::jsonb
);
