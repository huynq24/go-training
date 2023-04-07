-- +migrate Up
CREATE TABLE product_tags (
                              product_id int NOT NULL,
                              tag_id int NOT NULL,
                              updated_at datetime DEFAULT NULL,
                              created_at datetime DEFAULT NULL,
                              status int DEFAULT NULL,
                              PRIMARY KEY (product_id,tag_id),
                              KEY tag_id (tag_id),
                              CONSTRAINT product_tags_ibfk_1 FOREIGN KEY (product_id) REFERENCES products (id),
                              CONSTRAINT product_tags_ibfk_2 FOREIGN KEY (tag_id) REFERENCES tags (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +migrate Down
DROP TABLE IF EXISTS product_tags;