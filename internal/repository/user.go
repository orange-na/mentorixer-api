package repository

import (
	"database/sql"

	"main/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	rows, err := r.db.Query("SELECT * FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.Name, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) GetByNameAndPassword(name, password string) (model.User, error) {
	var user model.User
	err := r.db.QueryRow("SELECT * FROM user WHERE name = ? AND password = ?", name, password).Scan(&user.Id, &user.Name, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, nil
		}
		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepository) Create(user model.User) error {
	_, err := r.db.Exec("INSERT INTO user (id, name, password) VALUES (?, ?, ?)", user.Id, user.Name, user.Password)
	return err
}