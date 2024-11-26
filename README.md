# golang-jwt-auth

## API Spec

### SignIn

Endpoint: POST /api/auth/signin

Request Body:

```json
{
  "first_name": "test",
  "last_name": "1",
  "email": "test@mail.com",
  "password": "password"
}
```

Response Body (Success):

```json
{
  "data": {
    "id": 1,
    "first_name": "test",
    "last_name": "1",
    "email": "test@mail.com",
    "password": "$2a$10$gdoHTs80QxRCsxgTIsZaqeSCNaDYGUJ4HSIzj080/TpBgVYwpb97G",
    "created_at": "2024-11-26T20:33:23.6600308+07:00",
    "updated_at": "2024-11-26T20:33:23.6600308+07:00"
  }
}
```

Response Body (Failed):

```json
{
    "error": "Email already use", "Internal server error"
}
```

### SignUp

Endpoint: POST /api/auth/signup

Request Body:

```json
{
  "email": "test@mail.com",
  "password": "password"
}
```

Response Body (Success):

```json
{
  "message": "Login success"
}
```

Response Body (Failed):

```json
{
    "error": "Invalid email or password", "Internal server error"
}
```
