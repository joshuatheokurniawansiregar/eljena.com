CREATE TABLE IF NOT EXISTS sub_categories(
 id BIGSERIAL PRIMARY KEY NOT NULL,
 user_id BIGINT NOT NULL,
 category_id BIGINT NOT NULL,
 sub_category_name TEXT NOT NULL,
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NULL,
 CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id),
 CONSTRAINT fk_category_id FOREIGN KEY(category_id) REFERENCES categories(id)
);