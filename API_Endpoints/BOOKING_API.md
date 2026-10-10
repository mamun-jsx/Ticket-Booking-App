# Booking API Documentation

**Base URL:** `http://localhost:4001`

**_ Booking Related _**

> All booking endpoints require a valid JWT token in the `Authorization` header.
> `Authorization: Bearer <token>`

---

### 1. Create Booking

**Endpoint:** `POST /api/v1/booking/create`

#### Request Body

```json
{
  "event_id": 1,
  "quantity": 2
}
```

#### Response (201 Created)

```json
{
  "id": 1,
  "user_id": 3,
  "event_id": 1,
  "quantity": 2,
  "total_price": 3000,
  "status": "confirmed",
  "booking_code": "GT-a1b2c3d4",
  "created_at": "2026-10-10T11:00:00.000000+06:00"
}
```

#### Error Responses

```json
// 401 Unauthorized — missing or invalid token
{
  "code": 401,
  "message": "Unauthorize"
}

// 400 Bad Request — invalid body
{
  "code": 400,
  "message": "Invalid request body",
  "details": "..."
}

// 404 Not Found — event does not exist
{
  "code": 404,
  "message": "Event Not Found",
  "details": "Event Not found"
}

// 500 Internal Server Error — not enough tickets
{
  "code": 500,
  "message": "Internal Server error",
  "details": "booking not found"
}
```

---

### 2. Get My Bookings

**Endpoint:** `GET /api/v1/booking`

#### Response (200 OK)

```json
[
  {
    "id": 1,
    "user_id": 3,
    "event_id": 1,
    "quantity": 2,
    "total_price": 3000,
    "status": "confirmed",
    "booking_code": "GT-a1b2c3d4",
    "created_at": "2026-10-10T11:00:00.000000+06:00"
  },
  {
    "id": 2,
    "user_id": 3,
    "event_id": 2,
    "quantity": 1,
    "total_price": 2000,
    "status": "cancelled",
    "booking_code": "GT-e5f6g7h8",
    "created_at": "2026-10-10T12:00:00.000000+06:00"
  }
]
```

---

### 3. Cancel Booking

**Endpoint:** `PUT /api/v1/booking/cancel/:id`

#### Response (200 OK)

```json
{
  "message": "Booking cancelled successfully"
}
```

#### Error Responses

```json
// 404 Not Found — booking does not exist
{
  "code": 404,
  "message": "Booking Not Found",
  "details": "booking not found"
}

// 403 Forbidden — booking belongs to another user
{
  "code": 403,
  "message": "Forbidden",
  "details": "you are not authorized to do that"
}

// 409 Conflict — booking already cancelled
{
  "code": 409,
  "message": "Already Cancelled",
  "details": "booking already canceled"
}
```
