
## 📦 Go CRUD API

A basic RESTful API built with Go and the Gin framework, featuring JWT-based authentication and CRUD operations for products.

---

### 🚀 Features

* 🔐 JWT-based login system
* 🛡️ Authentication middleware
* 📦 Product CRUD (Create, Read, Update, Delete)


---

### 📁 Project Structure

```
go-crud-api/
├── controllers/
│   ├── auth_controller.go
│   └── product_controller.go
├── middleware/
│   └── auth.go
├── main.go
```

---

### 🔧 Setup Instructions

1. **Clone the project**

   ```bash
   git clone https://github.com/yourusername/go-crud-api.git
   cd go-crud-api
   ```

2. **Install Go dependencies**

   ```bash
   go mod tidy
   ```

3. **Run the server**

   ```bash
   go run main.go
   ```

4. **Server starts at**

   ```
   http://localhost:8080
   ```

---

### 👤 Sample Users

| Username | Password   | Role  |
| -------- | ---------- | ----- |
| `admin`  | `admin123` | admin |
| `user`   | `user123`  | user  |

---

### 🔐 Authentication

#### Login Endpoint

```http
POST /login
```

#### Request Body (JSON)

```json
{
  "username": "admin",
  "password": "admin123"
}
```

#### Response

```json
{
  "token": "<JWT_TOKEN>"
}
```

Use the token in `Authorization` header for protected routes:

```
Authorization: Bearer <JWT_TOKEN>
```

---

### 📦 Product API Endpoints

> All product endpoints require JWT token in the header.

| Method | Endpoint        | Description          |
| ------ | --------------- | -------------------- |
| POST   | `/products`     | Create a new product |
| GET    | `/products`     | List all products    |
| GET    | `/products/:id` | Get product by ID    |
| PUT    | `/products/:id` | Update product by ID |
| DELETE | `/products/:id` | Delete product by ID |

#### Example: Create Product

```http
POST /products
Authorization: Bearer <JWT_TOKEN>
Content-Type: application/json

{
  "name": "Keyboard",
  "price": 45.99
}
```

---
