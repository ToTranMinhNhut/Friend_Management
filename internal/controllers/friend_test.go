package controllers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/vinhnv1/babycrab/internal/models"
)

func TestControllers_GetFriends(t *testing.T) {
	tcs := map[string]struct {
		input              string
		expResult          string
		expError           error
		mockUser           models.User
		mockFriendSlice    models.FriendSlice
		mockUserBlockSlice models.UserBlockSlice
	}{
		"success with an input": {
			input:    `{"Email":"andy@example.com"}`,
			mockUser: models.User{ID: 100, Name: "John", Email: "john@example.com"},
			mockFriendSlice: models.FriendSlice{
				&models.Friend{UserID: 100, FriendID: 101},
				&models.Friend{UserID: 100, FriendID: 102},
			},
			mockUserBlockSlice: models.UserBlockSlice{
				&models.UserBlock{RequestorID: 100, TargetID: 102},
			},
			expResult: `{"count":1,"friends":["andy@example.com"],"success":true}`,
		},
		"failed with an unknow format input": {
			input:    `{}`,
			expError: errors.New(`{"message":"Request body is empty","success":false}`),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/v1/friends", bytes.NewBuffer([]byte(tc.input)))
			require.NoError(t, err)

			var mockRepo SpecRepo
			mockRepo.ExpectedCalls = []*mock.Call{
				mockRepo.On("GetUserIDByEmail", mock.Anything, mock.Anything).Return(tc.mockUser.ID, nil),
				mockRepo.On("GetFriendsByID", mock.Anything, mock.Anything).Return(tc.mockFriendSlice, nil),
				mockRepo.On("GetUserBlocksByID", mock.Anything, mock.Anything).Return(tc.mockUserBlockSlice, nil),
				mockRepo.On("GetEmailsByUserIDs", mock.Anything, mock.Anything).Return([]string{"andy@example.com"}, nil),
			}
			friendController := NewFriendController(mockRepo)
			handler := http.HandlerFunc(friendController.GetFriends)
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if tc.expError != nil {
				require.EqualError(t, tc.expError, rr.Body.String())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, rr.Body.String())
			}
		})
	}
}

//TODO
func TestControllers_CreateFriends(t *testing.T) {
	tcs := map[string]struct {
		input          string
		expResult      string
		expError       error
		mockFirstUser  models.User
		mockSecondUser models.User
	}{
		"success with an input": {
			input:          `{ "friends": ["andy@example.com","john@example.com"]}`,
			mockFirstUser:  models.User{ID: 100, Name: "John", Email: "john@example.com"},
			mockSecondUser: models.User{ID: 101, Name: "Andy", Email: "andy@example.com"},
			expResult:      `{"success":true}`,
		},
		"failed with an unknow format input": {
			input:    `{}`,
			expError: errors.New(`{"message":"Request body is empty","success":false}`),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/v1/friends", bytes.NewBuffer([]byte(tc.input)))
			require.NoError(t, err)

			var mockRepo SpecRepo
			mockRepo.ExpectedCalls = []*mock.Call{
				mockRepo.On("GetUserIDByEmail", mock.Anything, mock.Anything).Return(tc.mockFirstUser.ID, nil),
				mockRepo.On("GetUserIDByEmail", mock.Anything, mock.Anything).Return(tc.mockSecondUser.ID, nil),
				mockRepo.On("IsExistedFriend", mock.Anything, mock.Anything, mock.Anything).Return(false, nil),
				mockRepo.On("IsBlockedFriend", mock.Anything, mock.Anything, mock.Anything).Return(false, nil),
				mockRepo.On("CreateFriend", mock.Anything, mock.Anything, mock.Anything).Return(nil),
			}
			friendController := NewFriendController(mockRepo)
			handler := http.HandlerFunc(friendController.CreateFriend)
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if tc.expError != nil {
				require.EqualError(t, tc.expError, rr.Body.String())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, rr.Body.String())
			}
		})
	}
}
func TestControllers_GetCommonFriends(t *testing.T) {
	tcs := map[string]struct {
		input                    string
		expResult                string
		expError                 error
		mockFirstUser            models.User
		mockFirstFriendSlice     models.FriendSlice
		mockFirstUserBlockSlice  models.UserBlockSlice
		mockSecondUser           models.User
		mockSecondFriendSlice    models.FriendSlice
		mockSecondUserBlockSlice models.UserBlockSlice
	}{
		"success with an input": {
			input:         `{ "friends": ["andy@example.com","john@example.com"]}`,
			mockFirstUser: models.User{ID: 100, Name: "John", Email: "john@example.com"},
			mockFirstFriendSlice: models.FriendSlice{
				&models.Friend{UserID: 100, FriendID: 101},
				&models.Friend{UserID: 100, FriendID: 102},
			},
			mockFirstUserBlockSlice: models.UserBlockSlice{
				&models.UserBlock{RequestorID: 100, TargetID: 101},
			},
			mockSecondUser: models.User{ID: 101, Name: "Andy", Email: "andy@example.com"},
			mockSecondFriendSlice: models.FriendSlice{
				&models.Friend{UserID: 101, FriendID: 103},
				&models.Friend{UserID: 101, FriendID: 102},
			},
			mockSecondUserBlockSlice: models.UserBlockSlice{
				&models.UserBlock{RequestorID: 100, TargetID: 103},
			},

			expResult: `{"count":1,"friends":["common@example.com"],"success":true}`,
		},
		"failed with an unknow format input": {
			input:    `{}`,
			expError: errors.New(`{"message":"Request body is empty","success":false}`),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/v1/commonFriends", bytes.NewBuffer([]byte(tc.input)))
			require.NoError(t, err)

			var mockRepo SpecRepo
			mockRepo.ExpectedCalls = []*mock.Call{
				mockRepo.On("GetUserIDByEmail", mock.Anything, mock.Anything).Return(tc.mockFirstUser.ID, nil),
				mockRepo.On("GetFriendsByID", mock.Anything, mock.Anything).Return(tc.mockFirstFriendSlice, nil),
				mockRepo.On("GetUserBlocksByID", mock.Anything, mock.Anything).Return(tc.mockFirstUserBlockSlice, nil),
				mockRepo.On("GetEmailsByUserIDs", mock.Anything, mock.Anything).Return([]string{"common@example.com"}, nil),

				mockRepo.On("GetUserIDByEmail", mock.Anything, mock.Anything).Return(tc.mockSecondUser.ID, nil),
				mockRepo.On("GetFriendsByID", mock.Anything, mock.Anything).Return(tc.mockSecondFriendSlice, nil),
				mockRepo.On("GetUserBlocksByID", mock.Anything, mock.Anything).Return(tc.mockSecondUserBlockSlice, nil),
				mockRepo.On("GetEmailsByUserIDs", mock.Anything, mock.Anything).Return([]string{"common@example.com"}, nil),
			}
			friendController := NewFriendController(mockRepo)
			handler := http.HandlerFunc(friendController.GetCommonFriends)
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if tc.expError != nil {
				require.EqualError(t, tc.expError, rr.Body.String())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, rr.Body.String())
			}
		})
	}
}
func TestControllers_CreateSubcription(t *testing.T) {
	tcs := map[string]struct {
		input             string
		expResult         string
		expError          error
		mockRequestorUser models.User
		mockTargetUser    models.User
	}{
		"success with an input": {
			input:             `{"requestor": "andy@example.com","target": "lisa@example.com"}`,
			mockRequestorUser: models.User{ID: 100, Name: "John", Email: "john@example.com"},
			mockTargetUser:    models.User{ID: 101, Name: "Andy", Email: "andy@example.com"},
			expResult:         `{"success":true}`,
		},
		"failed with an unknow format input": {
			input:    `{}`,
			expError: errors.New(`{"message":"Request body is empty","success":false}`),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/v1/subscription", bytes.NewBuffer([]byte(tc.input)))
			require.NoError(t, err)

			var mockRepo SpecRepo
			mockRepo.ExpectedCalls = []*mock.Call{
				mockRepo.On("GetUserIDByEmail", mock.Anything, mock.Anything).Return(tc.mockRequestorUser.ID, nil),
				mockRepo.On("GetUserIDByEmail", mock.Anything, mock.Anything).Return(tc.mockTargetUser.ID, nil),
				mockRepo.On("IsSubscribedFriend", mock.Anything, mock.Anything, mock.Anything).Return(false, nil),
				mockRepo.On("IsBlockedFriend", mock.Anything, mock.Anything, mock.Anything).Return(false, nil),
				mockRepo.On("CreateSubscription", mock.Anything, mock.Anything, mock.Anything).Return(nil),
			}
			friendController := NewFriendController(mockRepo)
			handler := http.HandlerFunc(friendController.CreateSubcription)
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if tc.expError != nil {
				require.EqualError(t, tc.expError, rr.Body.String())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, rr.Body.String())
			}
		})
	}
}
func TestControllers_CreateUserBlocks(t *testing.T) {
	tcs := map[string]struct {
		input             string
		expResult         string
		expError          error
		mockRequestorUser models.User
		mockTargetUser    models.User
	}{
		"success with an input": {
			input:             `{"requestor": "andy@example.com","target": "lisa@example.com"}`,
			mockRequestorUser: models.User{ID: 100, Name: "John", Email: "john@example.com"},
			mockTargetUser:    models.User{ID: 101, Name: "Andy", Email: "andy@example.com"},
			expResult:         `{"success":true}`,
		},
		"failed with an unknow format input": {
			input:    `{}`,
			expError: errors.New(`{"message":"Request body is empty","success":false}`),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/v1/blocking", bytes.NewBuffer([]byte(tc.input)))
			require.NoError(t, err)

			var mockRepo SpecRepo
			mockRepo.ExpectedCalls = []*mock.Call{
				mockRepo.On("GetUserIDByEmail", mock.Anything, mock.Anything).Return(tc.mockRequestorUser.ID, nil),
				mockRepo.On("GetUserIDByEmail", mock.Anything, mock.Anything).Return(tc.mockTargetUser.ID, nil),
				mockRepo.On("IsBlockedFriend", mock.Anything, mock.Anything, mock.Anything).Return(false, nil),
				mockRepo.On("CreateUserBlock", mock.Anything, mock.Anything, mock.Anything).Return(nil),
			}
			friendController := NewFriendController(mockRepo)
			handler := http.HandlerFunc(friendController.CreateUserBlock)
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if tc.expError != nil {
				require.EqualError(t, tc.expError, rr.Body.String())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, rr.Body.String())
			}
		})
	}
}
func TestControllers_GetRecipientEmails(t *testing.T) {
	tcs := map[string]struct {
		input         string
		expResult     string
		expError      error
		mockUser      models.User
		mockRecipient models.User
	}{
		"success with an input": {
			input:         `{"sender": "andy@example.com","text": "Hello World! kate@example.com"}`,
			mockUser:      models.User{ID: 100, Name: "John", Email: "john@example.com"},
			mockRecipient: models.User{ID: 103, Name: "Lisa", Email: "lisa@example.com"},
			expResult:     `{"recipients":["lisa@example.com","kate@example.com"],"success":true}`,
		},
		"failed with an unknow format input": {
			input:    `{}`,
			expError: errors.New(`{"message":"Request body is empty","success":false}`),
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/v1/recipients", bytes.NewBuffer([]byte(tc.input)))
			require.NoError(t, err)

			var mockRepo SpecRepo
			mockRepo.ExpectedCalls = []*mock.Call{
				mockRepo.On("GetUserIDByEmail", mock.Anything, mock.Anything).Return(tc.mockUser.ID, nil),
				mockRepo.On("GetRecipientEmails", mock.Anything, mock.Anything).Return([]models.User{tc.mockRecipient}, nil),
			}
			friendController := NewFriendController(mockRepo)
			handler := http.HandlerFunc(friendController.GetRecipientEmails)
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if tc.expError != nil {
				require.EqualError(t, tc.expError, rr.Body.String())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expResult, rr.Body.String())
			}
		})
	}
}
