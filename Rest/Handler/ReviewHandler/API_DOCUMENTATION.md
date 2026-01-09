# Review API Documentation

## Overview
This API allows buyers to create, read, update, and delete product reviews. All buyer endpoints require JWT authentication.

---

## Endpoints

### 1. Check if Buyer Can Review Product
**Endpoint:** `GET /api/buyer/review/can-review`  
**Auth Required:** Yes (JWT Token)  
**Description:** Check if the authenticated buyer can review a specific product.

#### Query Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| product_id | uint | Yes | ID of the product to review |

#### Request Example
```http
GET /api/buyer/review/can-review?product_id=9
Authorization: Bearer <your_jwt_token>
```

#### Response Example (200 OK)
```json
{
  "success": true,
  "message": "Review Created successfully",
  "data": {
    "can_review": true,
    "has_purchased": true,
    "already_reviewed": false,
    "order_id": 91,
    "message": "You can review this product"
  }
}
```

#### Response Example - Cannot Review (200 OK)
```json
{
  "success": true,
  "data": {
    "can_review": false,
    "has_purchased": false,
    "already_reviewed": false,
    "order_id": null,
    "message": "You can only review products you have purchased and received"
  }
}
```

---

### 2. Create Product Review
**Endpoint:** `POST /api/buyer/reviews`  
**Auth Required:** Yes (JWT Token)  
**Description:** Create a new review for a purchased and delivered product.

#### Request Body
```json
{
  "product_id": 9,
  "order_id": 91,
  "rating": 5,
  "title": "Great product!",
  "comment": "This product exceeded my expectations. Highly recommended!",
  "images": [
    "https://example.com/review-image-1.jpg",
    "https://example.com/review-image-2.jpg"
  ]
}
```

#### Field Validation
| Field | Type | Required | Validation |
|-------|------|----------|------------|
| product_id | uint | Yes | Must exist |
| order_id | uint | Yes | Must exist |
| rating | int | Yes | 1-5 |
| title | string | Yes | 3-255 characters |
| comment | string | Yes | 10-2000 characters |
| images | []string | No | Array of image URLs |

#### Response Example (200 OK)
```json
{
  "success": true,
  "message": "Review Created successfully",
  "data": {
    "id": 123,
    "product_id": 9,
    "buyer_id": 45,
    "order_id": 91,
    "rating": 5,
    "title": "Great product!",
    "comment": "This product exceeded my expectations. Highly recommended!",
    "images": [
      {
        "id": 1,
        "review_id": 123,
        "image_url": "https://example.com/review-image-1.jpg",
        "alt_text": "Review image 1"
      }
    ],
    "buyer": {
      "id": 45,
      "name": "John Doe",
      "avatar": "https://example.com/avatar.jpg"
    },
    "product": {
      "id": 9,
      "name": "Product Name",
      "image": "https://example.com/product.jpg"
    },
    "created_at": "2025-12-30T23:37:45Z",
    "updated_at": "2025-12-30T23:37:45Z"
  }
}
```

#### Error Responses
```json
// Already reviewed
{
  "success": false,
  "error": "you have already reviewed this product"
}

// Not purchased or not delivered
{
  "success": false,
  "error": "you can only review products you have purchased"
}

// Validation error
{
  "success": false,
  "error": "rating must be between 1 and 5"
}
```

---

### 3. Get Product Reviews (Public)
**Endpoint:** `GET /api/reviews/product`  
**Auth Required:** No  
**Description:** Get all reviews for a specific product with pagination and filtering.

#### Query Parameters
| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| product_id | uint | Yes | - | Product ID |
| rating | int | No | - | Filter by rating (1-5) |
| page | int | No | 1 | Page number |
| per_page | int | No | 10 | Items per page |
| sort_by | string | No | newest | Sort order: newest, oldest, highest_rated, lowest_rated |

#### Request Example
```http
GET /api/reviews/product?product_id=9&rating=5&page=1&per_page=10&sort_by=newest
```

#### Response Example (200 OK)
```json
{
  "success": true,
  "data": {
    "reviews": [
      {
        "id": 123,
        "product_id": 9,
        "rating": 5,
        "title": "Great product!",
        "comment": "This product exceeded my expectations.",
        "images": [...],
        "buyer": {...},
        "created_at": "2025-12-30T23:37:45Z"
      }
    ],
    "pagination": {
      "page": 1,
      "per_page": 10,
      "total": 50,
      "total_pages": 5
    }
  }
}
```

---

### 4. Get My Reviews
**Endpoint:** `GET /api/buyer/reviews/my-reviews`  
**Auth Required:** Yes (JWT Token)  
**Description:** Get all reviews created by the authenticated buyer.

#### Query Parameters
| Parameter | Type | Required | Default |
|-----------|------|----------|---------|
| page | int | No | 1 |
| per_page | int | No | 10 |

#### Request Example
```http
GET /api/buyer/reviews/my-reviews?page=1&per_page=10
Authorization: Bearer <your_jwt_token>
```

#### Response Example (200 OK)
```json
{
  "success": true,
  "data": {
    "reviews": [
      {
        "id": 123,
        "product_id": 9,
        "rating": 5,
        "title": "Great product!",
        "comment": "This product exceeded my expectations.",
        "product": {
          "id": 9,
          "name": "Product Name",
          "image": "https://example.com/product.jpg"
        },
        "created_at": "2025-12-30T23:37:45Z"
      }
    ],
    "pagination": {...}
  }
}
```

---

### 5. Update Review
**Endpoint:** `PUT /api/buyer/reviews/{id}`  
**Auth Required:** Yes (JWT Token)  
**Description:** Update an existing review. Only the review owner can update.

#### URL Parameters
| Parameter | Type | Description |
|-----------|------|-------------|
| id | uint | Review ID |

#### Request Body
```json
{
  "rating": 4,
  "title": "Updated title",
  "comment": "Updated comment with more details about the product.",
  "images": [
    "https://example.com/new-image.jpg"
  ]
}
```

#### Request Example
```http
PUT /api/buyer/reviews/123
Authorization: Bearer <your_jwt_token>
Content-Type: application/json

{
  "rating": 4,
  "title": "Updated title",
  "comment": "Updated comment with more details."
}
```

#### Response Example (200 OK)
```json
{
  "success": true,
  "message": "Review updated successfully",
  "data": {
    "id": 123,
    "rating": 4,
    "title": "Updated title",
    "comment": "Updated comment with more details.",
    "updated_at": "2025-12-30T23:45:00Z"
  }
}
```

---

### 6. Delete Review
**Endpoint:** `DELETE /api/buyer/reviews/{id}`  
**Auth Required:** Yes (JWT Token)  
**Description:** Delete a review. Only the review owner can delete.

#### URL Parameters
| Parameter | Type | Description |
|-----------|------|-------------|
| id | uint | Review ID |

#### Request Example
```http
DELETE /api/buyer/reviews/123
Authorization: Bearer <your_jwt_token>
```

#### Response Example (200 OK)
```json
{
  "success": true,
  "message": "Review deleted successfully"
}
```

#### Error Response (403 Forbidden)
```json
{
  "success": false,
  "error": "you can only delete your own reviews"
}
```

---

## Frontend Integration Guide

### 1. Check Before Showing Review Form
Before showing the review form, check if the user can review:

```javascript
async function checkCanReview(productId) {
  const response = await fetch(
    `/api/buyer/review/can-review?product_id=${productId}`,
    {
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('token')}`
      }
    }
  );
  const result = await response.json();
  
  if (result.data.can_review) {
    // Show review form
    showReviewForm(result.data.order_id);
  } else {
    // Show message: result.data.message
    showMessage(result.data.message);
  }
}
```

### 2. Submit Review
```javascript
async function submitReview(reviewData) {
  const response = await fetch('/api/buyer/reviews', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${localStorage.getItem('token')}`
    },
    body: JSON.stringify({
      product_id: reviewData.productId,
      order_id: reviewData.orderId,
      rating: reviewData.rating,
      title: reviewData.title,
      comment: reviewData.comment,
      images: reviewData.images // Optional
    })
  });
  
  const result = await response.json();
  
  if (result.success) {
    // Review created successfully
    console.log('Review created:', result.data);
  } else {
    // Show error
    alert(result.error);
  }
}
```

### 3. Load Product Reviews
```javascript
async function loadProductReviews(productId, page = 1) {
  const response = await fetch(
    `/api/reviews/product?product_id=${productId}&page=${page}&per_page=10&sort_by=newest`
  );
  const result = await response.json();
  
  if (result.success) {
    displayReviews(result.data.reviews);
    displayPagination(result.data.pagination);
  }
}
```

### 4. Update Review
```javascript
async function updateReview(reviewId, updatedData) {
  const response = await fetch(`/api/buyer/reviews/${reviewId}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${localStorage.getItem('token')}`
    },
    body: JSON.stringify(updatedData)
  });
  
  const result = await response.json();
  return result;
}
```

### 5. Delete Review
```javascript
async function deleteReview(reviewId) {
  if (!confirm('Are you sure you want to delete this review?')) {
    return;
  }
  
  const response = await fetch(`/api/buyer/reviews/${reviewId}`, {
    method: 'DELETE',
    headers: {
      'Authorization': `Bearer ${localStorage.getItem('token')}`
    }
  });
  
  const result = await response.json();
  
  if (result.success) {
    // Remove review from UI
    removeReviewFromUI(reviewId);
  }
}
```

---

## Important Notes

### Authentication
- All buyer endpoints require JWT token in Authorization header
- Token format: `Bearer <token>`
- Get token from login/register response

### Business Rules
1. **Can only review delivered products**: Order status must be "delivered"
2. **One review per product**: Cannot review the same product twice
3. **Must be purchased**: Can only review products you've actually bought
4. **Owner only**: Can only update/delete your own reviews

### Error Handling
Always check the `success` field in the response:
```javascript
if (response.success) {
  // Handle success
} else {
  // Handle error: response.error
}
```

### Image Upload
Images should be uploaded separately to your image storage service first, then include the URLs in the `images` array.

---

## Common Error Codes

| Status Code | Description |
|-------------|-------------|
| 200 | Success |
| 400 | Bad Request (validation error) |
| 401 | Unauthorized (missing/invalid token) |
| 403 | Forbidden (not owner) |
| 404 | Not Found |
| 500 | Internal Server Error |


---

## Seller Endpoints

### 7. Get Reviews for Seller's Products
**Endpoint:** `GET /api/seller/reviews/my-products`  
**Auth Required:** Yes (JWT Token - Seller)  
**Description:** Get all reviews for products owned by the authenticated seller.

#### Query Parameters
| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| product_id | uint | No | - | Filter by specific product |
| page | int | No | 1 | Page number |
| per_page | int | No | 10 | Items per page (max 100) |

#### Request Example
```http
GET /api/seller/reviews/my-products?page=1&per_page=10
Authorization: Bearer <seller_jwt_token>
```

#### Request Example - Filter by Product
```http
GET /api/seller/reviews/my-products?product_id=9&page=1&per_page=10
Authorization: Bearer <seller_jwt_token>
```

#### Response Example (200 OK)
```json
{
  "success": true,
  "message": "Reviews fetched successfully",
  "data": {
    "reviews": [
      {
        "id": 123,
        "product_id": 9,
        "buyer_id": 45,
        "rating": 5,
        "title": "Great product!",
        "comment": "This product exceeded my expectations.",
        "seller_response": "Thank you for your feedback!",
        "seller_responded_at": "2025-12-31T10:30:00Z",
        "images": [...],
        "buyer": {
          "id": 45,
          "name": "John Doe",
          "avatar": "https://example.com/avatar.jpg"
        },
        "product": {
          "id": 9,
          "name": "Product Name",
          "image": "https://example.com/product.jpg"
        },
        "created_at": "2025-12-30T23:37:45Z",
        "updated_at": "2025-12-31T10:30:00Z"
      }
    ],
    "pagination": {
      "page": 1,
      "per_page": 10,
      "total": 25,
      "total_pages": 3
    }
  }
}
```

---

### 8. Respond to Review
**Endpoint:** `POST /api/seller/reviews/{id}/respond`  
**Auth Required:** Yes (JWT Token - Seller)  
**Description:** Add a seller response to a review on your product.

#### URL Parameters
| Parameter | Type | Description |
|-----------|------|-------------|
| id | uint | Review ID |

#### Request Body
```json
{
  "response": "Thank you for your feedback! We're glad you enjoyed our product."
}
```

#### Field Validation
| Field | Type | Required | Validation |
|-------|------|----------|------------|
| response | string | Yes | 10-1000 characters |

#### Request Example
```http
POST /api/seller/reviews/123/respond
Authorization: Bearer <seller_jwt_token>
Content-Type: application/json

{
  "response": "Thank you for your feedback! We're glad you enjoyed our product."
}
```

#### Response Example (200 OK)
```json
{
  "success": true,
  "message": "Response added successfully",
  "data": {
    "id": 123,
    "product_id": 9,
    "rating": 5,
    "title": "Great product!",
    "comment": "This product exceeded my expectations.",
    "seller_response": "Thank you for your feedback! We're glad you enjoyed our product.",
    "seller_responded_at": "2025-12-31T10:30:00Z",
    "buyer": {...},
    "product": {...},
    "created_at": "2025-12-30T23:37:45Z",
    "updated_at": "2025-12-31T10:30:00Z"
  }
}
```

#### Error Responses
```json
// Not the product owner
{
  "success": false,
  "error": "you can only respond to reviews on your own products"
}

// Review not found
{
  "success": false,
  "error": "review not found"
}

// Validation error
{
  "success": false,
  "error": "response must be between 10 and 1000 characters"
}
```

---

## Seller Frontend Integration Guide

### 1. Load Seller's Product Reviews
```javascript
async function loadSellerReviews(page = 1, productId = null) {
  let url = `/api/seller/reviews/my-products?page=${page}&per_page=10`;
  
  if (productId) {
    url += `&product_id=${productId}`;
  }
  
  const response = await fetch(url, {
    headers: {
      'Authorization': `Bearer ${localStorage.getItem('sellerToken')}`
    }
  });
  
  const result = await response.json();
  
  if (result.success) {
    displaySellerReviews(result.data.reviews);
    displayPagination(result.data.pagination);
  }
}
```

### 2. Respond to Review
```javascript
async function respondToReview(reviewId, responseText) {
  const response = await fetch(`/api/seller/reviews/${reviewId}/respond`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${localStorage.getItem('sellerToken')}`
    },
    body: JSON.stringify({
      response: responseText
    })
  });
  
  const result = await response.json();
  
  if (result.success) {
    // Update UI with seller response
    updateReviewWithResponse(reviewId, result.data);
    alert('Response added successfully!');
  } else {
    alert(result.error);
  }
}
```

### 3. Filter Reviews by Product
```javascript
async function filterReviewsByProduct(productId) {
  await loadSellerReviews(1, productId);
}
```

---

## Complete Endpoint Summary

### Public Endpoints
- `GET /api/reviews/product` - Get product reviews (no auth)

### Buyer Endpoints (Require Buyer JWT)
- `GET /api/buyer/review/can-review` - Check if can review
- `POST /api/buyer/reviews` - Create review
- `GET /api/buyer/reviews/my-reviews` - Get my reviews
- `PUT /api/buyer/reviews/{id}` - Update review
- `DELETE /api/buyer/reviews/{id}` - Delete review

### Seller Endpoints (Require Seller JWT)
- `GET /api/seller/reviews/my-products` - Get reviews for seller's products
- `POST /api/seller/reviews/{id}/respond` - Respond to review

---

## Notes for Sellers

### Business Rules
1. **Product Ownership**: Can only respond to reviews on your own products
2. **Response Length**: Response must be 10-1000 characters
3. **Update Response**: Can update response by calling the endpoint again with new text
4. **View All Reviews**: See all reviews across all your products or filter by specific product

### Best Practices
1. **Respond Promptly**: Reply to reviews within 24-48 hours
2. **Be Professional**: Keep responses courteous and helpful
3. **Address Concerns**: If negative review, acknowledge and offer solutions
4. **Thank Customers**: Show appreciation for positive feedback
5. **Monitor Regularly**: Check for new reviews daily


---

## Admin Endpoints

### 9. Get All Reviews (Admin)
**Endpoint:** `GET /api/admin/reviews`  
**Auth Required:** Yes (JWT Token - Admin/Super Admin)  
**Description:** Get all reviews with filtering, search, and pagination.

#### Query Parameters
| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| page | int | No | 1 | Page number |
| limit | int | No | 20 | Items per page |
| status | string | No | - | Filter by status: pending, approved, rejected |
| rating | int | No | - | Filter by rating (1-5) |
| is_reported | bool | No | - | Filter reported reviews |
| search | string | No | - | Search in title and comment |

#### Request Example
```http
GET /api/admin/reviews?page=1&limit=20&status=pending&rating=1
Authorization: Bearer <admin_jwt_token>
```

#### Response Example (200 OK)
```json
{
  "status": "success",
  "message": "Reviews retrieved successfully",
  "data": [
    {
      "id": 1,
      "product_id": 10,
      "product_name": "Product Name",
      "product_image": "https://example.com/product.jpg",
      "buyer_id": 5,
      "buyer_name": "John Doe",
      "buyer_avatar": "https://example.com/avatar.jpg",
      "seller_id": 3,
      "seller_name": "Tech Store",
      "rating": 1,
      "title": "Poor quality",
      "comment": "Not as described",
      "status": "pending",
      "is_verified_purchase": true,
      "helpful_count": 0,
      "report_count": 3,
      "is_reported": true,
      "created_at": "2025-12-30T10:00:00Z",
      "images": ["https://example.com/review1.jpg"]
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 20,
    "total": 150,
    "total_pages": 8
  }
}
```

---

### 10. Get Pending Reviews (Admin)
**Endpoint:** `GET /api/admin/reviews/pending`  
**Auth Required:** Yes (JWT Token - Admin/Super Admin)  
**Description:** Get all reviews awaiting moderation.

#### Query Parameters
| Parameter | Type | Required | Default |
|-----------|------|----------|---------|
| page | int | No | 1 |
| limit | int | No | 20 |

#### Request Example
```http
GET /api/admin/reviews/pending?page=1&limit=20
Authorization: Bearer <admin_jwt_token>
```

---

### 11. Get Reported Reviews (Admin)
**Endpoint:** `GET /api/admin/reviews/reported`  
**Auth Required:** Yes (JWT Token - Admin/Super Admin)  
**Description:** Get all reviews that have been reported by users.

#### Query Parameters
| Parameter | Type | Required | Default |
|-----------|------|----------|---------|
| page | int | No | 1 |
| limit | int | No | 20 |

#### Request Example
```http
GET /api/admin/reviews/reported?page=1&limit=20
Authorization: Bearer <admin_jwt_token>
```

---

### 12. Moderate Review (Admin)
**Endpoint:** `POST /api/admin/reviews/moderate`  
**Auth Required:** Yes (JWT Token - Admin/Super Admin)  
**Description:** Approve or reject a review.

#### Query Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | uint | Yes | Review ID |

#### Request Body
```json
{
  "action": "approve",
  "reason": "Spam content"
}
```

#### Field Validation
| Field | Type | Required | Validation |
|-------|------|----------|------------|
| action | string | Yes | "approve" or "reject" |
| reason | string | No | Required only for reject action |

#### Request Example - Approve
```http
POST /api/admin/reviews/moderate?id=5
Authorization: Bearer <admin_jwt_token>
Content-Type: application/json

{
  "action": "approve"
}
```

#### Request Example - Reject
```http
POST /api/admin/reviews/moderate?id=5
Authorization: Bearer <admin_jwt_token>
Content-Type: application/json

{
  "action": "reject",
  "reason": "Contains inappropriate content"
}
```

#### Response Example (200 OK)
```json
{
  "status": "success",
  "message": "Review approved successfully"
}
```

---

### 13. Delete Review (Admin)
**Endpoint:** `DELETE /api/admin/reviews/delete`  
**Auth Required:** Yes (JWT Token - Admin/Super Admin)  
**Description:** Permanently delete a review.

#### Query Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | uint | Yes | Review ID |

#### Request Example
```http
DELETE /api/admin/reviews/delete?id=5
Authorization: Bearer <admin_jwt_token>
```

#### Response Example (200 OK)
```json
{
  "status": "success",
  "message": "Review deleted successfully"
}
```

---

### 14. Get Review Statistics (Admin)
**Endpoint:** `GET /api/admin/stats`  
**Auth Required:** Yes (JWT Token - Admin/Super Admin)  
**Description:** Get comprehensive review statistics.

#### Request Example
```http
GET /api/admin/stats
Authorization: Bearer <admin_jwt_token>
```

#### Response Example (200 OK)
```json
{
  "status": "success",
  "data": {
    "total": 500,
    "pending": 25,
    "approved": 450,
    "rejected": 15,
    "reported": 10,
    "average_rating": 4.2,
    "rating_distribution": {
      "1_star": 10,
      "2_star": 20,
      "3_star": 50,
      "4_star": 150,
      "5_star": 220
    }
  }
}
```

---

### 15. Report Review (Any User)
**Endpoint:** `POST /api/admin/reviews/report`  
**Auth Required:** Yes (JWT Token)  
**Description:** Report a review for inappropriate content.

#### Query Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| id | uint | Yes | Review ID |

#### Request Body
```json
{
  "reason": "spam",
  "details": "This is clearly a fake review"
}
```

#### Field Validation
| Field | Type | Required | Validation |
|-------|------|----------|------------|
| reason | string | Yes | spam, inappropriate, fake, offensive, other |
| details | string | No | Additional details |

#### Request Example
```http
POST /api/admin/reviews/report?id=5
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "reason": "spam",
  "details": "This review is fake"
}
```

#### Response Example (200 OK)
```json
{
  "status": "success",
  "message": "Review reported successfully"
}
```

---

### 16. Get Review Reports (Admin)
**Endpoint:** `POST /api/admin/reviews/reports`  
**Auth Required:** Yes (JWT Token - Admin/Super Admin)  
**Description:** Get all review reports with filtering.

#### Query Parameters
| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| status | string | No | - | pending, reviewed, dismissed |
| page | int | No | 1 | Page number |
| limit | int | No | 20 | Items per page |

#### Request Example
```http
POST /api/admin/reviews/reports?status=pending&page=1&limit=20
Authorization: Bearer <admin_jwt_token>
```

#### Response Example (200 OK)
```json
{
  "status": "success",
  "data": [
    {
      "id": 10,
      "review_id": 5,
      "review": {
        "id": 5,
        "title": "Bad product",
        "comment": "Fake review content"
      },
      "reporter_type": "buyer",
      "reason": "spam",
      "details": "This is a fake review",
      "status": "pending",
      "created_at": "2025-12-30T10:00:00Z"
    }
  ]
}
```

---

## Admin Frontend Integration Guide

### 1. Get All Reviews with Filters
```javascript
async function loadAdminReviews(filters = {}) {
  const params = new URLSearchParams({
    page: filters.page || 1,
    limit: filters.limit || 20,
    ...(filters.status && { status: filters.status }),
    ...(filters.rating && { rating: filters.rating }),
    ...(filters.is_reported && { is_reported: filters.is_reported }),
    ...(filters.search && { search: filters.search })
  });
  
  const response = await fetch(`/api/admin/reviews?${params}`, {
    headers: {
      'Authorization': `Bearer ${localStorage.getItem('adminToken')}`
    }
  });
  
  const result = await response.json();
  
  if (result.status === 'success') {
    displayReviews(result.data);
    displayPagination(result.pagination);
  }
}
```

### 2. Moderate Review (Approve/Reject)
```javascript
async function moderateReview(reviewId, action, reason = null) {
  const body = { action };
  if (action === 'reject' && reason) {
    body.reason = reason;
  }
  
  const response = await fetch(`/api/admin/reviews/moderate?id=${reviewId}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${localStorage.getItem('adminToken')}`
    },
    body: JSON.stringify(body)
  });
  
  const result = await response.json();
  
  if (result.status === 'success') {
    alert(result.message);
    loadAdminReviews(); // Refresh list
  }
}
```

### 3. Delete Review
```javascript
async function deleteReview(reviewId) {
  if (!confirm('Are you sure you want to permanently delete this review?')) {
    return;
  }
  
  const response = await fetch(`/api/admin/reviews/delete?id=${reviewId}`, {
    method: 'DELETE',
    headers: {
      'Authorization': `Bearer ${localStorage.getItem('adminToken')}`
    }
  });
  
  const result = await response.json();
  
  if (result.status === 'success') {
    alert('Review deleted successfully');
    loadAdminReviews();
  }
}
```

### 4. Get Review Statistics
```javascript
async function loadReviewStats() {
  const response = await fetch('/api/admin/stats', {
    headers: {
      'Authorization': `Bearer ${localStorage.getItem('adminToken')}`
    }
  });
  
  const result = await response.json();
  
  if (result.status === 'success') {
    displayStats(result.data);
  }
}
```

### 5. Report Review (Any User)
```javascript
async function reportReview(reviewId, reason, details) {
  const response = await fetch(`/api/admin/reviews/report?id=${reviewId}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${localStorage.getItem('token')}`
    },
    body: JSON.stringify({ reason, details })
  });
  
  const result = await response.json();
  
  if (result.status === 'success') {
    alert('Review reported successfully');
  }
}
```

---

## Complete Endpoint Summary (Updated)

### Public Endpoints
- `GET /api/reviews/product` - Get product reviews (no auth)

### Buyer Endpoints (Require Buyer JWT)
- `GET /api/buyer/review/can-review` - Check if can review
- `POST /api/buyer/reviews` - Create review
- `GET /api/buyer/reviews/my-reviews` - Get my reviews
- `PUT /api/buyer/reviews/{id}` - Update review
- `DELETE /api/buyer/reviews/{id}` - Delete review

### Seller Endpoints (Require Seller JWT)
- `GET /api/seller/reviews/my-products` - Get reviews for seller's products
- `POST /api/seller/reviews/{id}/respond` - Respond to review

### Admin Endpoints (Require Admin JWT)
- `GET /api/admin/reviews` - Get all reviews with filters
- `GET /api/admin/reviews/pending` - Get pending reviews
- `GET /api/admin/reviews/reported` - Get reported reviews
- `POST /api/admin/reviews/moderate` - Approve/reject review
- `DELETE /api/admin/reviews/delete` - Delete review
- `GET /api/admin/stats` - Get review statistics
- `POST /api/admin/reviews/report` - Report a review (any user)
- `POST /api/admin/reviews/reports` - Get all reports

---

## Admin Review Workflow

### Review Moderation Flow
1. **Buyer creates review** → Status: `pending`
2. **Admin reviews** → Approves or Rejects
3. **If approved** → Review visible to public, updates product rating
4. **If rejected** → Hidden from public, reason stored

### Report Handling Flow
1. **User reports review** → Creates ReviewReport entry
2. **Review flagged** → `is_reported = true`, `report_count++`
3. **Admin reviews report** → Can approve/reject/delete the review
4. **Admin marks report** → Status: `reviewed` or `dismissed`

---

## Admin Best Practices

### Review Moderation
1. **Check Context**: Review product, buyer history, and seller
2. **Look for Patterns**: Multiple reports, suspicious language
3. **Be Fair**: Don't reject negative reviews unless they violate policies
4. **Document Reasons**: Always provide clear rejection reasons
5. **Act Quickly**: Review pending items within 24 hours

### Report Management
1. **Prioritize**: Handle high report counts first
2. **Investigate**: Check both review and reporter history
3. **Take Action**: Approve, reject, or delete based on findings
4. **Communicate**: Update report status after action

### Statistics Monitoring
1. **Track Trends**: Monitor approval/rejection rates
2. **Quality Control**: Watch for spam patterns
3. **User Behavior**: Identify problematic users
4. **Platform Health**: Maintain high-quality review ecosystem
