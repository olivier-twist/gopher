package user

import (
	"strings"
)

type User struct {
	firstName string
	lastName  string
}

func (u *User) CreateNew(fn, ln string) *User {
	u.firstName = fn
	u.lastName = ln
	return u
}

func (u User) String() string {
	names := []string{u.firstName, u.lastName}
	return strings.Join(names, "  ")
}
