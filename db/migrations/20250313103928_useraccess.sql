-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_access (
                             id INT AUTO_INCREMENT PRIMARY KEY,
                             user_id INT NOT NULL,
                             token CHAR(26) NOT NULL UNIQUE, -- ULID
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             expires_at TIMESTAMP NOT NULL,
                             FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_access;
-- +goose StatementEnd
