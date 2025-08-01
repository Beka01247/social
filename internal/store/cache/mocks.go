package cache

import (
	"context"

	"github.com/Beka01247/social/internal/store"
	"github.com/stretchr/testify/mock"
)

func NewMockStore() Storage {
	return Storage{
		Users: &MockUserStore{},
	}
}

type MockUserStore struct {
	mock.Mock
}

func (m *MockUserStore) Get(ctx context.Context, userID int64) (*store.User, error) {
	args := m.Called(userID)
	user := args.Get(0)
	if user == nil {
		return nil, args.Error(1)
	}
	return user.(*store.User), args.Error(1)
}

func (m *MockUserStore) Set(ctx context.Context, user *store.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserStore) Delete(ctx context.Context, userID int64) error {
	args := m.Called(userID)
	return args.Error(0)
}
