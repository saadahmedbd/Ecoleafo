# Review API Documentation

## Overview
The Review API allows buyers to create, read, update, and delete product reviews. It also provides endpoints to check review eligibility and retrieve reviews for products.

## Base URL
`/api`

## Authentication
Most endpoints require JWT authentication with buyer role.

---

## Endpoints

### 1. Get Product Reviews
**GET** `/api/reviews/product`

Retrieve all reviews for a specific product (public endpoint).

**Query Parameters:**
- `product_id` (uint, required) - Product ID
- `page` (int, optional, default: 1) - Page number
- `limit` (int, optional, default: 10) - Reviews per page
- `sort` (string, optional) - Sort by: `recent`, `rating_high`, `rating_low`

**Response:**
```json
{
  "status": "success",
  "message": "Reviews retrieved successfully",
  "data": {
    "reviews": [
      {
        "id": 1,
        "product_id": 10,
        "buyer_id": 5,
        "buyer_name": "John Doe",
        "rating": 5,
        "title": "Excellent product!",
        "comment": "Very satisfied with the quality",
        "images": ["https://..."],
        "is_verified_purchase": true,
        "helpful_count": 10,
        "created_at": "2024-01-15T10:30:00Z",
        "updated_at": "2024-01-15T10:30:00Z"
      }
    ],
    "total": 50,
    "page": 1,
    "limit": 10,
    "average_rating": 4.5
  }
}
```

---

### 2. Create Review
**POST** `/api/buyer/reviews`

Create a new product review (requires authentication).

**Request Body:**
```json
{
  "product_id": 10,
  "order_id": 25,
  "rating": 5,
  "title": "Great product",
  "comment": "Highly recommend this product",
  "images": ["https://image1.jpg", "https://image2.jpg"]
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Review created successfully",
  "data": {
    "id": 1,
    "product_id": 10,
    "rating": 5,
    "title": "Great product",
    "comment": "Highly recommend this product",
    "created_at": "2024-01-15T10:30:00Z"
  }
}
```

---

### 3. Get My Reviews
**GET** `/api/buyer/reviews/my-reviews`

Retrieve all reviews created by the authenticated buyer.

**Query Parameters:**
- `page` (int, optional, default: 1) - Page number
- `limit` (int, optional, default: 10) - Reviews per page

**Response:**
```json
{
  "status": "success",
  "message": "Reviews retrieved successfully",
  "data": {
    "reviews": [
      {
        "id": 1,
        "product_id": 10,
        "product_name": "Product Name",
        "rating": 5,
        "title": "Great product",
        "comment": "Highly recommend",
        "created_at": "2024-01-15T10:30:00Z"
      }
    ],
    "total": 5,
    "page": 1,
    "limit": 10
  }
}
```

---

### 4. Check Review Eligibility
**GET** `/api/buyer/review/can-review`

Check if the buyer can review a specific product.

**Query Parameters:**
- `product_id` (uint, required) - Product ID
- `order_id` (uint, optional) - Order ID

**Response:**
```json
{
  "status": "success",
  "message": "Review eligibility checked",
  "data": {
    "can_review": true,
    "reason": "Product purchased and delivered",
    "has_existing_review": false
  }
}
```

---

### 5. Update Review
**PUT** `/api/buyer/reviews/{id}`

Update an existing review.

**Path Parameters:**
- `id` (uint, required) - Review ID

**Request Body:**
```json
{
  "rating": 4,
  "title": "Updated title",
  "comment": "Updated comment",
  "images": ["https://new-image.jpg"]
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Review updated successfully",
  "data": {
    "id": 1,
    "rating": 4,
    "title": "Updated title",
    "comment": "Updated comment",
    "updated_at": "2024-01-16T10:30:00Z"
  }
}
```

---

### 6. Delete Review
**DELETE** `/api/buyer/reviews/{id}`

Delete a review.

**Path Parameters:**
- `id` (uint, required) - Review ID

**Response:**
```json
{
  "status": "success",
  "message": "Review deleted successfully",
  "data": null
}
```

---

## Error Responses

**401 Unauthorized:**
```json
{
  "error": "unauthorized"
}
```

**400 Bad Request:**
```json
{
  "status": "error",
  "message": "Invalid request body",
  "data": null
}
```

**403 Forbidden:**
```json
{
  "status": "error",
  "message": "You can only review products you have purchased",
  "data": null
}
```

**404 Not Found:**
```json
{
  "status": "error",
  "message": "Review not found",
  "data": null
}
```

**500 Internal Server Error:**
```json
{
  "status": "error",
  "message": "Failed to create review",
  "data": null
}
```

---

## Business Rules

### Review Creation
- Buyers can only review products they have purchased
- Order must be delivered before review can be created
- One review per product per buyer
- Rating must be between 1-5
- Title is optional but recommended
- Comment is required (minimum 10 characters)
- Maximum 5 images per review

### Review Update
- Buyers can only update their own reviews
- Can update within 30 days of creation
- Rating, title, comment, and images can be updated

### Review Deletion
- Buyers can only delete their own reviews
- Soft delete - review is marked as deleted but retained in database
- Deletion affects product average rating

---

## Rating System
- **5 stars** - Excellent
- **4 stars** - Good
- **3 stars** - Average
- **2 stars** - Below Average
- **1 star** - Poor

---

## Notes
- All timestamps are in ISO 8601 format (UTC)
- Reviews are publicly visible on product pages
- Verified purchase badge shown for confirmed orders
- Helpful count tracks how many users found the review useful
- Images are optional and stored via CDN
- Reviews can be reported for inappropriate content
