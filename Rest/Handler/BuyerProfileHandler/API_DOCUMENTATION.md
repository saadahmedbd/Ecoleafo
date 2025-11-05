# Buyer Profile API Documentation

## Overview
Complete API documentation for buyer profile management including profile operations, address management, and order history.

---

## Base URL
```
http://localhost:3000/api/buyer
```

---

## Authentication
All endpoints require JWT authentication via Bearer token in the Authorization header.

```http
Authorization: Bearer <your_jwt_token>
```

---

## Endpoints

### 1. Get Buyer Profile
Retrieve the authenticated buyer's profile information.

**Endpoint:** `GET /api/buyer/profile`

**Headers:**
```http
Authorization: Bearer <token>
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Profile retrieved successfully",
  "data": {
    "id": 1,
    "user_id": 5,
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@example.com",
    "phone": "+1234567890",
    "profile_photo": "https://cloudinary.com/...",
    "date_of_birth": "1990-01-15",
    "gender": "male",
    "status": "active",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-15T00:00:00Z"
  }
}
```

**cURL Example:**
```bash
curl -X GET "http://localhost:3000/api/buyer/profile" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

### 2. Update Buyer Profile
Update buyer profile information.

**Endpoint:** `PUT /api/buyer/profile`

**Headers:**
```http
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "first_name": "John",
  "last_name": "Doe",
  "phone": "+1234567890",
  "date_of_birth": "1990-01-15",
  "gender": "male"
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Profile updated successfully",
  "data": {
    "id": 1,
    "first_name": "John",
    "last_name": "Doe",
    "phone": "+1234567890",
    "date_of_birth": "1990-01-15",
    "gender": "male"
  }
}
```

**cURL Example:**
```bash
curl -X PUT "http://localhost:3000/api/buyer/profile" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "phone": "+1234567890",
    "date_of_birth": "1990-01-15",
    "gender": "male"
  }'
```

---

### 3. Upload Profile Picture
Upload or update buyer profile picture.

**Endpoint:** `POST /api/buyer/profile/upload-picture`

**Headers:**
```http
Authorization: Bearer <token>
Content-Type: multipart/form-data
```

**Request Body (Form Data):**
- `photo` (file): Image file (max 2MB, formats: jpg, jpeg, png)

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Profile photo uploaded successfully",
  "data": {
    "photo_url": "https://res.cloudinary.com/your-cloud/image/upload/v123456/buyer-profiles/photo.jpg",
    "message": "Profile photo updated successfully"
  }
}
```

**cURL Example:**
```bash
curl -X POST "http://localhost:3000/api/buyer/profile/upload-picture" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -F "photo=@/path/to/image.jpg"
```

**Validation:**
- File size: Max 2MB
- Allowed formats: JPG, JPEG, PNG
- Image is automatically optimized and resized to 400x400px

---

### 4. Delete Profile Picture
Remove buyer profile picture.

**Endpoint:** `DELETE /api/buyer/profile/delete-picture`

**Headers:**
```http
Authorization: Bearer <token>
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Profile photo deleted successfully",
  "data": null
}
```

**cURL Example:**
```bash
curl -X DELETE "http://localhost:3000/api/buyer/profile/delete-picture" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

### 5. Change Password
Change buyer account password.

**Endpoint:** `PUT /api/buyer/profile/change-password`

**Headers:**
```http
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "current_password": "OldPassword123",
  "new_password": "NewPassword123",
  "confirm_password": "NewPassword123"
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Password changed successfully",
  "data": null
}
```

**cURL Example:**
```bash
curl -X PUT "http://localhost:3000/api/buyer/profile/change-password" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "current_password": "OldPassword123",
    "new_password": "NewPassword123",
    "confirm_password": "NewPassword123"
  }'
```

**Validation:**
- Current password must be correct
- New password must be at least 8 characters
- New password and confirm password must match

---

## Address Management

### 6. Get All Addresses
Retrieve all saved addresses for the buyer.

**Endpoint:** `GET /api/buyer/addresses`

**Headers:**
```http
Authorization: Bearer <token>
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Addresses retrieved successfully",
  "data": [
    {
      "id": 1,
      "buyer_id": 1,
      "address_type": "home",
      "street_address": "123 Main St",
      "city": "New York",
      "state": "NY",
      "postal_code": "10001",
      "country": "USA",
      "is_default": true,
      "created_at": "2024-01-01T00:00:00Z"
    },
    {
      "id": 2,
      "buyer_id": 1,
      "address_type": "work",
      "street_address": "456 Office Blvd",
      "city": "New York",
      "state": "NY",
      "postal_code": "10002",
      "country": "USA",
      "is_default": false,
      "created_at": "2024-01-05T00:00:00Z"
    }
  ]
}
```

**cURL Example:**
```bash
curl -X GET "http://localhost:3000/api/buyer/addresses" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

### 7. Create Address
Add a new address for the buyer.

**Endpoint:** `POST /api/buyer/addresses`

**Headers:**
```http
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "address_type": "home",
  "street_address": "123 Main St",
  "city": "New York",
  "state": "NY",
  "postal_code": "10001",
  "country": "USA",
  "is_default": true
}
```

**Response:** `201 Created`
```json
{
  "success": true,
  "message": "Address created successfully",
  "data": {
    "id": 1,
    "buyer_id": 1,
    "address_type": "home",
    "street_address": "123 Main St",
    "city": "New York",
    "state": "NY",
    "postal_code": "10001",
    "country": "USA",
    "is_default": true,
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

**cURL Example:**
```bash
curl -X POST "http://localhost:3000/api/buyer/addresses" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "address_type": "home",
    "street_address": "123 Main St",
    "city": "New York",
    "state": "NY",
    "postal_code": "10001",
    "country": "USA",
    "is_default": true
  }'
```

---

### 8. Update Address
Update an existing address.

**Endpoint:** `PUT /api/buyer/addresses/{id}`

**Headers:**
```http
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "address_type": "home",
  "street_address": "789 New St",
  "city": "Boston",
  "state": "MA",
  "postal_code": "02101",
  "country": "USA",
  "is_default": false
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Address updated successfully",
  "data": {
    "id": 1,
    "street_address": "789 New St",
    "city": "Boston",
    "state": "MA",
    "postal_code": "02101",
    "country": "USA"
  }
}
```

**cURL Example:**
```bash
curl -X PUT "http://localhost:3000/api/buyer/addresses/1" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "street_address": "789 New St",
    "city": "Boston",
    "state": "MA",
    "postal_code": "02101"
  }'
```

---

### 9. Delete Address
Delete a saved address.

**Endpoint:** `DELETE /api/buyer/addresses/{id}`

**Headers:**
```http
Authorization: Bearer <token>
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Address deleted successfully",
  "data": null
}
```

**cURL Example:**
```bash
curl -X DELETE "http://localhost:3000/api/buyer/addresses/1" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

## Order History

### 10. Get Order History
Retrieve buyer's order history with pagination.

**Endpoint:** `GET /api/buyer/orders`

**Headers:**
```http
Authorization: Bearer <token>
```

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10, max: 50)
- `status` (optional): Filter by order status (pending, confirmed, shipped, delivered, cancelled)

**Response:** `200 OK`
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "order_number": "ORD-20240101-001",
      "total_amount": 299.99,
      "status": "delivered",
      "payment_status": "paid",
      "created_at": "2024-01-01T00:00:00Z",
      "items": [
        {
          "product_name": "Oak Tree",
          "quantity": 2,
          "price": 149.99
        }
      ]
    }
  ],
  "page": 1,
  "limit": 10,
  "total": 25
}
```

**cURL Example:**
```bash
# Get all orders
curl -X GET "http://localhost:3000/api/buyer/orders?page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"

# Filter by status
curl -X GET "http://localhost:3000/api/buyer/orders?status=delivered" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

### 11. Get Buyer Statistics
Retrieve buyer account statistics.

**Endpoint:** `GET /api/buyer/stats`

**Headers:**
```http
Authorization: Bearer <token>
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Statistics retrieved successfully",
  "data": {
    "total_orders": 25,
    "pending_orders": 2,
    "completed_orders": 20,
    "cancelled_orders": 3,
    "total_spent": 2499.75,
    "saved_addresses": 3,
    "wishlist_items": 5
  }
}
```

**cURL Example:**
```bash
curl -X GET "http://localhost:3000/api/buyer/stats" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

## Error Responses

### 400 Bad Request
```json
{
  "success": false,
  "error": "Invalid request data"
}
```

### 401 Unauthorized
```json
{
  "success": false,
  "error": "Authentication required"
}
```

### 403 Forbidden
```json
{
  "success": false,
  "error": "Access denied"
}
```

### 404 Not Found
```json
{
  "success": false,
  "error": "Resource not found"
}
```

### 500 Internal Server Error
```json
{
  "success": false,
  "error": "Internal server error"
}
```

---

## Status Codes

| Code | Description |
|------|-------------|
| 200  | Success |
| 201  | Created |
| 400  | Bad Request |
| 401  | Unauthorized |
| 403  | Forbidden |
| 404  | Not Found |
| 500  | Internal Server Error |

---

## Data Validation Rules

### Profile Update
- `first_name`: Required, 2-50 characters
- `last_name`: Required, 2-50 characters
- `phone`: Optional, valid phone format
- `date_of_birth`: Optional, valid date format (YYYY-MM-DD)
- `gender`: Optional, enum (male, female, other)

### Address
- `address_type`: Required, enum (home, work, other)
- `street_address`: Required, 5-200 characters
- `city`: Required, 2-100 characters
- `state`: Required, 2-100 characters
- `postal_code`: Required, 3-20 characters
- `country`: Required, 2-100 characters

### Password Change
- `current_password`: Required
- `new_password`: Required, min 8 characters, must contain uppercase, lowercase, and number
- `confirm_password`: Required, must match new_password

### Profile Photo
- Max file size: 2MB
- Allowed formats: JPG, JPEG, PNG
- Automatically resized to 400x400px
- Stored in Cloudinary

---

## Testing Workflow

### 1. Get Profile
```bash
curl -X GET "http://localhost:3000/api/buyer/profile" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### 2. Update Profile
```bash
curl -X PUT "http://localhost:3000/api/buyer/profile" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"first_name":"John","last_name":"Doe"}'
```

### 3. Upload Photo
```bash
curl -X POST "http://localhost:3000/api/buyer/profile/upload-picture" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -F "photo=@profile.jpg"
```

### 4. Add Address
```bash
curl -X POST "http://localhost:3000/api/buyer/addresses" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"address_type":"home","street_address":"123 Main St","city":"NYC","state":"NY","postal_code":"10001","country":"USA"}'
```

### 5. View Orders
```bash
curl -X GET "http://localhost:3000/api/buyer/orders" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

---

## Notes

- All timestamps are in UTC format
- Profile photos are stored in Cloudinary with automatic optimization
- Default address is automatically set if it's the first address
- Order history includes pagination for better performance
- All endpoints require valid JWT authentication
- Rate limiting may apply to prevent abuse

---

## Support

For issues or questions, contact: support@treestore.com
