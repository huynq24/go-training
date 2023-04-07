-- +migrate Up
CREATE TABLE tags (
                      id int NOT NULL AUTO_INCREMENT,
                      title varchar(255) NOT NULL,
                      updated_at datetime DEFAULT NULL,
                      created_at datetime DEFAULT NULL,
                      status int DEFAULT NULL,
                      PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +migrate Down
DROP TABLE IF EXISTS tags;