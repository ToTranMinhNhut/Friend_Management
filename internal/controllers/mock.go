package controllers

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/vinhnv1/babycrab/internal/models"
)

type SpecRepo struct {
	mock.Mock
}

func (m SpecRepo) CreateFriend(ctx context.Context, userId int, friendId int) error {
	args := m.Called(ctx, userId, friendId)
	var r error
	if args.Get(0) != nil {
		r = args.Get(0).(error)
	}
	return r
}

func (m SpecRepo) GetFriendsByID(ctx context.Context, userId int) (models.FriendSlice, error) {
	args := m.Called(userId)
	t := args.Get(0).(models.FriendSlice)
	return t, args.Error(1)
}

func (m SpecRepo) GetUserBlocksByID(ctx context.Context, userId int) (models.UserBlockSlice, error) {
	args := m.Called(userId)
	r1 := args.Get(0).(models.UserBlockSlice)

	var r2 error
	if args.Get(1) != nil {
		r2 = args.Get(1).(error)
	}
	return r1, r2
}

/* func (m SpecRepo) GetCommonFriends(ctx context.Context, firstUser int, secondUser int) (models.FriendSlice, error) {
	args := m.Called(ctx, firstUser, secondUser)
	r1 := args.Get(0).(models.FriendSlice)

	var r2 error
	if args.Get(1) != nil {
		r2 = args.Get(1).(error)
	}
	return r1, r2
} */

func (m SpecRepo) CreateSubscription(ctx context.Context, requestorId int, targetId int) error {
	args := m.Called(ctx, requestorId, targetId)
	var r error
	if args.Get(0) != nil {
		r = args.Get(0).(error)
	}
	return r
}

func (m SpecRepo) GetRecipientEmails(ctx context.Context, senderId int) ([]models.User, error) {
	args := m.Called(ctx, senderId)
	r1 := args.Get(0).([]models.User)

	var r2 error
	if args.Get(1) != nil {
		r2 = args.Get(1).(error)
	}
	return r1, r2
}

func (m SpecRepo) CreateUserBlock(ctx context.Context, requestorId int, targetId int) error {
	args := m.Called(ctx, requestorId, targetId)
	var r error
	if args.Get(0) != nil {
		r = args.Get(0).(error)
	}
	return r
}

func (m SpecRepo) IsExistedFriend(ctx context.Context, userId int, friendId int) (bool, error) {
	args := m.Called(ctx, userId, friendId)
	r1 := args.Get(0).(bool)

	var r2 error
	if args.Get(1) != nil {
		r2 = args.Get(1).(error)
	}
	return r1, r2
}

func (m SpecRepo) IsBlockedFriend(ctx context.Context, userId int, friendId int) (bool, error) {
	args := m.Called(ctx, userId, friendId)
	r1 := args.Get(0).(bool)

	var r2 error
	if args.Get(1) != nil {
		r2 = args.Get(1).(error)
	}
	return r1, r2
}

func (m SpecRepo) IsSubscribedFriend(ctx context.Context, requestorId int, targetId int) (bool, error) {
	args := m.Called(ctx, requestorId, targetId)
	r1 := args.Get(0).(bool)

	var r2 error
	if args.Get(1) != nil {
		r2 = args.Get(1).(error)
	}
	return r1, r2
}

func (m SpecRepo) GetUserIDByEmail(ctx context.Context, email string) (int, error) {
	args := m.Called(email)

	t := args.Get(0).(int)
	return t, args.Error(1)
}

func (m SpecRepo) GetEmailsByUserIDs(ctx context.Context, userIDs []int) ([]string, error) {
	args := m.Called(userIDs)
	r1 := args.Get(0).([]string)

	var r2 error
	if args.Get(1) != nil {
		r2 = args.Get(1).(error)
	}
	return r1, r2
}
