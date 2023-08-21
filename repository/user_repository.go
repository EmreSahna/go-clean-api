package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"go-clean-api/model"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) createUser(user model.User) error {
	id := uuid.New().String()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = r.DB.Exec("INSERT INTO users (id, username, password) VALUES (?, ?, ?)", id, user.Username, hashedPassword)
	return err
}

func (r *UserRepository) getUsers(users *[]model.User) error {
	rows, err := r.DB.Query("SELECT id, username, password FROM users")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			return err
		}
		*users = append(*users, user)
	}

	return nil
}

func (r *UserRepository) getUser(id string, user *model.User) error {
	row := r.DB.QueryRow("SELECT id, username, password FROM users WHERE id = ?", id)
	if err := row.Scan(&user.Id, &user.Username, &user.Password); err != nil {
		return err
	}
	return nil
}
