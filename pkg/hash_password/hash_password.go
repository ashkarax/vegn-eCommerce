package hashpassword

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// func HashPassword(password string) (string, error) {
// 	p, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	hashedPassword := string(p)
// 	return hashedPassword, err
// }

// func CompareHashedPassword(dbPassword, password string) error {
// 	return bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
// }

func HashPassword(password string) string {

	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err, "problem at hashing ")
	}
	fmt.Println(HashedPassword)
	return string(HashedPassword)
}

func CompairPassword(hashedPassword string, plainPassword string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))

	if err != nil {
		return errors.New("passwords does not match")
	}

	return nil
}
