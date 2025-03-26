# 📈 Stock Broker Application  

A backend service for user authentication and management in a stock broker platform.

---

## 🚀 Features

- **User Signup & Sign-In** with validations and hashed passwords
- **Data Validation** for email, phone number, PAN, and password
- **Database Integration** using PostgreSQL and GORM
- **Automated Migrations** handled in `sqlSetup.go`
- **Error Handling** with structured responses

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
  dbname: "postgres"
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

#### 📌 POST `/signup`
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

---

### 🔹 User Sign-In

#### 📌 POST `/signin`
Authenticates an existing user.

#### 🔹 Request Body

```json
{
  "email": "johndoe@example.com",
  "password": "Strong@123"
}
```

#### 🔹 Response

```json
{
  "message": "User has been logged in",
}
```

---

## ❌ Error Handling

All errors now follow a **structured response format** with field-specific errors.

### 🔹 Example Validation Error

```json
{
  "errors": {
    "email": "email is required",
    "password": "password is required"
  }
}
```

### 🔹 Example Authentication Error

```json
{
  "error": "Invalid credentials"
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

### 🔹 Test Signup with Postman

1. Open Postman.
2. Select a **POST** request.
3. Enter the URL: `http://localhost:8080/signup`
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

### 🔹 Test Sign-In with cURL

✅ **Successful Login**

```sh
curl -X POST "http://localhost:8080/signin" -H "Content-Type: application/json" -d '{
  "email": "testuser@example.com",
  "password": "Test@1234"
}'
```

❌ **Error Case (Wrong Password)**

```sh
curl -X POST "http://localhost:8080/signin" -H "Content-Type: application/json" -d '{
  "email": "testuser@example.com",
  "password": "WrongPass123"
}'
```

📌 **Response:**

```json
{
  "error": "Invalid credentials"
}
```


![image](https://github.com/user-attachments/assets/0a8779c2-ef42-4592-95d2-b8dab1af544b)
![image](https://github.com/user-attachments/assets/28b50f7f-fb8f-4a50-9f52-af978fe081a9)
![image](https://github.com/user-attachments/assets/9e852d69-8053-4f0c-bb5f-9557f7761496)
![image](https://github.com/user-attachments/assets/d25b8114-71de-4a47-82c2-ee96beb13342)
![image](https://github.com/user-attachments/assets/c7721ff0-c421-46f1-9700-dbae65417b91)
![image](https://github.com/user-attachments/assets/9d2b987f-f060-485b-a608-ba784c8512e3)
![image](https://github.com/user-attachments/assets/4a5d3ec3-e344-41fc-8407-6836515a732c)
![image](https://github.com/user-attachments/assets/4de11705-994e-44e3-88b5-a655a1a8f71a)
![image](https://github.com/user-attachments/assets/a25c8278-0334-4567-83e3-59bcefa6cd19)
![image](https://github.com/user-attachments/assets/004eda99-a181-4f9a-85bb-90b703d0b51a)
![image](https://github.com/user-attachments/assets/85aa5235-020b-4333-8d79-96284eabcc38)
![image](https://github.com/user-attachments/assets/d3c8ca8c-bf29-452e-89f5-e39be924ac75)













