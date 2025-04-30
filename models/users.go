package models

import (
	"errors"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username     string
	PasswordHash string
	UsedStorage  int64 // in bytes
	TotalStorage int64 // e.g., 50MB
}

var (
	users = make(map[string]*User)
	mu    sync.Mutex
)

func RegisterUser(username, password string) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := users[username]; exists {
		return errors.New("user already exists")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	users[username] = &User{
		Username:     username,
		PasswordHash: string(hashed),
		TotalStorage: 50 * 1024 * 1024, // 50MB
	}

	return nil
}

func AuthenticateUser(username, password string) error {
	mu.Lock()
	defer mu.Unlock()

	user, exists := users[username]
	if !exists {
		return errors.New("user not found")
	}
	return bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
}

func GetUser(username string) *User {
	mu.Lock()
	defer mu.Unlock()
	return users[username]
}
