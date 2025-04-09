# 📘 User Authentication API Documentation (v1)

## 🧾 Base URL:
```
http://localhost:8080/api/v1
```

---

## 📌 User Routes

| Method | Endpoint                   | Description                | Auth Required | Status Codes |
|--------|----------------------------|----------------------------|----------------|----------------|
| POST   | `/users/register`          | Register a new user        | ❌             | 201, 400       |
| POST   | `/users/login`             | Login and get JWT token    | ❌             | 200, 401       |
| GET    | `/users/`                  | Get all users              | ✅             | 200, 401       |
| GET    | `/users/:id`               | Get user by ID             | ✅             | 200, 404, 401  |
| PUT    | `/users/:id`               | Update user by ID          | ✅             | 200, 400, 404  |
| DELETE | `/users/:id`               | Delete user by ID          | ✅             | 204, 404, 401  |

---

## 📋 Detailed API Documentation

### Register a New User
**Endpoint:** `POST /users/register`  
**Auth Required:** No  
**Description:** Creates a new user account.

#### Request Body:
```json
{
  "username": "johndoe",
  "email": "john.doe@example.com",
  "password": "StrongPassword123!",
  "firstName": "John",
  "lastName": "Doe"
}
```

#### Successful Response (201 Created):
```json
{
  "status": "success",
  "message": "User registered successfully",
  "data": {
    "id": "60d5f3a7e5f39e001feebb50",
    "username": "johndoe",
    "email": "john.doe@example.com",
    "firstName": "John",
    "lastName": "Doe",
    "createdAt": "2023-06-25T12:30:45Z"
  }
}
```

#### Error Response (400 Bad Request):
```json
{
  "status": "error",
  "message": "Validation failed",
  "errors": [
    {
      "field": "email",
      "message": "Invalid email format"
    },
    {
      "field": "password",
      "message": "Password must be at least 8 characters with at least one uppercase letter, one number, and one special character"
    }
  ]
}
```

---

### User Login
**Endpoint:** `POST /users/login`  
**Auth Required:** No  
**Description:** Authenticates a user and returns a JWT token.

#### Request Body:
```json
{
  "email": "john.doe@example.com",
  "password": "StrongPassword123!"
}
```

#### Successful Response (200 OK):
```json
{
  "status": "success",
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": "60d5f3a7e5f39e001feebb50",
      "username": "johndoe",
      "email": "john.doe@example.com",
      "firstName": "John",
      "lastName": "Doe"
    }
  }
}
```

#### Error Response (401 Unauthorized):
```json
{
  "status": "error",
  "message": "Invalid email or password"
}
```

---

### Get All Users
**Endpoint:** `GET /users`  
**Auth Required:** Yes (Admin only)  
**Description:** Retrieves a list of all users.

#### Headers:
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

#### Query Parameters:
```
page: 1 (default)
limit: 10 (default)
sort: createdAt (default)
order: desc (default)
```

#### Successful Response (200 OK):
```json
{
  "status": "success",
  "data": {
    "users": [
      {
        "id": "60d5f3a7e5f39e001feebb50",
        "username": "johndoe",
        "email": "john.doe@example.com",
        "firstName": "John",
        "lastName": "Doe",
        "createdAt": "2023-06-25T12:30:45Z"
      },
      {
        "id": "60d5f3a7e5f39e001feebb51",
        "username": "janedoe",
        "email": "jane.doe@example.com",
        "firstName": "Jane",
        "lastName": "Doe",
        "createdAt": "2023-06-26T10:15:30Z"
      }
    ],
    "pagination": {
      "totalCount": 25,
      "totalPages": 3,
      "currentPage": 1,
      "limit": 10
    }
  }
}
```

#### Error Response (401 Unauthorized):
```json
{
  "status": "error",
  "message": "Authentication required"
}
```

---

### Get User by ID
**Endpoint:** `GET /users/:id`  
**Auth Required:** Yes  
**Description:** Retrieves a specific user by ID.

#### Headers:
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

#### Successful Response (200 OK):
```json
{
  "status": "success",
  "data": {
    "user": {
      "id": "60d5f3a7e5f39e001feebb50",
      "username": "johndoe",
      "email": "john.doe@example.com",
      "firstName": "John",
      "lastName": "Doe",
      "createdAt": "2023-06-25T12:30:45Z"
    }
  }
}
```

#### Error Response (404 Not Found):
```json
{
  "status": "error",
  "message": "User not found"
}
```

#### Error Response (401 Unauthorized):
```json
{
  "status": "error",
  "message": "Authentication required"
}
```

---

### Update User by ID
**Endpoint:** `PUT /users/:id`  
**Auth Required:** Yes (User can update own profile, Admin can update any)  
**Description:** Updates a user's profile information.

#### Headers:
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

#### Request Body:
```json
{
  "firstName": "Johnny",
  "lastName": "Doe",
  "email": "johnny.doe@example.com"
}
```

#### Successful Response (200 OK):
```json
{
  "status": "success",
  "message": "User updated successfully",
  "data": {
    "user": {
      "id": "60d5f3a7e5f39e001feebb50",
      "username": "johndoe",
      "email": "johnny.doe@example.com",
      "firstName": "Johnny",
      "lastName": "Doe",
      "updatedAt": "2023-06-27T09:45:12Z"
    }
  }
}
```

#### Error Response (400 Bad Request):
```json
{
  "status": "error",
  "message": "Validation failed",
  "errors": [
    {
      "field": "email",
      "message": "Invalid email format"
    }
  ]
}
```

#### Error Response (404 Not Found):
```json
{
  "status": "error",
  "message": "User not found"
}
```

#### Error Response (401 Unauthorized):
```json
{
  "status": "error",
  "message": "Authentication required"
}
```

---

### Delete User by ID
**Endpoint:** `DELETE /users/:id`  
**Auth Required:** Yes (User can delete own account, Admin can delete any)  
**Description:** Deletes a user account.

#### Headers:
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

#### Successful Response (204 No Content):
```
No content returned on successful deletion
```

#### Error Response (404 Not Found):
```json
{
  "status": "error",
  "message": "User not found"
}
```

#### Error Response (401 Unauthorized):
```json
{
  "status": "error",
  "message": "Authentication required"
}
```

---

## 🔐 Authentication

Authentication is handled via JWT tokens. After a successful login, the token should be:
- Included in the Authorization header as `Bearer <token>` for all protected endpoints
- Stored securely on the client side
- Refreshed before expiration (token validity: 24 hours)

## 🛡️ Security Considerations

- All passwords are hashed using bcrypt before storing
- API uses HTTPS for all communications
- JWT tokens are encrypted and signed
- Rate limiting is applied to prevent brute force attacks
- Input validation is applied to all endpoints