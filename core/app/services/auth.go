package services

import (
	"context"
	"database/sql"
	"encoding/hex"
	"errors"
	"playpost/core/app/databases"
	"playpost/core/app/models"
	coreutils "playpost/core/app/utils"
)

type AuthService struct{}

// Register new user. Returns saved user id and any error
func (AuthService) Register(userdata *models.Register) (int64, error) {
	if userdata.Password != userdata.PasswordMatch {
		return 0, errors.New("Passwords don't match.")
	}
	var db *sql.DB
	db = databases.GetDB()

	rows, err := db.Query("SELECT * FROM user WHERE email = ?", userdata.Email)
	if err != nil {
		return 0, err
	}

	if rows != nil {

	}
	defer rows.Close()

	if rows.Next() {
		return 0, errors.New("Email already registered")
	}

	salt, err := coreutils.GenerateRandomSalt(coreutils.SALTSIZE)
	if err != nil {
		return 0, err
	}

	hashedPassword, err := coreutils.HashPassword(userdata.Password, salt)
	if err != nil {
		return 0, err
	}

	// save to database
	dml := "INSERT INTO `user`(`create_time`, `email`, `username`, `full_name`, `password`, `password_salt`) " +
		"VALUES(NOW(), ?, ?, ?, ?, ?)"

	insertResult, err := db.ExecContext(context.Background(), dml, userdata.Email, userdata.Username, userdata.Fullname, hashedPassword, hex.EncodeToString(salt))
	if err != nil {
		return 0, err
	}

	return insertResult.LastInsertId()
}
