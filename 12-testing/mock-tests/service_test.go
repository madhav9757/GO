package mock_test

import (
	"errors"
	"testing"
)

// MockDatabase is a manual mock implementation of the Database interface
type MockDatabase struct {
	GetUserFn func(id int) (*User, error)
}

func (m *MockDatabase) GetUser(id int) (*User, error) {
	return m.GetUserFn(id)
}

func TestGetUserName(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Setup mock
		mockDB := &MockDatabase{
			GetUserFn: func(id int) (*User, error) {
				return &User{ID: 1, Name: "John Doe"}, nil
			},
		}
		service := NewUserService(mockDB)

		name, err := service.GetUserName(1)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if name != "John Doe" {
			t.Errorf("expected John Doe, got %s", name)
		}
	})

	t.Run("not found", func(t *testing.T) {
		// Setup mock to return error
		mockDB := &MockDatabase{
			GetUserFn: func(id int) (*User, error) {
				return nil, errors.New("db error")
			},
		}
		service := NewUserService(mockDB)

		_, err := service.GetUserName(1)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}
