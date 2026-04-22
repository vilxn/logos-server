package models

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID           int64  `db:"id" json:"id"`
	Email        string `db:"email" json:"email"`
	PasswordHash string `db:"password_hash" json:"-"`

	FirstName string `db:"first_name" json:"firstName"`
	LastName  string `db:"last_name" json:"lastName"`

	Role      UserRole  `db:"role" json:"role"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{
		db: database,
	}
}

func (ur *UserRepository) InsertUser(newUser User) error {
	query := `INSERT INTO users 
    (email, password_hash, first_name, last_name, role, created_at) VALUES (?, ?, ?, ?, ?, ?)`

	_, err := ur.db.Exec(
		query,
		newUser.Email,
		newUser.PasswordHash,
		newUser.FirstName,
		newUser.LastName,
		newUser.Role,
		newUser.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetUserByID(id int64) (User, error) {
	query := `
	SELECT id, email, password_hash, first_name, last_name, role, created_at
	FROM users
	WHERE id = ?
	`

	row := ur.db.QueryRow(query, id)

	var user User

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		return User{}, err
	}

	return user, nil
}
