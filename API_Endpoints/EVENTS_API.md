# Ticket Booking App - API Documentation

**Base URL:** `http://localhost:4001`

**_ Event Related _**

---

### 1. Create Event

**Endpoint:** `POST /api/v1/event/create`

### Request Body

```json
{
  "title": "Advanced Machine Learning with Python",
  "description": "Deep dive into data preprocessing, model training with Scikit-Learn, and deploying ML models as REST APIs.",
  "location": "Chittagong, Bangladesh",
  "start_at": "2026-11-20T10:00:00Z",
  "total_tickets": 50,
  "price": 2500
}
```

### Response (201 Created)

```json
{
  "id": 3,
  "title": "Advanced Machine Learning with Python",
  "description": "Deep dive into data preprocessing, model training with Scikit-Learn, and deploying ML models as REST APIs.",
  "location": "Chittagong, Bangladesh",
  "start_at": "2026-11-20T10:00:00Z",
  "total_tickets": 50,
  "available_tickets": 50,
  "price": 2500,
  "created_at": "2026-10-03T23:58:24.552898+06:00",
  "updated_at": "2026-10-03T23:58:24.552898+06:00"
}
```

### 2. Get All Events

**Endpoint:** `GET /api/v1/event`

```json
[
  {
    "id": 1,
    "title": "Golang Backend Architecture Workshop",
    "description": "Learn how to build production-ready APIs with Fiber, GORM, and clean architecture principles.",
    "location": "Dhaka, Bangladesh",
    "start_at": "2026-10-15T18:30:00Z",
    "total_tickets": 100,
    "available_tickets": 100,
    "price": 1500,
    "created_at": "2026-10-03T17:52:09.138217Z",
    "updated_at": "2026-10-03T17:52:09.138217Z"
  },
  {
    "id": 2,
    "title": "Syber Security System",
    "description": "Learn how to build production-ready APIs with Fiber, GORM, and clean architecture principles.",
    "location": "Dhaka, Bangladesh",
    "start_at": "2026-10-15T18:30:00Z",
    "total_tickets": 100,
    "available_tickets": 100,
    "price": 2000,
    "created_at": "2026-10-03T17:52:36.13283Z",
    "updated_at": "2026-10-03T17:52:36.13283Z"
  }
]
```

### 3. Get Event By ID || get a single events

**Endpoint:** `GET /api/v1/event/:id`

```json
{
  "id": 1,
  "title": "Golang Backend Architecture Workshop",
  "description": "Learn how to build production-ready APIs with Fiber, GORM, and clean architecture principles.",
  "location": "Dhaka, Bangladesh",
  "start_at": "2026-10-15T18:30:00Z",
  "total_tickets": 100,
  "available_tickets": 100,
  "price": 1500,
  "created_at": "2026-10-03T17:52:09.138217Z",
  "updated_at": "2026-10-03T17:52:09.138217Z"
}
```

