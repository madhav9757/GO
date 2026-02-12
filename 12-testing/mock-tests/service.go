package mock_test

import "errors"

// User represents a user in our system
type User struct {
	ID   int
	Name string
}

// Database interface defines the behavior we need
type Database interface {
	GetUser(id int) (*User, error)
}

// UserService uses a Database to perform operations
type UserService struct {
	db Database
}

// NewUserService creates a new UserService
func NewUserService(db Database) *UserService {
	return &UserService{db: db}
}

// GetUserName returns the name of a user or an error
func (s *UserService) GetUserName(id int) (string, error) {
	user, err := s.db.GetUser(id)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not found")
	}
	return user.Name, nil
}
