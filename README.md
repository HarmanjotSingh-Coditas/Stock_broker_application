# 📈 Stock Broker Application  

A backend service for user authentication and management in a stock broker platform.

---

## 🚀 Features

- **User Signup** with validations and hashed passwords
- **Data Validation** for email, phone number, PAN, and password
- **Database Integration** using PostgreSQL and GORM
- **Automated Migrations** handled in `sqlSetup.go`
- **Error Handling** with detailed responses

---

## 🛠 Technologies Used

- **Golang** (Gin, GORM)
- **PostgreSQL**
- **Validator for Input Validation**

---

## 📌 Setup Instructions

### 🔹 Prerequisites

- Install **Go** (1.20+)
- Install **PostgreSQL**

### 🔹 Clone the Repository

```sh
git clone https://github.com/HarmanjotSingh-Coditas/stock_broker_application.git
cd stock_broker_application
🔹 Configure Environment Variables
Modify config.yaml (or .env if used) to set up database credentials:

yaml
Copy
Edit
database:
  host: "localhost"
  port: 5432
  user: "your_user"
  password: "your_password"
  dbname: "stock_broker_db"
🔹 Run Migrations
Migrations are automatically handled in sqlSetup.go, so just run the main file.

🔹 Start the Server
sh
Copy
Edit
go run main.go
✅ The server will start on localhost:8080

📡 API Endpoints
🔹 User Signup
📌 POST /Sign
Registers a new user with validation checks.

🔹 Request Body
json
Copy
Edit
{
  "name": "John Doe",
  "password": "Strong@123",
  "confirmpassword": "Strong@123",
  "email": "johndoe@example.com",
  "phoneno": 9876543210,
  "pan": "ABCDE1234F"
}
🔹 Response
json
Copy
Edit
{
  "message": "User created successfully"
}
❌ Error Handling
The application provides detailed validation errors. Example:

json
Copy
Edit
{
  "errors": [
    "email is required",
    "password must be at least 8 characters"
  ]
}
⚠ Additional Notes
The application uses hashed passwords for security.

Duplicate username/email results in a conflict error.

Phone numbers are stored as BIGINT to avoid string conversion issues.

🧪 Running Tests
You can test the API using Postman or cURL.

🔹 Test with Postman
Open Postman.

Select a POST request.

Enter the URL:

arduino
Copy
Edit
http://localhost:8080/Sign
Go to Body → raw → JSON.

Paste the following:

json
Copy
Edit
{
  "name": "Test User",
  "password": "Test@1234",
  "confirmpassword": "Test@1234",
  "email": "testuser@example.com",
  "phoneno": 9876543211,
  "pan": "XYZAB5678L"
}
Click Send.

You should receive:

json
Copy
Edit
{
  "message": "User created successfully"
}
🔹 Test with cURL
✅ Successful Request
sh
Copy
Edit
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
❌ Error Case (Missing Password)
sh
Copy
Edit
curl -X POST "http://localhost:8080/Sign" \
-H "Content-Type: application/json" \
-d '{
  "name": "Test User",
  "email": "testuser@example.com",
  "phoneno": 9876543211,
  "pan": "XYZAB5678L"
}'
📌 Response:

json
Copy
Edit
{
  "errors": [
    "password is required",
    "confirmpassword is required"
  ]
}
