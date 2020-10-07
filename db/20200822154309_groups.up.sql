CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS groups(
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  created_at timestamptz WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  deleted_at timestamptz WITH TIME ZONE DEFAULT NULL,
  uuid uuid NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
  name varchar(30) NOT NULL UNIQUE CHECK(name <> ''),
  description text,
  attributes jsonb NOT NULL DEFAULT '{}'::jsonb
);
