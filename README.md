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
1 - Create friend
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

2 - List Friends
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

3 - Get common friends
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

4 - Create subscription
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

5 - Create user block
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

6 - Get Recipients
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

