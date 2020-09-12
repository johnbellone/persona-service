CREATE TABLE IF NOT EXISTS roles(
       id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
       created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
       updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
       deleted_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

       uuid uuid NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
       name varchar(30) NOT NULL UNIQUE CHECK(name <> ''),
       description text,
       meta jsonb NOT NULL DEFAULT '{}'::jsonb
);
