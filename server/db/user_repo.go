package db

import (
	"database/sql"
	"paper-trader/model"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return UserRepo{db: db}
}

func (r *UserRepo) CreateUser(email string) (*model.User, error) {
	user := new(model.User)
	err := r.db.QueryRow("INSERT INTO users (email) VALUES ($1) RETURNING id", email).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	user.Email = email
	return user, nil
}

func (r *UserRepo) GetUserById(id int32) (*model.User, error) {
	user := new(model.User)
	err := r.db.QueryRow("SELECT id, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) GetUserByEmail(email string) (*model.User, error) {
	user := new(model.User)
	err := r.db.QueryRow("SELECT id, email FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) UpdateUser(user *model.User) error {
	_, err := r.db.Exec("UPDATE users SET email = $1 WHERE id = $2", user.Email, user.ID)
	return err
}

func (r *UserRepo) DeleteUser(id int32) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
