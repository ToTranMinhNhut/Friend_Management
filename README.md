# S3_FriendManagementAPI_NhutTo
CURL app with Go and Postgresql

## Backend Setup
- Make sure that your Docker is up and running
- From folder `S3_FriendManagementAPI_NhutTo/`, run `make setup`

## Run app
- Start the server: `make run`
- Server running on: `http://localhost:8080`

## Run test
- Run command `make test`

## Teardown
```
make teardown
```

## API information
1 - Get users
- GET: http://localhost:8080/v1/users
- Parameter request: none
- Success with status code: 200 OK
```
{
    "count": 5,
    "success": true,
    "users": [
        "john@example.com",
        "andy@example.com",
        "common@example.com",
        "lisa@example.com",
        "kate@example.com"
    ]
}
```

2 - Create friend
- POST: http://localhost:8080/v1/friends
- Parameter request:
```
{ 
    "friends": [
        "andy@example.com",
        "john@example.com"
    ]
}
```

- Success with status code: 200 OK
```
{
    "success": true
}
```

3 - List Friends
- GET: http://localhost:8080/v1/friends
- Parameter request:
```
{
    "Email":"andy@example.com"
}
```

- Success with status code: 200 OK

```
{
    "count": 1,
    "friends": [
        "lisa@example.com"
    ],
    "success": true
}
```

4 - Get common friends
- GET: http://localhost:8080/v1/commonFriends
- Parameter request:
```
{ 
    "friends": [
        "andy@example.com",
        "john@example.com"
    ]
}
```

- Success with status code: 200 OK
```
{
    "count": 1,
    "friends": [
        "common@example.com"
    ],
    "success": true
}
```

5 - Create subscription
- POST: http://localhost:8080/v1/subscription
- Parameter request:
```
{
  "requestor": "andy@example.com",
  "target": "lisa@example.com"
}
```

- Success with status code: 200 OK
```
{
    "success": true
}
```

6 - Create user block
- POST: http://localhost:8080/v1/blocking
- Parameter request:
```
{
    "requestor": "common@example.com",
    "target": "kate@example.com"
}
```

- Success with status code: 200 OK
```
{
    "success": true
}
```

7 - Get Recipients
- GET: http://localhost:8080 /v1/recipients
- Parameter request:
```
{
    "sender": "lisa@example.com",
    "text": "Hello World! kate@example.com"
}
```

- Success with status code: 200 OK
```
{
    "recipients": [
        "common@example.com",
        "kate@example.com"
    ],
    "success": true
}
```

## Unit Test results

?   	github.com/ToTranMinhNhut/S3_FriendManagementAPI_NhutTo	[no test files]
?   	github.com/ToTranMinhNhut/S3_FriendManagementAPI_NhutTo/config	[no test files]
=== RUN   TestControllers_GetFriends
=== RUN   TestControllers_GetFriends/success_with_an_input
=== RUN   TestControllers_GetFriends/failed_with_an_unknow_format_input
--- PASS: TestControllers_GetFriends (0.00s)
    --- PASS: TestControllers_GetFriends/success_with_an_input (0.00s)
    --- PASS: TestControllers_GetFriends/failed_with_an_unknow_format_input (0.00s)
=== RUN   TestControllers_CreateFriends
=== RUN   TestControllers_CreateFriends/success_with_an_input
=== RUN   TestControllers_CreateFriends/failed_with_an_unknow_format_input
--- PASS: TestControllers_CreateFriends (0.00s)
    --- PASS: TestControllers_CreateFriends/success_with_an_input (0.00s)
    --- PASS: TestControllers_CreateFriends/failed_with_an_unknow_format_input (0.00s)
=== RUN   TestControllers_GetCommonFriends
=== RUN   TestControllers_GetCommonFriends/success_with_an_input
=== RUN   TestControllers_GetCommonFriends/failed_with_an_unknow_format_input
--- PASS: TestControllers_GetCommonFriends (0.00s)
    --- PASS: TestControllers_GetCommonFriends/success_with_an_input (0.00s)
    --- PASS: TestControllers_GetCommonFriends/failed_with_an_unknow_format_input (0.00s)
=== RUN   TestControllers_CreateSubcription
=== RUN   TestControllers_CreateSubcription/success_with_an_input
=== RUN   TestControllers_CreateSubcription/failed_with_an_unknow_format_input
--- PASS: TestControllers_CreateSubcription (0.00s)
    --- PASS: TestControllers_CreateSubcription/success_with_an_input (0.00s)
    --- PASS: TestControllers_CreateSubcription/failed_with_an_unknow_format_input (0.00s)
=== RUN   TestControllers_CreateUserBlocks
=== RUN   TestControllers_CreateUserBlocks/success_with_an_input
=== RUN   TestControllers_CreateUserBlocks/failed_with_an_unknow_format_input
--- PASS: TestControllers_CreateUserBlocks (0.00s)
    --- PASS: TestControllers_CreateUserBlocks/success_with_an_input (0.00s)
    --- PASS: TestControllers_CreateUserBlocks/failed_with_an_unknow_format_input (0.00s)
=== RUN   TestControllers_GetRecipientEmails
=== RUN   TestControllers_GetRecipientEmails/success_with_an_input
=== RUN   TestControllers_GetRecipientEmails/failed_with_an_unknow_format_input
--- PASS: TestControllers_GetRecipientEmails (0.00s)
    --- PASS: TestControllers_GetRecipientEmails/success_with_an_input (0.00s)
    --- PASS: TestControllers_GetRecipientEmails/failed_with_an_unknow_format_input (0.00s)
=== RUN   TestControllers_GetUsers
=== RUN   TestControllers_GetUsers/success_with_an_empty_input
=== RUN   TestControllers_GetUsers/failed_with_an_unknow_format_input
--- PASS: TestControllers_GetUsers (0.00s)
    --- PASS: TestControllers_GetUsers/success_with_an_empty_input (0.00s)
    --- PASS: TestControllers_GetUsers/failed_with_an_unknow_format_input (0.00s)
PASS
ok  	github.com/ToTranMinhNhut/S3_FriendManagementAPI_NhutTo/internal/controllers	0.012s
?   	github.com/ToTranMinhNhut/S3_FriendManagementAPI_NhutTo/internal/models	[no test files]
=== RUN   TestRepository_CreateFriend
=== RUN   TestRepository_CreateFriend/success_with_adding_input_of_userIds
=== RUN   TestRepository_CreateFriend/query_by_an_unknown_input_userIds
--- PASS: TestRepository_CreateFriend (0.10s)
    --- PASS: TestRepository_CreateFriend/success_with_adding_input_of_userIds (0.06s)
    --- PASS: TestRepository_CreateFriend/query_by_an_unknown_input_userIds (0.05s)
=== RUN   TestRepository_IsExistedFriend
=== RUN   TestRepository_IsExistedFriend/success_with_adding_input_of_userIds
=== RUN   TestRepository_IsExistedFriend/query_by_an_unknown_input_userIds
--- PASS: TestRepository_IsExistedFriend (0.09s)
    --- PASS: TestRepository_IsExistedFriend/success_with_adding_input_of_userIds (0.05s)
    --- PASS: TestRepository_IsExistedFriend/query_by_an_unknown_input_userIds (0.04s)
=== RUN   TestRepository_IsBlockedUser
=== RUN   TestRepository_IsBlockedUser/success_with_adding_input_of_userIds
=== RUN   TestRepository_IsBlockedUser/query_by_an_unknown_input_userIds
--- PASS: TestRepository_IsBlockedUser (0.09s)
    --- PASS: TestRepository_IsBlockedUser/success_with_adding_input_of_userIds (0.05s)
    --- PASS: TestRepository_IsBlockedUser/query_by_an_unknown_input_userIds (0.05s)
=== RUN   TestRepository_GetFriendsByID
=== RUN   TestRepository_GetFriendsByID/query_by_an_unknown_input_userId
=== RUN   TestRepository_GetFriendsByID/success_with_adding_input_of_userId
--- PASS: TestRepository_GetFriendsByID (0.09s)
    --- PASS: TestRepository_GetFriendsByID/query_by_an_unknown_input_userId (0.04s)
    --- PASS: TestRepository_GetFriendsByID/success_with_adding_input_of_userId (0.04s)
=== RUN   TestRepository_GetUserBlocksByID
=== RUN   TestRepository_GetUserBlocksByID/query_by_an_unknown_input_userId
=== RUN   TestRepository_GetUserBlocksByID/success_with_adding_input_of_userId
--- PASS: TestRepository_GetUserBlocksByID (0.07s)
    --- PASS: TestRepository_GetUserBlocksByID/query_by_an_unknown_input_userId (0.04s)
    --- PASS: TestRepository_GetUserBlocksByID/success_with_adding_input_of_userId (0.04s)
=== RUN   TestRepository_CreateSubscription
=== RUN   TestRepository_CreateSubscription/success_with_adding_input_of_userIds
=== RUN   TestRepository_CreateSubscription/query_by_an_unknown_input_userIds
--- PASS: TestRepository_CreateSubscription (0.08s)
    --- PASS: TestRepository_CreateSubscription/success_with_adding_input_of_userIds (0.04s)
    --- PASS: TestRepository_CreateSubscription/query_by_an_unknown_input_userIds (0.04s)
=== RUN   TestRepository_GetRecipientEmails
=== RUN   TestRepository_GetRecipientEmails/query_by_an_unknown_input_userId
=== RUN   TestRepository_GetRecipientEmails/success_with_adding_input_of_userId
--- PASS: TestRepository_GetRecipientEmails (0.08s)
    --- PASS: TestRepository_GetRecipientEmails/query_by_an_unknown_input_userId (0.04s)
    --- PASS: TestRepository_GetRecipientEmails/success_with_adding_input_of_userId (0.04s)
=== RUN   TestRepository_CreateUserBlock
=== RUN   TestRepository_CreateUserBlock/query_by_an_unknown_input_userIds
=== RUN   TestRepository_CreateUserBlock/success_with_adding_input_of_userIds
--- PASS: TestRepository_CreateUserBlock (0.08s)
    --- PASS: TestRepository_CreateUserBlock/query_by_an_unknown_input_userIds (0.04s)
    --- PASS: TestRepository_CreateUserBlock/success_with_adding_input_of_userIds (0.04s)
=== RUN   TestRepository_IsSubscribedFriend
=== RUN   TestRepository_IsSubscribedFriend/success_with_adding_input_of_userIds
=== RUN   TestRepository_IsSubscribedFriend/query_by_an_unknown_input_userIds
--- PASS: TestRepository_IsSubscribedFriend (0.08s)
    --- PASS: TestRepository_IsSubscribedFriend/success_with_adding_input_of_userIds (0.04s)
    --- PASS: TestRepository_IsSubscribedFriend/query_by_an_unknown_input_userIds (0.04s)
=== RUN   TestRepository_GetUserIDByEmail
=== RUN   TestRepository_GetUserIDByEmail/query_by_an_unknown_input_email
=== RUN   TestRepository_GetUserIDByEmail/success_with_adding_input_of_email
--- PASS: TestRepository_GetUserIDByEmail (0.07s)
    --- PASS: TestRepository_GetUserIDByEmail/query_by_an_unknown_input_email (0.04s)
    --- PASS: TestRepository_GetUserIDByEmail/success_with_adding_input_of_email (0.04s)
=== RUN   TestRepository_GetEmailsByUserIDs
=== RUN   TestRepository_GetEmailsByUserIDs/success_with_adding_input_of_userIds
=== RUN   TestRepository_GetEmailsByUserIDs/query_by_an_unknown_input_userIds
--- PASS: TestRepository_GetEmailsByUserIDs (0.08s)
    --- PASS: TestRepository_GetEmailsByUserIDs/success_with_adding_input_of_userIds (0.03s)
    --- PASS: TestRepository_GetEmailsByUserIDs/query_by_an_unknown_input_userIds (0.04s)
=== RUN   TestRepository_GetUsers
=== RUN   TestRepository_GetUsers/successfully_get_all_users
--- PASS: TestRepository_GetUsers (0.04s)
    --- PASS: TestRepository_GetUsers/successfully_get_all_users (0.04s)
PASS
ok  	github.com/ToTranMinhNhut/S3_FriendManagementAPI_NhutTo/internal/repository