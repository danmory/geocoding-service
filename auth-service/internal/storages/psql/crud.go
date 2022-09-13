package psql

import (
	"context"

	"github.com/danmory/geocoding-service/auth-service/internal/core"
)

func GetByUsername(username string) (*core.User, error) {
	db := GetDatabase()
	row := db.QueryRow(
		context.Background(),
		"SELECT username, password FROM users WHERE username=$1", username)
	var user core.User
	if err := row.Scan(&user.Username, &user.Password); err != nil {
		return nil, err
	}
	return &user, nil
}

func SaveUser(username, hashedPassword string) error {
	db := GetDatabase()
	if _, err := db.Exec(
		context.Background(),
		"INSERT INTO users VALUES($1, $2)",
		username,
		hashedPassword); err != nil {
		return err
	}
	return nil
}
