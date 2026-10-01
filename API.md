### Create User

**Endpoint:** `POST {BASE_URL}/api/v1/users/create`

#### Request Body

```json
{
  "name": "Mamun Hossain",
  "email": "mamun@example.com",
  "password": "secretpassword123"
}
```

#### Response (201 Created)

```json
{
  "code": 0,
  "message": "User created successfully",
  "data": {
    "id": 1,
    "name": "Mamun Hossain",
    "email": "mamun@example.com",
    "created_at": "2026-10-01T17:09:26.3078584+06:00",
    "updated_at": "2026-10-01T17:09:26.3078584+06:00"
  }
}
```
