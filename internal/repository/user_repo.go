package repository

import (
	"crudapp/internal/domain"
	"database/sql"
	"errors"
)

type UserRepository interface {
	Create(user *domain.User) 	error
	GetById(id int64) 			(*domain.User, error)
	GetAll() 					([]*domain.User, error)
	Update(user *domain.User) 	error
	Delete(id int64) 			error
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) Create(user *domain.User) error {
	query := "INSERT INTO users (name, email) VALUES ($1, $2)"
	return r.db.QueryRow(query, user.Name, user.Email).Scan(&user.Id)
}

func (r *userRepo) GetById(id int64) (*domain.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = $1"
	user := &domain.User{}
	err := r.db.QueryRow(query, id).Scan(&user.Id, &user.Name, &user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("пользователь не найден")
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepo) GetAll() ([]*domain.User, error) {
	query := "SELECT id, name, email FROM users"
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}	
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		user := &domain.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepo) Update(user *domain.User) error {
	query := "UPDATE users SET name = $1, email = $2 WHERE id = $3"
	_, err := r.db.Exec(query, user.Name, user.Email, user.Id)
	return err
}

func (r *userRepo) Delete(id int64) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}