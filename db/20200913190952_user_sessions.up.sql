CREATE TABLE IF NOT EXISTS user_sessions(
  user_id bigint REFERENCES users(id),
  uuid uuid NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
  last_login timestamptz WITH TIME ZONE,
  last_active timestamptz WITH TIME ZONE
);
