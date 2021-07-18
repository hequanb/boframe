package usermodels

import (
	"errors"
	
	"boframe/models"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	models.Model
	
	UserId   int64  `gorm:"column:user_id;type:int;unique;notnull"`
	Username string `gorm:"column:username;type:varchar(255);notnull;unique"`
	NickName string `gorm:"column:nick_name;type:varchar(255);notnull"`
	Password string `gorm:"column:password;type:varchar(255);notnull"`
	Email    string `gorm:"column:email;type:varchar(255)"`
	Gender   int8   `gorm:"column:gender;type:tinyint;notnull"`
}


var InvalidPasswordErr = errors.New("invalid password")

func Encrypt(password string) (passwordHash string, err error) {
	if len(password) <= 0 {
		return "", InvalidPasswordErr
	}
	pHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	passwordHash = string(pHash)
	return passwordHash, nil
}

func Decrypt(passwordHash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
