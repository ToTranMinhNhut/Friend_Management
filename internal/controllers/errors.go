package controllers

import "errors"

var (
	ErrBodyRequestInvalid    = errors.New("Body request invalid format")
	ErrExistedFriendship     = errors.New("The friend relationship has been existed")
	ErrExistedBlockedUser    = errors.New("The users have blocked each other")
	ErrExistedSubscription   = errors.New("The users have subscribed each other")
	ErrCreatedFriendship     = errors.New("Users cannot be created a new friendship")
	ErrBodyRequestEmpty      = errors.New("Request body is empty")
	ErrNumberOfEmail         = errors.New("Number of email addresses must be 2")
	ErrDifferentEmail        = errors.New("Two email addresses must be different")
	ErrRequestorFieldInvalid = errors.New("Requestor field invalid format")
	ErrTargetFieldInvalid    = errors.New("Target field invalid format")
	ErrSenderFieldInvalid    = errors.New("Sender field invalid format")
	ErrTextFieldInvalid      = errors.New("Text field invalid format")
)
