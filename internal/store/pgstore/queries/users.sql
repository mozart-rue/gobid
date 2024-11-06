-- name: CreateUser :one
INSERT INTO users (user_name, email, password_hash, bio)
VALUES ($1, $2, $3, $4)
RETURNING id;


-- name: GetUserByID :one
SELECT
  id,
  user_name,
  email,
  password_hash,
  bio,
  updated_at,
  created_at
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT
  id,
  user_name,
  email,
  password_hash,
  bio,
  updated_at,
  created_at
FROM users
WHERE email = $1;
