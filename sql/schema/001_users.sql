-- +goose UP
CREATE TABLE users (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE users;
