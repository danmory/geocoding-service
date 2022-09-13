package service

import (
	"context"
	"errors"

	"github.com/danmory/geocoding-service/auth-service/internal/core"
	"github.com/danmory/geocoding-service/auth-service/internal/storages/psql"
	"github.com/jackc/pgx/v4"
)

func retrieveUser(username string) (*core.User, error) {
	db := psql.GetDatabase()
	row := db.QueryRow(
		context.Background(),
		"SELECT username, password FROM users WHERE username=$1", username) // TODO: to CRUD
	var user core.User
	err := row.Scan(&user.Username, &user.Password)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func RegisterUser(user *core.User) (string, error) {
	existingUser, err := retrieveUser(user.Username)
	if err != nil {
		return "", err
	}
	if existingUser != nil {
		return "", errors.New("user already exists")
	}
	db := psql.GetDatabase()
	if !isPasswordStrong(user.Password) {
		return "", errors.New("weak password")
	}
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return "", err
	}
	if _, err = db.Exec(
		context.Background(),
		"INSERT INTO users VALUES($1, $2)",
		user.Username,
		hashedPassword); err != nil {
		return "", err
	} // TODO: to CRUD
	token, err := generateJWT(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func LoginUser(user *core.User) (string, error) {
	existingUser, err := retrieveUser(user.Username)
	if err != nil {
		return "", err
	}
	if !checkPasswordHash(user.Password, existingUser.Password) {
		return "", errors.New("unknown user")

	}
	token, err := generateJWT(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}
