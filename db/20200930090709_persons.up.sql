CREATE TABLE IF NOT EXISTS persons(
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  uuid uuid NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
  created_at timestamptz WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  deleted_at timestamptz WITH TIME ZONE DEFAULT NULL,

  given_name varchar NOT NULL,
  family_name varchar NOT NULL,
  nick_name varchar,
  additional_name varchar,
  birth_date date,
  birth_place varchar,
  death_place varchar,
  death_date date,

  attributes jsonb NOT NULL DEFAULT '{}'::jsonb
);
