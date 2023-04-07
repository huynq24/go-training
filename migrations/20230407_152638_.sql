-- +migrate Down
DROP TABLE IF EXISTS product_tags;

-- +migrate Down
DROP TABLE IF EXISTS products;

-- +migrate Down
DROP TABLE IF EXISTS tags;

-- +migrate Down
DROP TABLE IF EXISTS categories;

-- +migrate Up
CREATE TABLE tags (
                      id int NOT NULL AUTO_INCREMENT,
                      title varchar(255) NOT NULL,
                      updated_at datetime DEFAULT NULL,
                      created_at datetime DEFAULT NULL,
                      status int DEFAULT NULL,
                      PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +migrate Up
CREATE TABLE categories (
                            id int NOT NULL AUTO_INCREMENT,
                            title varchar(255) NOT NULL,
                            updated_at datetime DEFAULT NULL,
                            created_at datetime DEFAULT NULL,
                            status int DEFAULT NULL,
                            PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

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