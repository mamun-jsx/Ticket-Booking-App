# Ticket Booking App - API Documentation

**Base URL:** `http://localhost:4001`

---

## 1. Register User

**Endpoint:** `POST /api/v1/auth/register`

### Request Body

```json
{
  "name": "Mamun Hossain",
  "email": "mamun@example.com",
  "password": "secretpassword123"
}
```

### Response (201 Created)

```json
{
  "id": 1,
  "name": "Mamun Hossain",
  "email": "mamun@example.com",
  "created_at": "2026-10-01 17:09:26.3078584 +0600 +06 m=+0.001"
}
```

---

## 2. Login User

**Endpoint:** `POST /api/v1/auth/login`

### Request Body

```json
{
  "email": "mamun@example.com",
  "password": "secretpassword123"
}
```

### Response (200 OK)

```json
{
  "id": 1,
  "name": "Mamun Hossain",
  "email": "mamun@example.com",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "created_at": "2026-10-01 17:09:26.3078584 +0600 +06 m=+0.001"
}
```

### Error Responses

#### 400 Bad Request
```json
{
  "code": 400,
  "message": "Invalid request body Input",
  "details": "..."
}
```

#### 401 Unauthorized
```json
{
  "code": 401,
  "message": "Authentication failed",
  "details": "Invalid Access ID password Not Match"
}
```

