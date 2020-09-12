CREATE TYPE user_state AS ENUM('Inactive', 'Staged', 'Pending', 'Active', 'Recovery', 'Locked', 'Suspended', 'Expired');
CREATE TABLE IF NOT EXISTS users(
       id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
       uuid uuid NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
       created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
       updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
       deleted_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

       realm_id bigint REFERENCES realms,
       name varchar(200),
       login varchar(30) NOT NULL CHECK (slug <> ''),
       email varchar(30) NOT NULL CHECK (email <> ''),
       password varchar(128) NOT NULL CHECK (password <> ''),
       status user_state NOT NULL DEFAULT 'Staged',
       verified boolean NOT NULL DEFAULT FALSE,
       last_password_change timestamp WITH TIME ZONE,
       last_login timestamp WITH TIME ZONE,
       last_active timestamp WITH TIME ZONE,
       activated timestamp WITH TIME ZONE DEFAULT NULL,
       meta jsonb NOT NULL DEFAULT '{}'::jsonb,
       UNIQUE(realm_id, email),
       UNIQUE(realm_id, slug)
);
