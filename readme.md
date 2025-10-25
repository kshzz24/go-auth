# Go Gin Project

A RESTful API project built with [Gin](https://github.com/gin-gonic/gin) in Go. This project demonstrates basic CRUD operations, user authentication, environment variable management, and clean project structure.

## Features

- 🔐 User authentication (signup/login)
- ✨ CRUD operations
- 🔧 Environment variable configuration with `.env`
- 🗄️ MongoDB integration
- 📁 Structured and modular project setup

## Prerequisites

- [Go](https://golang.org/dl/) >= 1.21
- [MongoDB](https://www.mongodb.com/) (running locally or in the cloud)
- [Git](https://git-scm.com/)

## Getting Started

**1. Clone the repository**

```bash
git clone https://github.com/your-username/your-repo.git
cd your-repo
```

**2. Create a `.env` file in the root directory**

```env
PORT=8000
MONGO_URI=mongodb://localhost:27017/mydatabase
JWT_SECRET=your_jwt_secret
```

**3. Install dependencies**

```bash
go mod tidy
```

**4. Run the server**

```bash
go run main.go
```

**5. Access the API**
The server will start on `http://localhost:8000`. Test endpoints using Postman, cURL, or any API testing tool.

## Project Structure

```
.
├── controllers/        # Handlers for API routes
├── models/            # MongoDB models and schemas
├── routes/            # API route definitions
├── .env               # Environment variables (not tracked by Git)
├── main.go            # Application entry point
├── go.mod             # Go module dependencies
└── README.md          # Project documentation
```

## API Endpoints

**Authentication**

- `POST /api/auth/signup` - Register a new user
- `POST /api/auth/login` - Login and receive JWT token

**Users (Protected)**

- `GET /api/users` - Get all users
- `GET /api/users/:id` - Get user by ID
- `PUT /api/users/:id` - Update user
- `DELETE /api/users/:id` - Delete user

## Environment Variables

### 🔌 `PORT`

**Description:** Server port  
**Example:** `8000`  
**Default:** `8000`

### 🗄️ `MONGO_URI`

**Description:** MongoDB connection string  
**Example:** `mongodb://localhost:27017/mydatabase`  
**Required:** Yes

### 🔐 `JWT_SECRET`

**Description:** Secret key for JWT signing  
**Example:** `your_secure_random_key`  
**Required:** Yes  
**Note:** Use a strong, random string in production

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
