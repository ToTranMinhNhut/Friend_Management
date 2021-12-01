package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type FriendRequest struct {
	Emails []string `json:"friends"`
}

type UserRequest struct {
	Email string `json:"email"`
}

type RequestorRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

type RecipientsRequest struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

// Create a new friend relationship
func (_self FriendController) CreateFriend(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	friendReq := FriendRequest{}
	if err := json.NewDecoder(r.Body).Decode(&friendReq); err != nil {
		Respond(w, http.StatusBadRequest, MsgError(ErrBodyRequestInvalid))
		return
	}

	// Validate request body
	if err := friendReq.Validate(); err != nil {
		Respond(w, http.StatusBadRequest, MsgError(err))
		return
	}

	// Get user id and friend id from repository
	userId, err := _self.Repo.GetUserIDByEmail(ctx, friendReq.Emails[0])
	if err != nil {
		Respond(w, http.StatusInternalServerError, MsgError(fmt.Errorf("%s is not exists", friendReq.Emails[0])))
		return
	}
	friendId, err := _self.Repo.GetUserIDByEmail(ctx, friendReq.Emails[1])
	if err != nil {
		Respond(w, http.StatusInternalServerError, MsgError(fmt.Errorf("%s is not exists", friendReq.Emails[1])))
		return
	}

	// Check friend relationship is exists
	isExisted, err := _self.Repo.IsExistedFriend(ctx, userId, friendId)
	if err != nil {
		Respond(w, http.StatusInternalServerError, MsgError(err))
		return
	}
	if isExisted {
		Respond(w, http.StatusInternalServerError, MsgError(ErrExistedFriendship))
		return
	}

	// check blocking between 2 emails
	isBlocked, err := _self.Repo.IsBlockedUser(ctx, userId, friendId)
	if err != nil {
		Respond(w, http.StatusInternalServerError, MsgError(err))
		return
	}
	if isBlocked {
		Respond(w, http.StatusInternalServerError, MsgError(ErrExistedBlockedUser))
		return
	}

	//Call services to create friend relationship
	if err := _self.Repo.CreateFriend(ctx, userId, friendId); err != nil {
		Respond(w, http.StatusInternalServerError, MsgError(ErrCreatedFriendship))
		return
	}

	Respond(w, http.StatusOK, MsgOK())
}

// Get all of friends of a user without blocking relationship
func (_self FriendController) GetFriends(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	userReq := UserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		Respond(w, http.StatusBadRequest, MsgError(ErrBodyRequestInvalid))
		return
	}

	// Validation request body
	if err := userReq.Validate(); err != nil {
		Respond(w, http.StatusBadRequest, MsgError(err))
		return
	}

	// Get user id from an email
	userId, err := _self.Repo.GetUserIDByEmail(ctx, userReq.Email)
	if err != nil {
		Respond(w, http.StatusInternalServerError, MsgError(fmt.Errorf("%s is not exists", userReq.Email)))
		return
	}

	// Get friends available
	friendEmails, err := _self.getFriendEmailsWithoutBlocking(ctx, userId)
	if err != nil {
		Respond(w, http.StatusInternalServerError, MsgError(err))
		return
	}

	Respond(w, http.StatusOK, MsgGetFriendsOk(friendEmails, len(friendEmails)))
}

// Get common friends of 2 users
func (_self FriendController) GetCommonFriends(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	friendReq := FriendRequest{}
	if err := json.NewDecoder(r.Body).Decode(&friendReq); err != nil {
		Respond(w, http.StatusBadRequest, MsgError(ErrBodyRequestInvalid))
		return
	}

	// Validate request body
	if err := friendReq.Validate(); err != nil {
		Respond(w, http.StatusBadRequest, MsgError(err))
		return
	}

	// Get user id and friend id from repository
	firstUserID, err := _self.Repo.GetUserIDByEmail(ctx, friendReq.Emails[0])
	if err != nil {
		Respond(w, http.StatusInternalServerError, MsgError(fmt.Errorf("%s is not exists", friendReq.Emails[0])))
		return
	}
	secondUserID, err := _self.Repo.GetUserIDByEmail(ctx, friendReq.Emails[1])
	if err != nil {
		Respond(w, http.StatusInternalServerError, MsgError(fmt.Errorf("%s is not exists", friendReq.Emails[1])))
		return
	}

	// Get friends of first user and second user
	firstFriendEmails, err := _self.getFriendEmailsWithoutBlocking(ctx, firstUserID)
	if err != nil {
		Respond(w, http.StatusInternalServerError, MsgError(err))
		return
	}
	secondFriendEmails, err := _self.getFriendEmailsWithoutBlocking(ctx, secondUserID)
	if err != nil {
		Respond(w, http.StatusInternalServerError, MsgError(err))
		return
	}

	//Get common friends
	commonFriends := make([]string, 0)
	commonMap := make(map[string]bool)
	for _, firstEmail := range firstFriendEmails {
		commonMap[firstEmail] = true
	}

	for _, secondEmail := range secondFriendEmails {
		if _, ok := commonMap[secondEmail]; ok {
			commonFriends = append(commonFriends, secondEmail)
		}
	}

	Respond(w, http.StatusOK, MsgGetFriendsOk(commonFriends, len(commonFriends)))
}

// Create a subscription relationship of users
func (_self FriendController) CreateSubcription(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	//Decode request body
	requestorReq := RequestorRequest{}
	if err := json.NewDecoder(r.Body).Decode(&requestorReq); err != nil {
		Respond(w, http.StatusBadRequest, MsgError(ErrBodyRequestInvalid))
		return
	}

	//Validate request
	if err := requestorReq.Validate(); err != nil {
		Respond(w, http.StatusBadRequest, MsgError(err))
		return
	}

	// Get user id and friend id from repository
	requestorId, err := _self.Repo.GetUserIDByEmail(ctx, requestorReq.Requestor)
	if err != nil {
		Respond(w, http.StatusInternalServerError, MsgError(fmt.Errorf("%s is not exists", requestorReq.Requestor)))
		return
	}
	targetId, err := _self.Repo.GetUserIDByEmail(ctx, requestorReq.Target)
	if err != nil {
		Respond(w, http.StatusInternalServerError, MsgError(fmt.Errorf("%s is not exists", requestorReq.Target)))
		return
	}

	// Check subscription relationship is exists
	isSubscribed, err := _self.Repo.IsSubscribedUser(ctx, requestorId, targetId)
	if err != nil {
		Respond(w, http.StatusInternalServerError, MsgError(err))
		return
	}
	if isSubscribed {
		Respond(w, http.StatusInternalServerError, MsgError(ErrExistedSubscription))
		return
	}

	// check blocking between 2 user
	isBlocked, err := _self.Repo.IsBlockedUser(ctx, requestorId, targetId)
	if err != nil {
		Respond(w, http.StatusInternalServerError, MsgError(err))
		return
	}
	if isBlocked {
		Respond(w, http.StatusInternalServerError, MsgError(ErrExistedBlockedUser))
		return
	}

	//Call services
	if err := _self.Repo.CreateSubscription(ctx, requestorId, targetId); err != nil {
		Respond(w, http.StatusInternalServerError, MsgError(err))
		return
	}

	// Response
	Respond(w, http.StatusOK, MsgOK())
}

// Create a blocking relationship of users
func (_self FriendController) CreateUserBlock(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	requestorReq := RequestorRequest{}
	if err := json.NewDecoder(r.Body).Decode(&requestorReq); err != nil {
		Respond(w, http.StatusBadRequest, MsgError(ErrBodyRequestInvalid))
		return
	}

	//Validate request
	if err := requestorReq.Validate(); err != nil {
		Respond(w, http.StatusBadRequest, MsgError(err))
		return
	}

	// Get user id and friend id from repository
	requestorId, err := _self.Repo.GetUserIDByEmail(ctx, requestorReq.Requestor)
	if err != nil {
		Respond(w, http.StatusInternalServerError, MsgError(fmt.Errorf("%s is not exists", requestorReq.Requestor)))
		return
	}
	targetId, err := _self.Repo.GetUserIDByEmail(ctx, requestorReq.Target)
	if err != nil {
		Respond(w, http.StatusInternalServerError, MsgError(fmt.Errorf("%s is not exists", requestorReq.Target)))
		return
	}

	// check blocking between 2 user
	isBlocked, err := _self.Repo.IsBlockedUser(ctx, requestorId, targetId)
	if err != nil {
		Respond(w, http.StatusInternalServerError, MsgError(err))
		return
	}
	if isBlocked {
		Respond(w, http.StatusInternalServerError, MsgError(ErrExistedBlockedUser))
		return
	}

	//Call services
	if err := _self.Repo.CreateUserBlock(ctx, requestorId, targetId); err != nil {
		Respond(w, http.StatusInternalServerError, MsgError(err))
		return
	}

	// Response
	Respond(w, http.StatusOK, MsgOK())

}

// Get all of recipients who are friend, subscriber, and mention user without blocking by user
func (_self FriendController) GetRecipientEmails(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	recipient := RecipientsRequest{}
	if err := json.NewDecoder(r.Body).Decode(&recipient); err != nil {
		Respond(w, http.StatusBadRequest, MsgError(err))
		return
	}

	// Validate request body
	if err := recipient.Validate(); err != nil {
		Respond(w, http.StatusBadRequest, MsgError(err))
		return
	}

	// Check existed email and get userID
	senderID, err := _self.Repo.GetUserIDByEmail(ctx, recipient.Sender)
	if err != nil {
		Respond(w, http.StatusInternalServerError, MsgError(fmt.Errorf("%s is not exists", recipient.Sender)))
		return
	}

	//Call services
	recipients, err := _self.Repo.GetRecipientEmails(ctx, senderID)
	if err != nil {
		Respond(w, http.StatusInternalServerError, MsgError(err))
		return
	}

	result := make([]string, 0)
	existedEmailsMap := make(map[string]bool)

	for _, user := range recipients {
		result = append(result, user.Email)
		existedEmailsMap[user.Email] = true
	}

	//Add mentionedEmails to result
	mentionedEmails := GetMentionedEmailFromText(recipient.Text)
	for _, email := range mentionedEmails {
		if _, ok := existedEmailsMap[email]; !ok {
			result = append(result, email)
		}
	}

	Respond(w, http.StatusOK, MsgGetEmailReceiversOk(result))
}

// Get all of users
func (_self FriendController) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	if r.ContentLength != 0 {
		Respond(w, http.StatusBadRequest, MsgError(ErrBodyRequestInvalid))
		return
	}

	users, err := _self.Repo.GetUsers(ctx)
	if err != nil {
		Respond(w, http.StatusInternalServerError, MsgError(err))
		return
	}

	emails := []string{}
	for _, user := range users {
		emails = append(emails, user.Email)
	}

	Respond(w, http.StatusOK, MsgGetAllUsersOk(emails, len(users)))
}

// Get emails of users who are not being blocked by user
func (_self FriendController) getFriendEmailsWithoutBlocking(ctx context.Context, userId int) ([]string, error) {
	// get friends
	friendSlice, err := _self.Repo.GetFriendsByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	friendIds := make([]int, 0)
	for _, friend := range friendSlice {
		if friend.UserID == userId {
			friendIds = append(friendIds, friend.FriendID)
		}
		if friend.FriendID == userId {
			friendIds = append(friendIds, friend.UserID)
		}
	}

	//Get list friends who have blocked user
	userBlocksSlice, err := _self.Repo.GetUserBlocksByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	userBlocksID := make([]int, 0)
	for _, user := range userBlocksSlice {
		if user.RequestorID == userId {
			userBlocksID = append(userBlocksID, user.TargetID)
		}
		if user.TargetID == userId {
			userBlocksID = append(userBlocksID, user.RequestorID)
		}
	}
	//Get UserID list with no blocked
	blockList := make(map[int]bool)
	for _, id := range userBlocksID {
		blockList[id] = true
	}
	friendIDsNonBlock := make([]int, 0)
	for _, id := range friendIds {
		if _, isBlock := blockList[id]; !isBlock {
			friendIDsNonBlock = append(friendIDsNonBlock, id)
		}
	}
	emails, err := _self.Repo.GetEmailsByUserIDs(ctx, friendIDsNonBlock)
	if err != nil {
		return nil, err
	}
	return emails, nil
}
