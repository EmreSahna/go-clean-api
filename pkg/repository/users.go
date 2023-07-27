package repository

import (
	"database/sql"
	"gin-example/pkg/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UsersRepo struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *UsersRepo {
	return &UsersRepo{db}
}

func (r *UsersRepo) CreateUser(user models.User) error {
	id := uuid.New().String()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = r.db.Exec("INSERT INTO users (id, username, password) VALUES (?, ?, ?)", id, user.Username, hashedPassword)
	return err
}
