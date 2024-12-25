# USER API SPEC

## Register User

Endpoint: POST /api/v1/register

Request Body :

```json
{
  "username": "wign",
  "email": "wign@gmail.com",
  "password": "secret"
}
```

Response Body (Success) :

```json
{
    "message": "User created successfully"
}
```

Response Body (Failed) :

```json
{
    "errors": "user already exist"
}
```

## Login User

Endpoint: POST /api/v1/login

```json
{
  "username": "test",
  "password": "secret"
}
```

Response Body (Success):

```json
{
 "data":"token"
}
```

Response Body (Failed) :

```json
{
"errors": "Credentials error"
}
```

## Get User

Endpoint: GET /api/v1/users/:id

Request Header :

- authorization : TOKEN

Response Body (Success):

```json
{
    "id": "cm2o3fwes00039hqywupqf8j1",
    "username": "wignn",
    "profile_picture": null,
    "email": "wignn@gmail.com",
    "created_at": "2024-10-25T02:09:03.700Z",
    "created_at": "2705d068-a9a3-43e2-b787-6cc1a7ec7de1",
}
```

Response Body (Failed) :

```json
"erors":"unauthorized,..."
```

## Update User

Endpoint : PUT /api/v1/users/id

Request Header :

- authorization : TOKEN

Request Body :

```json
{
  "username": "wign66",/optional
  "email":"test@gmail.com",/optional
  "profile_picture":"file://exemple.com",/optional
}
```

Response Body (Success) :

```json
{    
   "message": "User profile updated successfully"
}
```

Response Body (Failed) :

```json
"erors":"Unauthorized...."
```

## Verification Password

Endpoint: PATCH /api/v1/verify-email/:id

Request Body :

```json
{

}
```

Response Body (Success) :

```json
{
 "message": "Email sent successfully"
}
```

## Reset Password

Endpoint: PUT /api/v1/users/:id/reset-password

Request Body :

```json
{
  "password": "newpassword",
  "token":"akjdadg-dfgfdgfdg-hfnsas-faghqb5tnm-adadrrw",
}
```

Response Body (Success) :

```json
{
  "message": "reset successful"
}
```

