# Seller Account Backend - Complete Setup Documentation

## Table of Contents
1. [Overview](#overview)
2. [Architecture](#architecture)
3. [Installation & Setup](#installation--setup)
4. [File Structure](#file-structure)
5. [Database Schema](#database-schema)
6. [API Endpoints](#api-endpoints)
7. [Code Examples](#code-examples)
8. [Testing](#testing)
9. [Deployment](#deployment)
10. [Troubleshooting](#troubleshooting)

---

## Overview

This is a production-ready backend implementation for seller profile management in an e-commerce platform built with:
- **Go** (Golang)
- **net/http** (HTTP server)
- **GORM** (ORM for PostgreSQL)
- **PostgreSQL** (Database)
- **Cloudinary** (Image uploads)
- **JWT** (Authentication)

### Features Implemented
✅ Profile management (account info, photos)
✅ Store management (logo, banner, details)
✅ Policies management (return, shipping, FAQ)
✅ Notification preferences
✅ Verification documents
✅ Security (2FA, login activity)
✅ Account management (deactivate, delete)
✅ Statistics dashboard
✅ Image uploads to Cloudinary
✅ Complete error handling

---

## Architecture

### Clean Architecture Pattern
```
┌─────────────────────────────────────────────────────┐
│                    HTTP Handler                      │
│  (Receives requests, validates, returns responses)   │
└───────────────────────┬─────────────────────────────┘
                        │
                        ▼
┌─────────────────────────────────────────────────────┐
│                   Service Layer                      │
│     (Business logic, validation, coordination)       │
└───────────────────────┬─────────────────────────────┘
                        │
                        ▼
┌─────────────────────────────────────────────────────┐
│                 Repository Layer                     │
│        (Database operations, queries, GORM)          │
└───────────────────────┬─────────────────────────────┘
                        │
                        ▼
┌─────────────────────────────────────────────────────┐
│                  PostgreSQL Database                 │
└─────────────────────────────────────────────────────┘
```

### Layer Responsibilities

**DTOs (Data Transfer Objects)**
- Define request/response structures
- Validation rules
- Type conversions

**Handlers**
- HTTP request parsing
- Call service methods
- Format responses
- Handle file uploads

**Services**
- Business logic
- Data validation
- Coordinate multiple operations
- Error handling

**Repositories**
- Database queries
- CRUD operations
- Transaction management
- Data retrieval

---

## Installation & Setup

### Prerequisites
```bash
# Go 1.21 or higher
go version

# PostgreSQL 14 or higher
psql --version

# Git
git version
```

### Step 1: Clone & Install Dependencies

```bash
# Navigate to backend directory
cd backend

# Install dependencies
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/go-playground/validator/v10
go get -u github.com/cloudinary/cloudinary-go/v2
go get -u github.com/joho/godotenv
go get -u github.com/dgrijalva/jwt-go  # For JWT
```

### Step 2: Environment Configuration

Create `.env` file:
```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password_here
DB_NAME=ecommerce_db

# Server Configuration
SERVER_PORT=8080

# JWT Configuration
JWT_SECRET=your-super-secret-key-change-this-in-production
JWT_EXPIRATION=24h

# Cloudinary Configuration
CLOUDINARY_CLOUD_NAME=your_cloud_name
CLOUDINARY_API_KEY=your_api_key
CLOUDINARY_API_SECRET=your_api_secret

# Upload Configuration
MAX_UPLOAD_SIZE=5242880  # 5MB in bytes
ALLOWED_IMAGE_TYPES=jpg,jpeg,png,webp,gif

# CORS Configuration
ALLOWED_ORIGINS=http://localhost:3000,http://localhost:5173
```

### Step 3: Database Setup

```sql
-- Create database
CREATE DATABASE ecommerce_db;

-- Connect to database
\c ecommerce_db;

-- Tables will be auto-created by GORM AutoMigrate
-- Run: go run main.go
-- GORM will create all tables based on models
```

### Step 4: Run Migrations

```bash
# Run application (GORM auto-migrates)
go run main.go

# Or create separate migration file
go run migrations/migrate.go
```

### Step 5: Start Server

```bash
# Development
go run main.go

# Production build
go build -o server main.go
./server
```

Expected output:
```
Initializing Cloudinary...
✓ Cloudinary initialized successfully
Connecting to database...
✓ Database connected successfully
Running database migrations...
✓ Migrations completed
🚀 Server starting on port 8080
📡 API available at http://localhost:8080/api
☁️  Cloudinary integration active
```

---

## File Structure

```
backend/
├── main.go                          # Application entry point
├── .env                             # Environment variables (DO NOT COMMIT)
├── .gitignore                       # Git ignore file
├── go.mod                           # Go module file
├── go.sum                           # Go dependencies
│
├── config/
│   └── cloudinary.go                # Cloudinary configuration
│
├── dto/
│   └── seller_profile_dto.go        # All DTOs for seller operations
│
├── models/
│   ├── user.go                      # User/Seller model (existing)
│   ├── reg_user.go                  # Registration user model
│   ├── role.go                      # User roles
│   ├── product.go                   # Product model
│   ├── seller_notification_preference.go  # Notification settings
│   ├── seller_login_activity.go     # Login tracking
│   ├── seller_verification_document.go    # Verification docs
│   └── seller_policy.go             # Store policies
│
├── repository/
│   └── seller_profile_repository.go # Database operations
│
├── service/
│   └── seller_profile_service.go    # Business logic
│
├── handler/
│   └── seller_profile_handler.go    # HTTP handlers
│
├── middleware/
│   └── auth_middleware.go           # JWT authentication
│
├── routes/
│   └── seller_routes.go             # Route definitions
│
├── utils/
│   ├── response.go                  # HTTP response helpers
│   ├── validation.go                # Input validation
│   ├── cloudinary_upload.go         # Image upload utilities
│   └── error_handler.go             # Error handling
│
└── migrations/
    └── migrate.go                   # Database migration script
```

---

## Database Schema

### Core Models

#### 1. User (Seller) Table
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    role_id INTEGER NOT NULL,
    user_id INTEGER UNIQUE NOT NULL,  -- References reg_users.id
    
    -- Business Info
    business_email VARCHAR(100) UNIQUE DEFAULT 'N/A',
    password VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    store_name VARCHAR(100) NOT NULL,
    store_slug VARCHAR(100) UNIQUE,
    store_desc TEXT,
    store_logo VARCHAR(500),
    store_banner VARCHAR(500),
    
    -- Business Details
    business_type VARCHAR(50) DEFAULT 'individual',
    tax_number VARCHAR(50),
    business_license VARCHAR(100),
    
    -- Financial
    total_sales DECIMAL(12,2) DEFAULT 0,
    total_earnings DECIMAL(12,2) DEFAULT 0,
    total_orders INTEGER DEFAULT 0,
    average_rating DECIMAL(3,2) DEFAULT 0,
    commission DECIMAL(5,2) DEFAULT 10,
    
    -- Address
    address TEXT DEFAULT 'N/A',
    city VARCHAR(50),
    state VARCHAR(50),
    country VARCHAR(50) DEFAULT 'Bangladesh',
    postal_code VARCHAR(20),
    
    -- Status
    is_active BOOLEAN DEFAULT TRUE,
    is_verified BOOLEAN DEFAULT FALSE,
    is_approved BOOLEAN DEFAULT FALSE,
    approved_at TIMESTAMP,
    rejected_at TIMESTAMP,
    rejection_reason TEXT,
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
```

#### 2. Seller Notification Preferences
```sql
CREATE TABLE seller_notification_preferences (
    id SERIAL PRIMARY KEY,
    seller_id INTEGER NOT NULL UNIQUE REFERENCES users(id),
    
    order_email BOOLEAN DEFAULT TRUE,
    order_sms BOOLEAN DEFAULT FALSE,
    order_push BOOLEAN DEFAULT TRUE,
    
    message_email BOOLEAN DEFAULT TRUE,
    message_sms BOOLEAN DEFAULT FALSE,
    message_push BOOLEAN DEFAULT TRUE,
    
    marketing_email BOOLEAN DEFAULT FALSE,
    marketing_sms BOOLEAN DEFAULT FALSE,
    marketing_push BOOLEAN DEFAULT FALSE,
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
```

#### 3. Seller Login Activity
```sql
CREATE TABLE seller_login_activities (
    id SERIAL PRIMARY KEY,
    seller_id INTEGER NOT NULL REFERENCES users(id),
    device VARCHAR(100),
    browser VARCHAR(50),
    os VARCHAR(50),
    location VARCHAR(100),
    ip_address VARCHAR(45),
    is_current BOOLEAN DEFAULT FALSE,
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    
    INDEX idx_seller_login (seller_id, created_at)
);
```

#### 4. Seller Verification Documents
```sql
CREATE TABLE seller_verification_documents (
    id SERIAL PRIMARY KEY,
    seller_id INTEGER NOT NULL REFERENCES users(id),
    document_type VARCHAR(50) NOT NULL,  -- identity, tax, bank
    document_url VARCHAR(500) NOT NULL,
    status VARCHAR(50) DEFAULT 'pending',  -- pending, verified, rejected
    rejected_at TIMESTAMP,
    verified_at TIMESTAMP,
    rejection_reason TEXT,
    verified_by INTEGER,  -- Admin user ID
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    
    INDEX idx_seller_doc (seller_id, document_type)
);
```

#### 5. Seller Policies
```sql
CREATE TABLE seller_policies (
    id SERIAL PRIMARY KEY,
    seller_id INTEGER NOT NULL UNIQUE REFERENCES users(id),
    return_policy TEXT,
    shipping_policy TEXT,
    faq TEXT,
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
```

### Indexes for Performance
```sql
CREATE INDEX idx_users_email ON users(business_email);
CREATE INDEX idx_users_store_slug ON users(store_slug);
CREATE INDEX idx_users_is_active ON users(is_active);
CREATE INDEX idx_seller_notifications ON seller_notification_preferences(seller_id);
CREATE INDEX idx_login_activity ON seller_login_activities(seller_id, created_at DESC);
CREATE INDEX idx_verification_docs ON seller_verification_documents(seller_id, document_type);
```

---

## API Endpoints

### Authentication Required
All endpoints require JWT Bearer token in Authorization header:
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### 1. Profile Management

#### Get Full Profile
```http
GET /api/seller/profile
```

**Response:**
```json
{
  "id": 42,
  "first_name": "John",
  "last_name": "Doe",
  "email": "john@example.com",
  "phone": "+8801712345678",
  "profile_photo": "https://res.cloudinary.com/...",
  "store_name": "John's Plant Store",
  "store_slug": "johns-plant-store-1234567890",
  "store_description": "Quality plants for your garden",
  "store_logo": "https://res.cloudinary.com/...",
  "store_banner": "https://res.cloudinary.com/...",
  "business_email": "business@johnsstore.com",
  "business_type": "nursery",
  "address": "123 Main St, Dhaka",
  "city": "Dhaka",
  "state": "Dhaka Division",
  "country": "Bangladesh",
  "postal_code": "1200",
  "total_sales": 15000.50,
  "total_earnings": 13500.45,
  "total_orders": 120,
  "average_rating": 4.8,
  "is_active": true,
  "is_verified": true,
  "is_approved": true,
  "two_factor_enabled": false,
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-11-02T14:25:00Z"
}
```

#### Get Statistics
```http
GET /api/seller/statistics
```

**Response:**
```json
{
  "total_sales": 15000.50,
  "total_orders": 120,
  "active_products": 45,
  "rating": 4.8,
  "total_reviews": 89,
  "pending_orders": 5,
  "completed_orders": 115,
  "total_earnings": 13500.45
}
```

### 2. Account Management

#### Update Account
```http
PUT /api/seller/account
Content-Type: application/json

{
  "first_name": "John",
  "last_name": "Doe",
  "phone": "+8801712345678"
}
```

#### Upload Profile Photo
```http
POST /api/seller/account/photo
Content-Type: multipart/form-data

photo: (binary file)
```

**Response:**
```json
{
  "photo_url": "https://res.cloudinary.com/demo/image/upload/...",
  "message": "Profile photo updated successfully"
}
```

#### Change Password
```http
POST /api/seller/password/change
Content-Type: application/json

{
  "current_password": "oldpassword123",
  "new_password": "newpassword456",
  "confirm_password": "newpassword456"
}
```

#### Deactivate Account
```http
POST /api/seller/account/deactivate
```

#### Delete Account
```http
DELETE /api/seller/account/delete
Content-Type: application/json

{
  "password": "currentpassword123"
}
```

### 3. Store Management

#### Update Store
```http
PUT /api/seller/store
Content-Type: application/json

{
  "store_name": "John's Plant Paradise",
  "store_description": "Premium quality plants and gardening supplies",
  "website": "https://johnsplants.com",
  "phone": "+8801712345678",
  "business_email": "business@johnsplants.com",
  "address": "123 Green Street",
  "city": "Dhaka",
  "state": "Dhaka Division",
  "country": "Bangladesh",
  "postal_code": "1200"
}
```

#### Update Branding
```http
PUT /api/seller/store/branding
Content-Type: multipart/form-data

logo: (binary file)
banner: (binary file)
```

#### Update Policies
```http
PUT /api/seller/store/policies
Content-Type: application/json

{
  "return_policy": "Returns accepted within 30 days...",
  "shipping_policy": "Standard shipping: 5-7 business days...",
  "faq": "Q: How do I care for my plants?\nA: Water regularly..."
}
```

### 4. Notifications

#### Get Preferences
```http
GET /api/seller/notifications/preferences
```

**Response:**
```json
{
  "order_email": true,
  "order_sms": false,
  "order_push": true,
  "message_email": true,
  "message_sms": false,
  "message_push": true,
  "marketing_email": false,
  "marketing_sms": false,
  "marketing_push": false
}
```

#### Update Preferences
```http
PUT /api/seller/notifications/preferences
Content-Type: application/json

{
  "order_email": true,
  "order_sms": true,
  "order_push": true,
  "message_email": true,
  "message_sms": false,
  "message_push": true,
  "marketing_email": false,
  "marketing_sms": false,
  "marketing_push": false
}
```

### 5. Verification

#### Get Status
```http
GET /api/seller/verification/status
```

**Response:**
```json
{
  "identity": "verified",
  "tax": "pending",
  "bank": "not-submitted"
}
```

#### Upload Document
```http
POST /api/seller/verification/upload
Content-Type: multipart/form-data

document_type: identity
document: (binary file)
```

### 6. Security

#### Toggle 2FA
```http
PUT /api/seller/security/2fa
Content-Type: application/json

{
  "enabled": true
}
```

#### Get Login Activity
```http
GET /api/seller/security/activity
```

**Response:**
```json
{
  "activities": [
    {
      "id": 1,
      "device": "Chrome on MacOS",
      "browser": "Chrome 120.0",
      "location": "Dhaka, Bangladesh",
      "ip_address": "192.168.1.1",
      "time_ago": "2 hours ago",
      "is_current": true,
      "created_at": "2024-11-02T12:30:00Z"
    }
  ]
}
```

---

## Code Examples

### 1. Creating a New Service Method

```go
// In service/seller_profile_service.go

func (s *SellerProfileService) UpdateBusinessInfo(sellerID uint, req *dto.UpdateBusinessInfoRequest) error {
    // Validate request
    if err := utils.ValidateStruct(req); err != nil {
        return err
    }
    
    // Prepare updates
    updates := map[string]interface{}{
        "business_type": req.BusinessType,
        "tax_number": req.TaxNumber,
        "business_license": req.BusinessLicense,
    }
    
    // Update in database
    if err := s.repo.UpdateSeller(sellerID, updates); err != nil {
        return errors.New("failed to update business information")
    }
    
    return nil
}
```

### 2. Adding a New Endpoint

```go
// In handler/seller_profile_handler.go

func (h *SellerProfileHandler) UpdateBusinessInfo(w http.ResponseWriter, r *http.Request) {
    sellerID := r.Context().Value("seller_id").(uint)
    
    var req dto.UpdateBusinessInfoRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        utils.RespondError(w, http.StatusBadRequest, "Invalid request body")
        return
    }
    
    if err := h.service.UpdateBusinessInfo(sellerID, &req); err != nil {
        utils.RespondError(w, http.StatusBadRequest, err.Error())
        return
    }
    
    utils.RespondJSON(w, http.StatusOK, map[string]string{
        "message": "Business information updated successfully",
    })
}

// In routes/seller_routes.go

mux.Handle("PUT /api/seller/business-info", 
    authMiddleware.RequireSeller(http.HandlerFunc(handler.UpdateBusinessInfo)))
```

### 3. Database Query Example

```go
// In repository/seller_profile_repository.go

func (r *SellerProfileRepository) GetTopSellers(limit int) ([]models.User, error) {
    var sellers []models.User
    err := r.db.
        Where("is_active = ? AND is_approved = ?", true, true).
        Order("average_rating DESC, total_orders DESC").
        Limit(limit).
        Preload("RegUser").
        Find(&sellers).Error
    
    return sellers, err
}
```

---

## Testing

### Unit Tests

```go
// service/seller_profile_service_test.go

package service

import (
    "testing"
    "your-project/dto"
)

func TestGetStatistics(t *testing.T) {
    // Mock repository
    mockRepo := &MockSellerProfileRepository{
        GetSellerStatisticsFunc: func(sellerID uint) (*dto.SellerStatisticsResponse, error) {
            return &dto.SellerStatisticsResponse{
                TotalSales: 15000.50,
                TotalOrders: 120,
            }, nil
        },
    }
    
    service := NewSellerProfileService(mockRepo)
    
    // Test
    stats, err := service.GetStatistics(42)
    
    // Assertions
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    
    if stats.TotalSales != 15000.50 {
        t.Errorf("Expected TotalSales 15000.50, got %f", stats.TotalSales)
    }
}
```

### Integration Tests

```bash
# Run all tests
go test ./...

# Run specific package
go test ./service/...

# With coverage
go test -cover ./...

# Verbose output
go test -v ./...
```

### API Testing with cURL

```bash
# Test profile endpoint
curl -X GET http://localhost:8080/api/seller/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"

# Test update account
curl -X PUT http://localhost:8080/api/seller/account \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "phone": "+8801712345678"
  }'

# Test file upload
curl -X POST http://localhost:8080/api/seller/account/photo \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -F "photo=@/path/to/photo.jpg"
```

---

## Deployment

### Production Build

```bash
# Build binary
go build -o seller-backend main.go

# Run binary
./seller-backend
```

### Docker Deployment

```dockerfile
# Dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/server .
COPY --from=builder /app/.env .

EXPOSE 8080
CMD ["./server"]
```

```bash
# Build image
docker build -t seller-backend .

# Run container
docker run -p 8080:8080 --env-file .env seller-backend
```

### Environment Variables for Production

```env
# Use production database
DB_HOST=production-db.example.com
DB_PORT=5432
DB_USER=prod_user
DB_PASSWORD=strong_production_password
DB_NAME=ecommerce_production

# Strong JWT secret
JWT_SECRET=your-very-strong-random-secret-key-here

# Production Cloudinary
CLOUDINARY_CLOUD_NAME=production_cloud
CLOUDINARY_API_KEY=production_api_key
CLOUDINARY_API_SECRET=production_api_secret

# CORS
ALLOWED_ORIGINS=https://yourdomain.com,https://www.yourdomain.com
```

---

## Troubleshooting

### Common Issues

#### 1. Database Connection Failed
```
Error: failed to connect database: dial tcp: connect: connection refused
```

**Solution:**
- Check PostgreSQL is running: `sudo systemctl status postgresql`
- Verify credentials in `.env`
- Check firewall settings

#### 2. Cloudinary Upload Failed
```
Error: failed to upload file to cloudinary: invalid credentials
```

**Solution:**
- Verify Cloudinary credentials in `.env`
- Check internet connection
- Ensure Cloudinary account is active

#### 3. JWT Token Invalid
```
Error: invalid or expired token
```

**Solution:**
- Check JWT_SECRET matches between token generation and verification
- Verify token hasn't expired
- Ensure Authorization header format: `Bearer <token>`

#### 4. CORS Error
```
Error: CORS policy: No 'Access-Control-Allow-Origin' header
```

**Solution:**
- Add frontend URL to ALLOWED_ORIGINS in `.env`
- Verify CORS middleware is enabled in main.go

### Debug Mode

Enable detailed logging:
```go
// In main.go
import "log"

// Add before server start
log.SetFlags(log.LstdFlags | log.Lshortfile)

// Use throughout code
log.Printf("Debug: sellerID=%d, operation=%s", sellerID, "update_profile")
```

---

## Performance Optimization

### Database Optimization
```go
// Use indexes
r.db.Model(&models.Product{}).
    Where("seller_id = ?", sellerID).
    Order("created_at DESC").
    Limit(10).
    Find(&products)

// Batch operations
r.db.Create(&products)  // Creates all at once

// Use Select to limit fields
r.db.Select("id", "store_name", "total_sales").
    Find(&sellers)
```

### Caching Strategy
```go
// Use Redis for caching statistics
func (s *SellerProfileService) GetStatisticsCached(sellerID uint) (*dto.SellerStatisticsResponse, error) {
    // Try cache first
    if cached := redis.Get(fmt.Sprintf("stats:%d", sellerID)); cached != nil {
        return cached, nil
    }
    
    // Get from database
    stats, err := s.GetStatistics(sellerID)
    if err != nil {
        return nil, err
    }
    
    // Cache for 5 minutes
    redis.Set(fmt.Sprintf("stats:%d", sellerID), stats, 5*time.Minute)
    
    return stats, nil
}
```

---

## Security Best Practices

1. **Environment Variables**: Never commit `.env` to git
2. **Password Hashing**: Always use bcrypt with cost 10+
3. **JWT Secrets**: Use strong, random secrets (32+ characters)
4. **Input Validation**: Validate all user inputs
5. **SQL Injection**: Use GORM parameterized queries (it does automatically)
6. **File Upload**: Validate size and type before uploading
7. **Rate Limiting**: Implement rate limiting for API endpoints
8. **HTTPS**: Use HTTPS in production
9. **CORS**: Restrict origins to your frontend domains only
10. **Error Messages**: Don't expose sensitive information in errors

---

## Maintenance

### Regular Tasks
- Monitor database size and clean old data
- Review and rotate JWT secrets periodically
- Check Cloudinary usage and storage
- Update Go dependencies: `go get -u ./...`
- Backup database regularly
- Monitor server logs for errors
- Check API response times

### Monitoring
```go
// Add request logging middleware
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}
```

---

## Support & Resources

### Documentation
- GORM: https://gorm.io/docs/
- Cloudinary Go SDK: https://github.com/cloudinary/cloudinary-go
- Go Validator: https://github.com/go-playground/validator
- JWT Go: https://github.com/dgrijalva/jwt-go

### Getting Help
- Check this documentation first
- Review error logs
- Search GitHub issues
- Ask in Go community forums

---

## Change Log

### Version 1.0.0 (Current)
- Initial release
- Complete seller profile management
- Cloudinary integration
- Security features (2FA, login tracking)
- Notification preferences
- Verification system
- Statistics dashboard

---

## License

[Your License Here]

---

## Contributors

[Your Name/Team]

---

**Last Updated**: November 2, 2024
**Version**: 1.0.0
**Status**: Production Ready ✅