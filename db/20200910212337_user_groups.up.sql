CREATE TABLE IF NOT EXISTS user_groups(
       user_id bigint REFERENCES users(id),
       group_id bigint REFERENCES groups,
       PRIMARY KEY(user_id, group_id)
);
