# toggl-test-signer

## How to Run?

### Native

**Requirements**
- Go 1.21.4
- MySQL 8 running on port 13316.

Run `make run` and then navigate to <http://localhost:8080>

### Docker

Run `make docker` and then navigate to <http://localhost:8080>

## How to Use?

#### `POST    /api/v1/users/sign-answers`

1. Create a JWT (e.g. <https://www.javainuse.com/jwtgenerator>) with the payload containing `userID`.
Use the value of `JWT_SECRET` from `.env` file as the secret.

```json
{
  "userID": "123"
}
```

2. Use this JWT as `Authorization: BEARER eyJhbGciOiJIUzUxMiJ9.eyJ1c2VySUQiOjEyM30.LeNw5iobDHBW1rrdsW5-P38qfK0b6N2BOF-rfZKCXXMEVqqhGrAw7_rwVpFXVs6p2_0Y-sjMsHfG6FyErKqnsQ`
3. Make request.

```json
{
    "answers": [
        {
            "questionID": 1,
            "answer": "answer 1"
        },
        {
            "questionID": 2,
            "answer": "answer 2"
        }
    ]
}
```

Response
```json
{
    "userID": 123,
    "signature": "4e1fce5c292feabce21c053f18a3ea41b4b684127069bce0c84a33beaa2ec206"
}
```

4. Error scenarios

```json
{
  "status": "bad request",
  "status_code": 400,
  "error": "userID cannot be casted into int"
}
```

```json
{
  "status": "bad request",
  "status_code": 400,
  "error": "userID is not found in JWT payload"
}
```

```json
{
  "status": "bad request",
  "status_code": 400,
  "error": "retrieving JWT from context failed"
}
```

#### `POST    /api/v1/admin/verify-signature`

1. Use the signature from previous step.
2. Make request.

````json
{
    "userID": 123,
    "signature": "4e1fce5c292feabce21c053f18a3ea41b4b684127069bce0c84a33beaa2ec206"
}
````

3. Check response

```json
{
  "answers": [
    {
      "questionID": 1,
      "answer": "answer 1"
    },
    {
      "questionID": 2,
      "answer": "answer 2"
    }
  ],
  "signedAt": "2023-12-12T15:45:41Z"
}
```

4. Error scenarios

```json
{
    "status": "not found",
    "status_code": 404,
    "error": "record not found"
}
```
## Design Considerations
Due to the time constraints,
- The design of questions and answers tables is omitted. The signing is done by saving QAndA as a text column.
- `uesr` table is not defined and there's foreign key with `test` table.

## Testing

Didn't have time to setup integration tests.
The test cases to be covered:

- POST /api/v1/users/sign-answers
  - 400 - JWT generated with the wrong secret
  - 400 - JWT has invalid payload
  - 400 - JWT `userID` is not a number
  - 400 - Request is invalid
  - 500 - unknown internal errors
  - 200 - Success

- POST /api/v1/admin/verify-signature
  - 400 - Request is invalid
  - 404 - cannot find the test with the userID and signature
  - 200 - OK