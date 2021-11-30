package repository

import (
	"context"

	"github.com/vinhnv1/babycrab/internal/models"
)

// SpecRepo is the interface for repository methods
type SpecRepo interface {
	CreateFriend(ctx context.Context, userId int, friendId int) error
	GetFriendsByID(ctx context.Context, userId int) (models.FriendSlice, error)
	GetUserBlocksByID(ctx context.Context, userId int) (models.UserBlockSlice, error)
	//GetCommonFriends(ctx context.Context, firstUser int, secondUser int) (models.FriendSlice, error)
	CreateSubscription(ctx context.Context, requestorId int, targetId int) error
	GetRecipientEmails(ctx context.Context, senderId int) ([]models.User, error)
	CreateUserBlock(ctx context.Context, requestorId int, targetId int) error
	IsExistedFriend(ctx context.Context, userId int, friendId int) (bool, error)
	IsBlockedFriend(ctx context.Context, userId int, friendId int) (bool, error)
	IsSubscribedFriend(ctx context.Context, requestorId int, targetId int) (bool, error)
	GetUserIDByEmail(ctx context.Context, email string) (int, error)
	GetEmailsByUserIDs(ctx context.Context, userIDs []int) ([]string, error)
}
