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
```

### 🔹 Configure Environment Variables
Modify `config.yaml`  to set up database credentials:

```yaml
database:
  host: "localhost"
  port: 5432
  user: "your_user"
  password: "your_password"
  dbname: "stock_broker_db"
```

### 🔹 Run Migrations
Migrations are automatically handled in `sqlSetup.go`, so just run the main file.

### 🔹 Start the Server

```sh
go run main.go
```

✅ The server will start on `localhost:8080`

---

## 📡 API Endpoints

### 🔹 User Signup

#### 📌 POST `/Signup`
Registers a new user with validation checks.

#### 🔹 Request Body

```json
{
  "name": "John Doe",
  "password": "Strong@123",
  "confirmpassword": "Strong@123",
  "email": "johndoe@example.com",
  "phoneno": 9876543210,
  "pan": "ABCDE1234F"
}
```

#### 🔹 Response

```json
{
  "message": "User created successfully"
}
```

### ❌ Error Handling
The application provides detailed validation errors. Example:

```json
{
  "errors": [
    "email is required",
    "password must be at least 8 characters"
  ]
}
```

---

## ⚠ Additional Notes

- The application uses hashed passwords for security.
- Duplicate username/email results in a conflict error.
- Phone numbers are stored as `BIGINT` to avoid string conversion issues.

---

## 🧪 Running Tests

You can test the API using **Postman** or **cURL**.

### 🔹 Test with Postman

1. Open Postman.
2. Select a **POST** request.
3. Enter the URL: `http://localhost:8080/Signup`
4. Go to **Body → raw → JSON**.
5. Paste the following:

```json
{
  "name": "Test User",
  "password": "Test@1234",
  "confirmpassword": "Test@1234",
  "email": "testuser@example.com",
  "phoneno": 9876543211,
  "pan": "XYZAB5678L"
}
```

6. Click **Send**.
7. You should receive:

```json
{
  "message": "User created successfully"
}
```

### 🔹 Test with cURL

✅ **Successful Request**

```sh
curl -X POST "http://localhost:8080/Signup" \
-H "Content-Type: application/json" \
-d '{
  "name": "Test User",
  "password": "Test@1234",
  "confirmpassword": "Test@1234",
  "email": "testuser@example.com",
  "phoneno": 9876543211,
  "pan": "XYZAB5678L"
}'
```

❌ **Error Case (Missing Password)**

```sh
curl -X POST "http://localhost:8080/Signup" \
-H "Content-Type: application/json" \
-d '{
  "name": "Test User",
  "email": "testuser@example.com",
  "phoneno": 9876543211,
  "pan": "XYZAB5678L"
}'
```

📌 **Response:**

```json
{
  "errors": [
    "password is required",
    "confirmpassword is required"
  ]
}
```

![image](https://github.com/user-attachments/assets/0a8779c2-ef42-4592-95d2-b8dab1af544b)


