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
    "id": 61,
    "user_id": 134,
    "phone": "01819240089",
    "profile_picture": "",
    "default_address": "",
    "is_active": true,
    "email_verified": false,
    "last_order_at": null,
    "total_orders_count": 0,
    "total_spent": 0,
    "created_at": "2025-11-14T22:15:53.964959+06:00",
    "updated_at": "2025-11-16T22:19:57.055216+06:00",
    "reg_user": {
        "first_name": "nasir ",
        "last_name": "sheikh",
        "email": "nasirsheikh@gmail.com"
    },
    "default_addr": {
        "id": 42,
        "full_name": "",
        "address_line_1": "janpur bankpara",
        "address_line_2": "",
        "street": "",
        "city": "sirajganj",
        "state": "",
        "district": "",
        "country": "Bangladesh",
        "postal_code": "5700",
        "phone": "",
        "is_default": true
    },
    "addresses": [
        {
            "id": 42,
            "full_name": "saad ahmed",
            "address_line_1": "janpur bankpara",
            "address_line_2": "",
            "street": "",
            "city": "sirajganj",
            "state": "",
            "district": "",
            "country": "Bangladesh",
            "postal_code": "5700",
            "phone": "01999999999",
            "is_default": true
        },
        {
            "id": 41,
            "full_name": "saad ahmed",
            "address_line_1": "janpur bankpara",
            "address_line_2": "sirajganj",
            "street": "",
            "city": "sirajganj",
            "state": "rajshahi Division",
            "district": "",
            "country": "Bangladesh",
            "postal_code": "5700",
            "phone": "01892444",
            "is_default": false
        },
        {
            "id": 40,
            "full_name": "saad ahmed",
            "address_line_1": "janpur bankpara",
            "address_line_2": "sirajganj",
            "street": "",
            "city": "sirajganj",
            "state": "rajshahi Division",
            "district": "",
            "country": "Bangladesh",
            "postal_code": "5700",
            "phone": "01892444",
            "is_default": false
        },
        {
            "id": 39,
            "full_name": "",
            "address_line_1": "",
            "address_line_2": "",
            "street": "",
            "city": "sirajagnj",
            "state": "rajshahi",
            "district": "",
            "country": "Bangladesh",
            "postal_code": "5700",
            "phone": "",
            "is_default": false
        }
    ],
    "wishlist_count": 0,
    "cart_item_count": 0
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
  "phone": "01819240089",
  "first_name": "nasir",
  "last_name": "Ahmed",
  "gender": "male",
  "date_of_birth": "2002-05-14T00:00:00.000Z"
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
    "current_password":"password123",
    "new_password":"password1234",
    "confirm_password":"password1234"
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

 [
    {
        "id": 43,
        "full_name": "saad ahmed",
        "address_line_1": "janpur bankpara",
        "address_line_2": "sirajganj",
        "street": "janpur bankpara",
        "city": "sirajganj",
        "state": "rajshahi Division",
        "district": "rajshahi",
        "country": "Bangladesh",
        "postal_code": "5700",
        "phone": "01892444",
        "is_default": true
    },
 ]
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
"full_name":"saad ahmed",
"address_line_1": "janpur bankpara",
  "address_line_2": "sirajganj",
  "street": "janpur bankpara",
  "city": "sirajganj",
  "district":"rajshahi",
  "country": "Bangladesh",
  "postal_code": "5700",
  "phone":"01892444",
  "is_default": true
}
```

**Response:** `201 Created`
```json
{
    "id": 43,
    "full_name": "saad ahmed",
    "address_line_1": "janpur bankpara",
    "address_line_2": "sirajganj",
    "street": "janpur bankpara",
    "city": "sirajganj",
    "state": "rajshahi Division",
    "district": "rajshahi",
    "country": "Bangladesh",
    "postal_code": "5700",
    "phone": "01892444",
    "is_default": true
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
        "address_line1": "ss road",
        "address_line2": "sirajganj sadar",
        "street": "janpur bankpara",
        "city": "sirajganj",
        "state": "rajshahi Division",
        "district":"rajshahi",
        "country": "Bangladesh",
        "postal_code": "5700",
        "is_default": false
}
```

**Response:** `200 OK`
```json
{
    "id": 43,
    "full_name": "saad ahmed",
    "address_line_1": "janpur bankpara",
    "address_line_2": "sirajganj",
    "street": "janpur bankpara",
    "city": "sirajganj",
    "state": "rajshahi Division",
    "district": "rajshahi",
    "country": "Bangladesh",
    "postal_code": "5700",
    "phone": "",
    "is_default": true
}
```

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
  
  "message": "Address deleted successfully",
  
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
  {
    "orders": [
        {
            "id": 20,
            "order_number": "ORD-1763360676-61",
            "total_amount": 419.98,
            "status": "pending",
            "item_count": 1,
            "sub_total": 0,
            "order_date": "2025-11-17T12:24:36.87553+06:00",
            "items": [
                {
                    "product_id": 40,
                    "product_name": "Bonsai",
                    "quantity": 1,
                    "price": 500,
                    "subtotal": 0
                }
            ]
        },
  }
   "pagination": {
        "page": 1,
        "limit": 10,
        "total": 11,
        "total_pages": 2,
        "has_next": true,
        "has_prev": false
    }
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
  "total_orders": 0,
    "pending_orders": 6,
    "completed_orders": 0,
    "cancelled_orders": 5,
    "total_spent": 0,
    "wishlist_count": 4,
    "cart_item_count": 0,
    "review_count": 0,
    "last_order_date": null,
    "recent_orders": [
        {
            "order_id": 20,
            "order_number": "ORD-1763360676-61",
            "total_amount": 419.98,
            "status": "pending",
            "item_count": 1,
            "order_date": "2025-11-17T12:24:36.87553+06:00"
        },
        {
            "order_id": 19,
            "order_number": "ORD-1763360352-61",
            "total_amount": 520,
            "status": "pending",
            "item_count": 1,
            "order_date": "2025-11-17T12:19:12.991869+06:00"
        },
        {
            "order_id": 18,
            "order_number": "ORD-1763359562-61",
            "total_amount": 520,
            "status": "cancelled",
            "item_count": 1,
            "order_date": "2025-11-17T12:06:02.514207+06:00"
        },
        {
            "order_id": 17,
            "order_number": "ORD-1763359124-61",
            "total_amount": 675,
            "status": "pending",
            "item_count": 1,
            "order_date": "2025-11-17T11:58:44.918772+06:00"
        },
        {
            "order_id": 16,
            "order_number": "ORD-1763358864-61",
            "total_amount": 520,
            "status": "pending",
            "item_count": 1,
            "order_date": "2025-11-17T11:54:24.996392+06:00"
        }
    ]
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
