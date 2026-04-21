package models

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	PassHash  string    `json:"passHash"`
	CreatedAt time.Time `json:"createdAd"`
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
	query := `INSERT INTO users (username, email, passHash, createdAt) 
				VALUES ($1, $2, $3, $4)`

	_, err := ur.db.Exec(
		query,
		newUser.Username,
		newUser.Email,
		newUser.PassHash,
		newUser.CreatedAt.Unix(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetUserByUsername(username string) (User, error) {
	query := "SELECT id, username, passHash, email, createdAt from users WHERE username = $1"

	rows, err := ur.db.Query(query, username)
	if err != nil {
		return User{}, err
	}
	if rows.Next() {
		var user User
		var ts int64

		err = rows.Scan(
			&user.Id,
			&user.Username,
			&user.PassHash,
			&user.Email,
			&ts,
		)
		user.CreatedAt = time.Unix(ts, 0)
		if err != nil {
			return User{}, err
		}
		return user, nil
	}

	return User{}, errors.New("No such user with this username")
}
