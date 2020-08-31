CREATE TABLE IF NOT EXISTS user_groups(
       user_id bigint REFERENCES users(id),
       group_id bigint REFERENCES groups,
       created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
       updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
