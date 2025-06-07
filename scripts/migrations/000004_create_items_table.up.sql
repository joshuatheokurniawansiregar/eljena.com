CREATE TABLE IF NOT EXISTS items(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    sub_category_id BIGINT NOT NULL,
    item_name TEXT NOT NULL,
    item_description TEXT NOT NULL,
    item_image_url TEXT NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_sub_category_id FOREIGN KEY(sub_category_id) REFERENCES sub_categories(id)
);