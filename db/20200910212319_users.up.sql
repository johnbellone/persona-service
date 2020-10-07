CREATE TYPE user_state AS ENUM(
  'Inactive',
  'Staged',
  'Pending',
  'Active',
  'Recovery',
  'Locked',
  'Suspended',
  'Expired'
);
CREATE TABLE IF NOT EXISTS users(
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  uuid uuid NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
  created_at timestamptz WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  deleted_at timestamptz WITH TIME ZONE DEFAULT NULL,
  realm_id bigint REFERENCES realms,
  name varchar(200),
  login varchar(30) NOT NULL CHECK (login <> ''),
  email varchar(30) NOT NULL CHECK (email <> ''),
  status user_state NOT NULL DEFAULT 'Staged',
  verified timestamptz WITH TIME ZONE DEFAULT NULL,
  attributes jsonb NOT NULL DEFAULT '{}'::jsonb,
  UNIQUE(realm_id, email),
  UNIQUE(realm_id, login)
);
