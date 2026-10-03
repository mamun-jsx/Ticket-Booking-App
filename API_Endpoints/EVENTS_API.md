## 1. Register User

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
