# Product API Documentation

## Overview
Complete API documentation for product management including CRUD operations, image uploads, and product listings.

---

## Base URL
```
http://localhost:3000/api
```

---

## Authentication
Seller endpoints require JWT authentication with seller role.

---

## Endpoints

### 1. Create Product
Create a new product listing.

**Endpoint:** `POST /api/addproducts`

**Headers:**
```http
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Mango Tree",
  "slug": "mango-tree-premium",
  "description": "Premium quality mango tree, 2 years old",
  "sku": "MANGO-001",
  "category_id": 5,
  "price": 299.99,
  "discount_price": 249.99,
  "discount_percent": 16.67,
  "height": "4-5 feet",
  "age": "2 years",
  "tree_type": "fruit",
  "pot_size": "12 inch",
  "scientific_name": "Mangifera indica",
  "common_names": "Mango, Aam",
  "quantity": 50,
  "min_quantity": 5,
  "weight": 15.5,
  "is_active": true,
  "meta_title": "Premium Mango Tree for Sale",
  "meta_description": "Buy premium quality mango tree online"
}
```

**Response:** `201 Created`
```json
{
  "success": true,
  "message": "Product created successfully",
  "data": {
    "id": 1,
    "name": "Mango Tree",
    "slug": "mango-tree-premium",
    "sku": "MANGO-001",
    "price": 299.99,
    "approval_status": "pending",
    "created_at": "2024-01-15T10:30:00Z"
  }
}
```

---

### 2. Get Product by ID
Retrieve detailed product information.

**Endpoint:** `GET /api/products/{productId}`

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "id": 1,
    "seller_id": 10,
    "name": "Mango Tree",
    "slug": "mango-tree-premium",
    "description": "Premium quality mango tree",
    "sku": "MANGO-001",
    "category_id": 5,
    "price": 299.99,
    "discount_price": 249.99,
    "quantity": 50,
    "height": "4-5 feet",
    "age": "2 years",
    "tree_type": "fruit",
    "images": [
      {
        "id": 1,
        "image_url": "https://cloudinary.com/...",
        "is_primary": true
      }
    ],
    "category": {
      "id": 5,
      "name": "Fruit Trees"
    },
    "seller": {
      "id": 10,
      "store_name": "Green Garden Nursery"
    }
  }
}
```

---

### 3. Get Seller Products
Get all products for authenticated seller.

**Endpoint:** `GET /api/seller/products`

**Headers:**
```http
Authorization: Bearer <token>
```

**Query Parameters:**
- `page` (int, optional, default: 1)
- `limit` (int, optional, default: 10)

**Response:** `200 OK`
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Mango Tree",
      "sku": "MANGO-001",
      "price": 299.99,
      "quantity": 50,
      "approval_status": "approved",
      "is_active": true
    }
  ],
  "page": 1,
  "limit": 10,
  "total": 25
}
```

---

### 4. Update Product
Update existing product details.

**Endpoint:** `PUT /api/updateproducts/{productId}`

**Headers:**
```http
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Premium Mango Tree",
  "price": 349.99,
  "quantity": 45,
  "description": "Updated description"
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Product updated successfully",
  "data": {
    "id": 1,
    "name": "Premium Mango Tree",
    "price": 349.99,
    "updated_at": "2024-01-15T11:00:00Z"
  }
}
```

---

### 5. Delete Product
Soft delete a product.

**Endpoint:** `DELETE /api/deleteproducts/{produtcsId}`

**Headers:**
```http
Authorization: Bearer <token>
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Product deleted successfully"
}
```

---

### 6. Upload Product Images
Upload multiple images for a product.

**Endpoint:** `POST /api/products/{ImagesId}/images/multiple`

**Headers:**
```http
Authorization: Bearer <token>
Content-Type: multipart/form-data
```

**Request Body (Form Data):**
- `images` (files): Multiple image files (max 5, each max 5MB)
- `is_primary` (boolean, optional): Set first image as primary

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Images uploaded successfully",
  "data": {
    "uploaded_count": 3,
    "images": [
      {
        "id": 1,
        "image_url": "https://cloudinary.com/...",
        "is_primary": true
      },
      {
        "id": 2,
        "image_url": "https://cloudinary.com/...",
        "is_primary": false
      }
    ]
  }
}
```

---

### 7. Delete Product Image
Remove a specific product image.

**Endpoint:** `DELETE /api/products/{produtcId}/images/{imageId}`

**Headers:**
```http
Authorization: Bearer <token>
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Image deleted successfully"
}
```

---

### 8. Get All Products (Public)
Get paginated list of approved products.

**Endpoint:** `GET /api/products`

**Query Parameters:**
- `page` (int, optional, default: 1)
- `limit` (int, optional, default: 20)
- `category` (int, optional): Filter by category ID
- `min_price` (float, optional): Minimum price filter
- `max_price` (float, optional): Maximum price filter
- `sort` (string, optional): Sort by (price_asc, price_desc, newest, popular)

**Response:** `200 OK`
```json
{
  "success": true,
  "data": [
    {
      "id": 41,
            "seller_id": 62,
            "name": "sundori",
            "slug": "sundori",
            "description": "where you change i say add discount product with original product into the sellerproduct page and prouct detail page",
            "sku": "SUNDORI-7721",
            "category_id": 7,
            "price": 600,
            "discount price": 400,
            "discount_percent": 33.33,
            "height": "6",
            "age": "5",
            "tree_type": "",
            "pot_size": "10",
            "scientific_name": "",
            "common_names": "",
            "quantity": 19,
            "min_quantity": 2,
            "weight": 10,
            "is_active": true,
            "is_approved": false,
            "is_featured": false,
            "approved_by": null,
            "approved_at": null,
            "approval_status": "pending",
            "rejection_reason": "",
            "created_at": "2025-11-13T23:14:48.076002+06:00",
            "updated_at": "2025-11-14T10:56:41.024118+06:00",
            "deleted_at": null,
            "meta_title": "",
            "meta_description": "",
            "view_count": 7,
            "sale_count": 0,
            "average_rating": 0,
            "review_count": 0,
            "seller": {
                "id": 0,
                "role_id": 0,
                "user_id": 0,
                "business_email": "",
                "phone": "",
                "store_name": "",
                "store_slug": "",
                "store_description": "",
                "store_logo": "",
                "store_banner": "",
                "website": "",
                "business_type": "",
                "tax_number": "",
                "business_license": "",
                "total_sales": 0,
                "total_reviews": 0,
                "total_earnings": 0,
                "total_orders": 0,
                "average_rating": 0,
                "commission": 0,
                "address": "",
                "city": "",
                "state": "",
                "country": "",
                "postal_code": "",
                "status": "",
                "approval_status": "",
                "is_active": false,
                "is_verified": false,
                "is_approved": false,
                "approved_at": null,
                "rejected_at": null,
                "rejection_reason": "",
                "is_profile_complete": false,
                "has_business_info": false,
                "has_address": false,
                "has_payment_method": false,
                "can_add_products": false,
                "missing_fields": null,
                "next_step": "",
                "approved_by": null,
                "created_at": "0001-01-01T00:00:00Z",
                "updated_at": "0001-01-01T00:00:00Z",
                "deleted_at": null,
                "role": {
                    "id": 0,
                    "name": "",
                    "description": "",
                    "is_active": false,
                    "created_at": "0001-01-01T00:00:00Z",
                    "updated_at": "0001-01-01T00:00:00Z",
                    "deleted_at": null,
                    "users": null,
                    "buyers": null
                },
                "products": null,
                "reg_user": null,
                "payment_methods": null,
                "seller_categories": null
            },
            "category": {
                "id": 0,
                "name": "",
                "slug": "",
                "description": "",
                "image": "",
                "icon": "",
                "parent_id": null,
                "sort_order": 0,
                "is_featured": false,
                "is_active": false,
                "meta_title": "",
                "meta_description": "",
                "meta_keywords": "",
                "created_at": "0001-01-01T00:00:00Z",
                "updated_at": "0001-01-01T00:00:00Z",
                "deleted_at": null,
                "children": null,
                "products": null
            },
            "images": [
                {
                    "id": 56,
                    "product_id": 41,
                    "image_url": "https://res.cloudinary.com/ddylnmsou/image/upload/v1763054090/products/pexels-ansel-lee-1635554-3192175_1763054088.webp",
                    "alt_text": "Product image 1",
                    "is_primary": true,
                    "sort_order": 1,
                    "type": "gallery",
                    "created_at": "2025-11-13T23:14:54.864605+06:00",
                    "updated_at": "2025-11-13T23:14:54.864605+06:00",
                    "product": {
                        "id": 0,
                        "seller_id": 0,
                        "name": "",
                        "slug": "",
                        "description": "",
                        "sku": "",
                        "category_id": 0,
                        "price": 0,
                        "discount price": 0,
                        "discount_percent": 0,
                        "height": "",
                        "age": "",
                        "tree_type": "",
                        "pot_size": "",
                        "scientific_name": "",
                        "common_names": "",
                        "quantity": 0,
                        "min_quantity": 0,
                        "weight": 0,
                        "is_active": false,
                        "is_approved": false,
                        "is_featured": false,
                        "approved_by": null,
                        "approved_at": null,
                        "approval_status": "",
                        "rejection_reason": "",
                        "created_at": "0001-01-01T00:00:00Z",
                        "updated_at": "0001-01-01T00:00:00Z",
                        "deleted_at": null,
                        "meta_title": "",
                        "meta_description": "",
                        "view_count": 0,
                        "sale_count": 0,
                        "average_rating": 0,
                        "review_count": 0,
                        "seller": {
                            "id": 0,
                            "role_id": 0,
                            "user_id": 0,
                            "business_email": "",
                            "phone": "",
                            "store_name": "",
                            "store_slug": "",
                            "store_description": "",
                            "store_logo": "",
                            "store_banner": "",
                            "website": "",
                            "business_type": "",
                            "tax_number": "",
                            "business_license": "",
                            "total_sales": 0,
                            "total_reviews": 0,
                            "total_earnings": 0,
                            "total_orders": 0,
                            "average_rating": 0,
                            "commission": 0,
                            "address": "",
                            "city": "",
                            "state": "",
                            "country": "",
                            "postal_code": "",
                            "status": "",
                            "approval_status": "",
                            "is_active": false,
                            "is_verified": false,
                            "is_approved": false,
                            "approved_at": null,
                            "rejected_at": null,
                            "rejection_reason": "",
                            "is_profile_complete": false,
                            "has_business_info": false,
                            "has_address": false,
                            "has_payment_method": false,
                            "can_add_products": false,
                            "missing_fields": null,
                            "next_step": "",
                            "approved_by": null,
                            "created_at": "0001-01-01T00:00:00Z",
                            "updated_at": "0001-01-01T00:00:00Z",
                            "deleted_at": null,
                            "role": {
                                "id": 0,
                                "name": "",
                                "description": "",
                                "is_active": false,
                                "created_at": "0001-01-01T00:00:00Z",
                                "updated_at": "0001-01-01T00:00:00Z",
                                "deleted_at": null,
                                "users": null,
                                "buyers": null
                            },
                            "products": null,
                            "reg_user": null,
                            "payment_methods": null,
                            "seller_categories": null
                        },
                        "category": {
                            "id": 0,
                            "name": "",
                            "slug": "",
                            "description": "",
                            "image": "",
                            "icon": "",
                            "parent_id": null,
                            "sort_order": 0,
                            "is_featured": false,
                            "is_active": false,
                            "meta_title": "",
                            "meta_description": "",
                            "meta_keywords": "",
                            "created_at": "0001-01-01T00:00:00Z",
                            "updated_at": "0001-01-01T00:00:00Z",
                            "deleted_at": null,
                            "children": null,
                            "products": null
                        },
                        "images": null,
                        "attributes": null,
                        "cart_items": null,
                        "order_items": null,
                        "reviews": null,
                        "wishlist_items": null
                    }
                }
            ],
            "attributes": null,
            "cart_items": null,
            "order_items": null,
            "reviews": null,
            "wishlist_items": null
        },
      }
    }
  ],
  "page": 1,
  "limit": 20,
  "total": 150
}
```

---

### 9. Search Products
Search products by name, description, or SKU.

**Endpoint:** `GET /api/products/search`

**Query Parameters:**
- `q` (string, required): Search query
- `page` (int, optional, default: 1)
- `limit` (int, optional, default: 20)

**Response:** `200 OK`
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Mango Tree",
      "price": 299.99,
      "image_url": "https://cloudinary.com/..."
    }
  ],
  "total": 5
}
```

---

### 10. Get Product by Slug
Get product details using slug (SEO-friendly).

**Endpoint:** `GET /api/product/slug/{slug}`

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "Mango Tree",
    "slug": "mango-tree-premium",
    "description": "Premium quality mango tree",
    "price": 299.99,
    "images": [...],
    "seller": {...}
  }
}
```

---

### 11. Get Public Seller Products
Get all products from a specific seller (public view).

**Endpoint:** `GET /api/seller/{sellerId}/products`

**Query Parameters:**
- `page` (int, optional, default: 1)
- `limit` (int, optional, default: 20)

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "seller": {
      "id": 10,
      "store_name": "Green Garden Nursery",
      "store_logo": "https://...",
      "average_rating": 4.7
    },
    "products": [
      {
        "id": 1,
        "name": "Mango Tree",
        "price": 299.99
      }
    ],
    "total": 25
  }
}
```

---

## Field Validations

### Product Creation
- `name`: Required, 3-255 characters
- `slug`: Required, unique, URL-friendly
- `sku`: Required, unique, 3-100 characters
- `category_id`: Required, must exist
- `price`: Required, > 0
- `quantity`: Required, >= 0
- `height`: Required
- `age`: Required
- `tree_type`: Optional, enum (fruit, ornamental, shade, medicinal)
- `weight`: Optional, for shipping calculation

### Image Upload
- Max 5 images per product
- Max file size: 5MB per image
- Allowed formats: JPG, JPEG, PNG
- Images automatically optimized and resized

---

## Product Status

### Approval Status
- `pending`: Awaiting admin approval
- `approved`: Approved and visible to buyers
- `rejected`: Rejected by admin

### Active Status
- `true`: Product is active and visible
- `false`: Product is inactive (hidden)

---

## Error Responses

**400 Bad Request:**
```json
{
  "success": false,
  "error": "Invalid request data"
}
```

**401 Unauthorized:**
```json
{
  "success": false,
  "error": "Authentication required"
}
```

**403 Forbidden:**
```json
{
  "success": false,
  "error": "You don't have permission to access this product"
}
```

**404 Not Found:**
```json
{
  "success": false,
  "error": "Product not found"
}
```

---

## cURL Examples

### Create Product
```bash
curl -X POST "http://localhost:3000/api/product/create" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Mango Tree",
    "slug": "mango-tree-premium",
    "sku": "MANGO-001",
    "category_id": 5,
    "price": 299.99,
    "quantity": 50,
    "height": "4-5 feet",
    "age": "2 years"
  }'
```

### Upload Images
```bash
curl -X POST "http://localhost:3000/api/product/1/images" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -F "images=@image1.jpg" \
  -F "images=@image2.jpg" \
  -F "is_primary=true"
```

### Get Products
```bash
curl -X GET "http://localhost:3000/api/products?page=1&limit=20&category=5"
```

### Search Products
```bash
curl -X GET "http://localhost:3000/api/products/search?q=mango&page=1"
```

---

## Notes

- All seller operations require authentication
- Products require admin approval before becoming visible to buyers
- Images are stored in Cloudinary with automatic optimization
- Slugs must be unique and URL-friendly
- SKUs must be unique across all products
- Soft delete allows product recovery
- Stock quantity is automatically updated on orders

---

## Support

For issues or questions, contact: support@treestore.com
