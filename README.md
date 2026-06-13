# User with DOB and Calculated Age API

## Overview

This project is a RESTful API developed using Go and Fiber for managing user information. Each user record contains a name and date of birth (DOB). The API calculates the user's age dynamically whenever user details are retrieved, ensuring that age is always accurate without storing it in the database.

The project follows a clean architecture approach with separate layers for handlers, services, repositories, and database access generated using SQLC.

## Technologies Used

* Go
* Fiber
* MySQL
* SQLC
* Uber Zap Logger

## Project Structure

```text
User-with-DOB-and-Calculated-Age/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── config/
│
├── db/
│   ├── migrations/
│   ├── queries/
│   └── sqlc/
│
├── internal/
│   ├── handler/
│   ├── repository/
│   ├── service/
│   ├── routes/
│   ├── middleware/
│   ├── models/
│   └── logger/
│
├── go.mod
├── go.sum
├── sqlc.yaml
└── README.md
```

## Database Setup

Create a database named `userdb` and run the following query:

```sql
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    dob DATE NOT NULL
);
```

## Installation

Clone the repository:

```bash
git clone https://github.com/mamilladruva0007/User-with-DOB-and-Calculated-Age.git
cd User-with-DOB-and-Calculated-Age
```

Install dependencies:

```bash
go mod tidy
```

Generate SQLC files:

```bash
sqlc generate
```

## Running the Application

Start the server using:

```bash
go run cmd/server/main.go
```

The application will run on:

```text
http://localhost:3000
```

## API Endpoints

### Create User

**POST** `/api/users`

Request:

```json
{
  "name": "Alice",
  "dob": "1990-05-10"
}
```

Response:

```json
{
  "name": "Alice",
  "dob": "1990-05-10"
}
```

### Get User by ID

**GET** `/api/users/{id}`

Response:

```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 36
}
```

### List All Users

**GET** `/api/users`

Response:

```json
[
  {
    "id": 1,
    "name": "Alice",
    "dob": "1990-05-10",
    "age": 36
  }
]
```

### Update User

**PUT** `/api/users/{id}`

Request:

```json
{
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

Response:

```json
{
  "id": 1,
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

### Delete User

**DELETE** `/api/users/{id}`

Response:

```text
204 No Content
```

## Age Calculation

The user's age is not stored in the database. Instead, it is calculated dynamically from the date of birth whenever user information is requested.

This approach keeps the data accurate and eliminates the need to update ages manually every year.

## Sample Testing

The API was tested using PowerShell and supports all CRUD operations:

* User creation
* User retrieval
* User update
* User deletion
* Listing all users with calculated age

## Future Improvements

* Request validation using go-playground/validator
* Docker support
* Pagination for listing users
* Unit tests for age calculation
* Request ID middleware
* Request duration logging middleware

## Author

**Druva Chandra**

Backend Development Task – User with DOB and Calculated Age API
