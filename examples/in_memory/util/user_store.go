package util

import "slices"

type InMemoryUserStore struct {
	users []User
}

func NewInMemoryUserStore(users ...User) *InMemoryUserStore {
	return &InMemoryUserStore{users: users}
}

func (u *InMemoryUserStore) VerifyUser(username, password string) bool {
	return slices.ContainsFunc(u.users, func(user User) bool {
		return user.Username == username && user.Password == password
	})
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
