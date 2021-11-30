package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
)

const EmailRegex = `[_A-Za-z0-9-\+]+(\.[_A-Za-z0-9-]+)*@[A-Za-z0-9-]+(\.[A-Za-z0-9]+)*(\.[A-Za-z]{2,})`

func isValidEmail(email string) (bool, error) {
	isValid, err := regexp.MatchString(EmailRegex, email)
	if err != nil || !isValid {
		return false, err
	}
	return true, nil
}

func (_self FriendRequest) Validate() error {
	if _self.Emails == nil && len(_self.Emails) == 0 {
		return ErrBodyRequestEmpty
	}
	if len(_self.Emails) != 2 {
		return ErrNumberOfEmail
	}
	if _self.Emails[0] == _self.Emails[1] {
		return ErrDifferentEmail
	}
	isValidUserEmail, err := isValidEmail(_self.Emails[0])
	if !isValidUserEmail || err != nil {
		return errors.New(_self.Emails[0] + " invalid format (ex: \"andy@example.com\")")
	}
	isValidFriendEmail, err := isValidEmail(_self.Emails[1])
	if !isValidFriendEmail || err != nil {
		return errors.New(_self.Emails[1] + " invalid format (ex: \"andy@example.com\")")
	}
	return nil
}

func (_self UserRequest) Validate() error {
	if _self.Email == "" {
		return ErrBodyRequestEmpty
	}
	isValidEmail, err := isValidEmail(_self.Email)
	if !isValidEmail || err != nil {
		return errors.New(_self.Email + " invalid format (ex: \"andy@example.com\")")
	}
	return nil
}

func (_self RequestorRequest) Validate() error {
	if _self.Requestor == "" && _self.Target == "" {
		return ErrBodyRequestEmpty
	}

	if _self.Requestor == "" {
		return ErrRequestorFieldInvalid
	}

	if _self.Target == "" {
		return ErrTargetFieldInvalid
	}

	if _self.Target == _self.Requestor {
		return ErrDifferentEmail
	}

	isValidRequestEmail, requestErr := isValidEmail(_self.Requestor)
	if !isValidRequestEmail || requestErr != nil {
		return errors.New(_self.Requestor + "invalid format (ex: \"andy@example.com\")")
	}

	isValidTargetEmail, targetErr := isValidEmail(_self.Target)
	if !isValidTargetEmail || targetErr != nil {
		return errors.New(_self.Target + " invalid format (ex: \"andy@example.com\")")
	}
	return nil
}

func (_self RecipientsRequest) Validate() error {
	if _self.Sender == "" && _self.Text == "" {
		return ErrBodyRequestEmpty
	}

	if _self.Sender == "" {
		return ErrSenderFieldInvalid
	}
	if _self.Text == "" {
		return ErrTextFieldInvalid
	}
	isValidEmail, err := isValidEmail(_self.Sender)

	if !isValidEmail || err != nil {
		return errors.New(_self.Sender + " invalid format (ex: \"andy@example.com\")")
	}
	return nil
}

func GetMentionedEmailFromText(text string) []string {
	regex := regexp.MustCompile(EmailRegex)

	emailChain := regex.FindAllString(text, -1)
	email := make([]string, len(emailChain))
	for index, emailCharacter := range emailChain {
		email[index] = emailCharacter
	}
	return email
}

func MsgOK() map[string]interface{} {
	return map[string]interface{}{"success": true}
}

func MsgError(err error) map[string]interface{} {
	return map[string]interface{}{"message": err.Error(), "success": false}
}

func Message(status bool, msg string) map[string]interface{} {
	return map[string]interface{}{"message": msg, "success": status}
}

func MsgGetFriendsOk(friends []string, count int) interface{} {
	return map[string]interface{}{"count": count, "friends": friends, "success": true}
}

func MsgGetEmailReceiversOk(emails []string) interface{} {
	return map[string]interface{}{"recipients": emails, "success": true}
}

func Respond(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.WriteHeader(statusCode)
	w.Write(response)
}
