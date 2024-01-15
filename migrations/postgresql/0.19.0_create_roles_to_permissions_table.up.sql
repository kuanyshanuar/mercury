CREATE TABLE roles_to_permissions (
    role_id       BIGINT  REFERENCES roles (id),
    permission_id BIGINT  REFERENCES permissions (id),
    is_active     BOOLEAN DEFAULT FALSE,
    created_at    BIGINT,
    updated_at    BIGINT
);
