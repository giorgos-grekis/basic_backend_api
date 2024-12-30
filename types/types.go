package types

import "time"

// IUserStore => Interface UserStore
// type IUserStore interface {
// 	GetUserByEmail(email string) (*User, error)
// 	GetUserByID(id int) (*User, error)
// 	CreateUser(User) error
// }

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	// GetUserByID(id int) (*User, error)
	// CreateUser(User) error
}

// type mockUserStore struct{}

// func GetUserByEmail(email string) (*User, error) {
// 	return nil, nil
// }

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
