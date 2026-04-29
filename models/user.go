package models

import (
	"database/sql"
	"log"
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

func (ur *UserRepository) InsertUser(newUser User) (int64, error) {
	query := `INSERT INTO users 
    (email, password_hash, first_name, last_name, role, created_at) VALUES (?, ?, ?, ?, ?, ?)`

	res, err := ur.db.Exec(
		query,
		newUser.Email,
		newUser.PasswordHash,
		newUser.FirstName,
		newUser.LastName,
		newUser.Role,
		newUser.CreatedAt,
	)
	if err != nil {
		return 0, err
	}
	newUser.ID, _ = res.LastInsertId()

	if newUser.Role == RoleParent {
		query := `INSERT INTO parents (user_id) VALUES (?)`
		_, err := ur.db.Exec(query, newUser.ID)
		if err != nil {
			return 0, err
		}
		log.Default().Println("Parent created")
	}

	return res.LastInsertId()
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

func (ur *UserRepository) GetUserByEmail(email string) (User, error) {
	query := `
	SELECT id, email, password_hash, first_name, last_name, role, created_at
	FROM users
	WHERE email = ?
	`

	row := ur.db.QueryRow(query, email)

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

func (ur *UserRepository) DoesEmailExist(email string) (bool, error) {
	var exists bool
	err := ur.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", email).Scan(&exists)
	if err != nil {
		panic(err)
	}
	return exists, nil
}
