CREATE TABLE IF NOT EXISTS group_roles(
       group_id bigint REFERENCES groups,
       role_id bigint REFERENCES roles,
       PRIMARY KEY(group_id, role_id)
);
