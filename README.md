# Ticket Booking App

A backend system built with Go and the Fiber framework that allows users to book tickets for events.

## Local Setup

### 1. Prerequisites
- [Go](https://go.dev/dl/) (version 1.22+ or compatible)
- [PostgreSQL](https://www.postgresql.org/) database

### 2. Clone the Repository
```bash
git clone https://github.com/mamun-jsx/Ticket-Booking-App.git
cd Ticket-Booking-App
```

### 3. Environment Configuration
Create a `.env` file from the provided example:
```bash
cp .env.example .env
```
Update the `.env` file with your database credentials and port.

### 4. Install Dependencies
```bash
go mod tidy
```

### 5. Run the Server
```bash
go run cmd/main.go
```
The server will start at `http://localhost:8080`.
