# 🌳 Ecoleafo - Multi-Vendor E-Commerce Platform

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-14+-316192?style=flat&logo=postgresql)](https://www.postgresql.org)
[![License](https://img.shields.io/badge/License-Proprietary-red.svg)](LICENSE)

> A comprehensive multi-vendor marketplace platform built with Go, designed specifically for plant nurseries and garden stores.

---

## 📋 Table of Contents

- [Overview](#-overview)
- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Getting Started](#-getting-started)
- [API Documentation](#-api-documentation)
- [Project Structure](#-project-structure)
- [Database Schema](#-database-schema)
- [Contributing](#-contributing)
- [License](#-license)

---

## 🎯 Overview

**Ecoleafo** is a production-ready, scalable e-commerce platform that enables multiple sellers to operate their online stores under one unified marketplace. Built with modern Go practices and clean architecture principles, it provides a robust foundation for plant and garden product sales.

### Why Ecoleafo?

- 🚀 **High Performance**: Built with Go for exceptional speed and concurrency
- 🏗️ **Clean Architecture**: Repository pattern, service layer, and dependency injection
- 🔒 **Enterprise Security**: JWT authentication, RBAC, and comprehensive audit logging
- 💰 **Financial Management**: Built-in commission system and earnings tracking
- 📊 **Advanced Analytics**: Real-time dashboards and detailed reports
- 🔌 **RESTful API**: Well-documented, consistent API design

---

## ✨ Features

### For Sellers
- ✅ Easy store setup with 4-step onboarding
- ✅ Product management with multiple images
- ✅ Real-time inventory tracking
- ✅ Order management and fulfillment
- ✅ Earnings dashboard with analytics
- ✅ Payout request system
- ✅ Customer messaging

### For Buyers
- ✅ Browse products by category
- ✅ Advanced search and filtering
- ✅ Shopping cart (guest & authenticated)
- ✅ Secure checkout process
- ✅ Order tracking
- ✅ Product reviews and ratings
- ✅ Wishlist management

### For Administrators
- ✅ User approval workflow
- ✅ Product moderation
- ✅ Commission management
- ✅ Financial analytics
- ✅ Platform earnings tracking
- ✅ Seller performance monitoring
- ✅ Comprehensive audit logging

---

## 🛠️ Tech Stack

### Backend
- **Language**: Go 1.21+
- **Database**: PostgreSQL 14+
- **ORM**: GORM
- **Authentication**: JWT with refresh tokens
- **File Storage**: Cloudinary
- **Architecture**: Clean Architecture with Repository Pattern

### Key Libraries
```go
github.com/golang-jwt/jwt/v5      // JWT authentication
gorm.io/gorm                       // ORM
gorm.io/driver/postgres            // PostgreSQL driver
github.com/cloudinary/cloudinary-go // Image storage
golang.org/x/crypto/bcrypt         // Password hashing
```

---

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher
- PostgreSQL 14 or higher
- Cloudinary account (for image uploads)

### Installation

1. **Clone the repository**
```bash
git clone https://github.com/saadahmedbd/Ecoleafo.git
cd Ecoleafo/Ecoleafo
```

2. **Install dependencies**
```bash
go mod download
```

3. **Configure environment variables**
```bash
cp .env.example .env
```

Edit `.env` with your configuration:
```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=Ecoleafo

# JWT Configuration
JWT_SECRET=your_super_secret_jwt_key
JWT_EXPIRY=24h
REFRESH_TOKEN_EXPIRY=168h

# Cloudinary Configuration
CLOUDINARY_CLOUD_NAME=your_cloud_name
CLOUDINARY_API_KEY=your_api_key
CLOUDINARY_API_SECRET=your_api_secret

# Server Configuration
PORT=3000
```

4. **Run database migrations**
```bash
go run main.go migrate
```

5. **Seed initial data (optional)**
```bash
go run main.go seed
```

6. **Start the server**
```bash
go run main.go
```

The API will be available at `http://localhost:3000`

---

## 📚 API Documentation

### Base URL
```
http://localhost:3000/api
```

### Authentication

All protected endpoints require a JWT token in the Authorization header:
```
Authorization: Bearer <your_jwt_token>
```

### Quick Start Examples

#### Register a Seller
```bash
curl -X POST http://localhost:3000/api/seller/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "seller@example.com",
    "password": "SecurePass123",
    "confirm_password": "SecurePass123",
    "first_name": "John",
    "last_name": "Doe",
    "store_name": "Green Garden Store",
    "phone": "01234567890",
    "agree_to_terms": true
  }'
```

#### Login
```bash
curl -X POST http://localhost:3000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "seller@example.com",
    "password": "SecurePass123"
  }'
```

#### Get Products
```bash
curl -X GET http://localhost:3000/api/products?page=1&limit=20
```

### API Endpoints Overview

| Category | Endpoint | Method | Description |
|----------|----------|--------|-------------|
| **Auth** | `/api/auth/register` | POST | User registration |
| | `/api/auth/login` | POST | User login |
| | `/api/auth/refresh` | POST | Refresh token |
| **Products** | `/api/products` | GET | List products |
| | `/api/products` | POST | Create product |
| | `/api/products/{id}` | GET | Get product |
| | `/api/products/{id}` | PUT | Update product |
| **Orders** | `/api/orders` | POST | Create order |
| | `/api/orders` | GET | List orders |
| | `/api/orders/{id}` | GET | Get order details |
| **Cart** | `/api/cart` | GET | Get cart |
| | `/api/cart` | POST | Add to cart |
| **Admin** | `/api/admin/sellers/earnings` | GET | All sellers earnings |
| | `/api/admin/platform/earnings` | GET | Platform analytics |

For complete API documentation, see [API_DOCUMENTATION.md](docs/API_DOCUMENTATION.md)

---

## 📁 Project Structure

```
Ecoleafo/
├── CMD/                          # Application entry points
│   └── server.go                 # Server initialization
├── Config/                       # Configuration management
│   ├── config.go                 # App configuration
│   ├── db.go                     # Database connection
│   └── cloudinary.go             # Cloudinary setup
├── Database/                     # Database operations
│   ├── MigrateRegUser.go         # Migrations
│   └── SeedSuperAdmin.go         # Seed data
├── Models/                       # Database models (40+ models)
│   ├── Users.go                  # User/Seller model
│   ├── Product.go                # Product model
│   ├── Order.go                  # Order model
│   └── ...                       # Other models
├── Rest/
│   ├── DTO/                      # Data Transfer Objects
│   │   ├── sellerAccount/        # Seller DTOs
│   │   ├── Order/                # Order DTOs
│   │   └── ...                   # Other DTOs
│   ├── Handler/                  # HTTP handlers
│   │   ├── ProductHandler/       # Product handlers
│   │   ├── OrderHandler/         # Order handlers
│   │   └── ...                   # Other handlers
│   ├── Middleware/               # HTTP middleware
│   │   ├── AuthenticateJwt.go    # JWT authentication
│   │   ├── Cors.go               # CORS handling
│   │   └── Logger.go             # Request logging
│   ├── Repository/               # Data access layer
│   │   ├── ProductRepo/          # Product repository
│   │   ├── OrderRepo/            # Order repository
│   │   └── ...                   # Other repositories
│   ├── Service/                  # Business logic layer
│   │   ├── ProductService/       # Product service
│   │   ├── OrderService/         # Order service
│   │   └── ...                   # Other services
│   └── Routes/                   # Route definitions
├── Util/                         # Utility functions
│   ├── CreateJwt.go              # JWT utilities
│   ├── HashPassword.go           # Password hashing
│   └── ...                       # Other utilities
├── .env                          # Environment variables
├── go.mod                        # Go module definition
├── go.sum                        # Go module checksums
├── main.go                       # Application entry point
└── README.md                     # This file
```

---

## 🗄️ Database Schema

### Key Tables

#### Users & Authentication
- `reg_users` - User authentication and basic information
- `users` - Seller profiles and store details
- `buyers` - Buyer profiles
- `admins` - Administrator accounts
- `roles` - User roles and permissions
- `refresh_tokens` - JWT refresh tokens

#### E-Commerce
- `products` - Product catalog
- `product_images` - Product images
- `categories` - Product categories
- `orders` - Order records
- `order_items` - Order line items
- `cart_items` - Shopping cart
- `reviews` - Product reviews

#### Financial
- `seller_earnings_summaries` - Seller earnings cache
- `order_commissions` - Commission per order
- `seller_payouts` - Payout requests
- `commission_settings` - Platform commission configuration

#### Communication
- `messages` - User messaging
- `conversations` - Message threads
- `notifications` - System notifications

### Entity Relationships

```
reg_users (1) ──→ (1) users (sellers)
reg_users (1) ──→ (1) buyers
users (1) ──→ (N) products
users (1) ──→ (N) orders
products (1) ──→ (N) order_items
orders (1) ──→ (N) order_items
users (1) ──→ (1) seller_earnings_summaries
```

---

## 🏗️ Architecture

### Design Patterns

#### Repository Pattern
Separates data access logic from business logic:
```go
type ProductRepository interface {
    Create(product *models.Product) error
    GetByID(id uint) (*models.Product, error)
    Update(product *models.Product) error
    Delete(id uint) error
}
```

#### Service Layer Pattern
Encapsulates business logic:
```go
type ProductService interface {
    CreateProduct(req *dto.CreateProductRequest) error
    GetProduct(id uint) (*dto.ProductResponse, error)
    UpdateProduct(id uint, req *dto.UpdateProductRequest) error
}
```

#### Dependency Injection
Promotes loose coupling:
```go
func NewProductService(
    productRepo repository.ProductRepository,
    imageService service.ImageService,
) ProductService {
    return &productService{
        productRepo: productRepo,
        imageService: imageService,
    }
}
```

---

## 🔒 Security

### Authentication & Authorization
- **JWT Tokens**: Secure, stateless authentication
- **Refresh Tokens**: Long-lived tokens for seamless re-authentication
- **Password Hashing**: Bcrypt with salt
- **Role-Based Access Control**: Fine-grained permissions

### Security Best Practices
- ✅ SQL injection prevention (GORM parameterized queries)
- ✅ CORS configuration
- ✅ Input validation on all endpoints
- ✅ Secure password requirements
- ✅ Audit logging for admin actions
- ✅ Rate limiting (recommended for production)

---

## 📊 Performance

### Optimizations
- Database indexing on frequently queried fields
- Pagination for large datasets
- Eager loading to prevent N+1 queries
- Connection pooling
- Cached earnings summaries

### Scalability
- Stateless API design
- Horizontal scaling ready
- Database read replicas support
- CDN for static assets (Cloudinary)

---

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./Rest/Service/ProductService/...
```

---

## 🚢 Deployment

### Docker (Recommended)

```dockerfile
# Dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o Ecoleafo main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/Ecoleafo .
COPY .env .
EXPOSE 3000
CMD ["./Ecoleafo"]
```

Build and run:
```bash
docker build -t Ecoleafo .
docker run -p 3000:3000 Ecoleafo
```

### Production Checklist
- [ ] Set strong JWT secret
- [ ] Configure production database
- [ ] Enable HTTPS
- [ ] Set up monitoring and logging
- [ ] Configure backup strategy
- [ ] Set up CI/CD pipeline
- [ ] Enable rate limiting
- [ ] Configure CDN

---

## 📈 Monitoring & Logging

### Logging
All requests are logged with:
- Request method and path
- Response status code
- Response time
- User information (if authenticated)

### Metrics (Recommended)
- API response times
- Database query performance
- Error rates
- Active users
- Order conversion rates

---

## 🤝 Contributing

We welcome contributions! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### Code Style
- Follow Go best practices and idioms
- Use `gofmt` for code formatting
- Write meaningful commit messages
- Add tests for new features
- Update documentation

---

## 📝 License

This project is proprietary software. All rights reserved.

---

## 👨‍💻 Author

**Saad Ahmed**
- GitHub: [@saadahmedbd](https://github.com/saadahmedbd)
- Email: saadahmedbd0@gmail.com

---

## 🙏 Acknowledgments

- [Go](https://golang.org/) - The Go Programming Language
- [GORM](https://gorm.io/) - The fantastic ORM library
- [PostgreSQL](https://www.postgresql.org/) - The world's most advanced open source database
- [Cloudinary](https://cloudinary.com/) - Image and video management

---

## 📞 Support

For support, email saadahmedbd0@gmail.com or open an issue in the repository.

---

<div align="center">

**Ecoleafo** - Empowering plant nurseries with modern e-commerce technology 🌳

Made with ❤️ using Go

</div>
