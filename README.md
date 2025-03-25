Stock Broker Application

A backend service for user authentication and management in a stock broker platform.

---

FEATURES

- User Signup with validations and hashed passwords
- Data Validation for email, phone number, PAN, and password
- Database Integration using PostgreSQL and GORM
- Automated Migrations handled in sqlSetup.go
- Error Handling with detailed responses

---

TECHNOLOGIES USED

- Golang (Gin, GORM)
- PostgreSQL
- Validator for Input Validation

---

SETUP INSTRUCTIONS

Prerequisites

- Install Go (1.20+)
- Install PostgreSQL

Clone the Repository

git clone https://github.com/HarmanjotSingh-Coditas/stock_broker_application.git
cd stock_broker_application

Configure Environment Variables

Modify config.yaml (or .env if used) to set up database credentials:

database:
  host: "localhost"
  port: 5432
  user: "your_user"
  password: "your_password"
  dbname: "stock_broker_db"

Run Migrations

Migrations are automatically handled in sqlSetup.go, so just run the main file.

Start the Server

go run main.go

The server will start on localhost:8080

---

API ENDPOINTS

User Signup

POST /Sign

Registers a new user with validation checks.

Request Body

{
  "name": "John Doe",
  "password": "Strong@123",
  "confirmpassword": "Strong@123",
  "email": "johndoe@example.com",
  "phoneno": 9876543210,
  "pan": "ABCDE1234F"
}

Response

{
  "message": "User created successfully"
}

Error Handling

The application provides detailed validation errors. Example:

{
  "errors": [
    "email is required",
    "password must be at least 8 characters"
  ]
}

---

ADDITIONAL NOTES

- The application uses hashed passwords for security.
- Duplicate username/email results in a conflict error.
- Phone numbers are stored as BIGINT to avoid string conversion issues.

---

RUNNING TESTS

You can test the API using Postman or cURL.

Test with Postman

1. Open Postman
2. Select POST request
3. Enter the URL:
   http://localhost:8080/Sign
4. Go to Body -> raw -> JSON
5. Paste the following:
   {
     "name": "Test User",
     "password": "Test@1234",
     "confirmpassword": "Test@1234",
     "email": "testuser@example.com",
     "phoneno": 9876543211,
     "pan": "XYZAB5678L"
   }
6. Click Send
   You should receive:
   {
     "message": "User created successfully"
   }

Test with cURL

Successful Request

curl -X POST "http://localhost:8080/Sign" \
-H "Content-Type: application/json" \
-d '{
  "name": "Test User",
  "password": "Test@1234",
  "confirmpassword": "Test@1234",
  "email": "testuser@example.com",
  "phoneno": 9876543211,
  "pan": "XYZAB5678L"
}'

Error Case (Missing Password)

curl -X POST "http://localhost:8080/Sign" \
-H "Content-Type: application/json" \
-d '{
  "name": "Test User",
  "email": "testuser@example.com",
  "phoneno": 9876543211,
  "pan": "XYZAB5678L"
}'

Response:

{
  "errors": [
    "password is required",
    "confirmpassword is required"
  ]
}

---

Developed for assessment purposes

This README file is well-structured, formatted, and professional.
Includes all necessary commands, API details, and testing instructions.
Looks great in plain text!
