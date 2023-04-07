-- +migrate Up
CREATE TABLE products (
                          id int NOT NULL AUTO_INCREMENT,
                          title varchar(255) NOT NULL,
                          image varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                          description text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
                          category_id int DEFAULT NULL,
                          updated_at datetime DEFAULT NULL,
                          created_at datetime DEFAULT NULL,
                          status int DEFAULT NULL,
                          PRIMARY KEY (id),
                          KEY category_id (category_id),
                          CONSTRAINT products_ibfk_1 FOREIGN KEY (category_id) REFERENCES categories (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +migrate Down
DROP TABLE IF EXISTS products;