package user

const (
	createUserStmt = `INSERT INTO users (username, password, role) VALUES ($1, $2, $3)`

	getUserByIDStmt = `SELECT id, username, password, role FROM users WHERE id = $1`

	getUserByUsernameStmt = `SELECT id, username, password, role FROM users WHERE username = $1`
)
