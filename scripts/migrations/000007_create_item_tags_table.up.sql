CREATE TABLE IF NOT EXISTS item_tags(
    item_id BIGINT NOT NULL,
    tag_id BIGINT NOT NULL,
    PRIMARY KEY(item_id, tag_id),
    CONSTRAINT fk_item_id FOREIGN KEY(item_id) REFERENCES items(id),
    CONSTRAINT fk_tag_id FOREIGN KEY(tag_id) REFERENCES tags(id)
);