CREATE TABLE IF NOT EXISTS refresh_tokens(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    expired_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    created_by TEXT NOT NULL,
    updated_by TEXT NOT NULL,
    CONSTRAINT fk_user_id_refresh_tokens FOREIGN KEY(user_id) REFERENCES users(id)
);