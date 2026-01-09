# API Usage Examples

Complete guide for testing all API endpoints with cURL examples.

## 🔐 Authentication

### 1. Login as Super Admin

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@treestore.com",
    "password": "YourSecurePassword123"
  }'
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "email": "admin@treestore.com",
    "role": "super_admin"
  }
}
```

**Note:** Copy the token and use it in subsequent requests.

---

## 👥 User Management

### 1. Get All Users

```bash
curl -X GET "http://localhost:8080/api/users?page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 2. Get Users by Status

```bash
# Get active users
curl -X GET "http://localhost:8080/api/users?page=1&limit=10&status=active" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"

# Get suspended users
curl -X GET "http://localhost:8080/api/users?page=1&limit=10&status=suspended" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 3. Search Users

```bash
curl -X GET "http://localhost:8080/api/users/search?q=john&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 4. Get User by ID

```bash
curl -X GET "http://localhost:8080/api/users/get?id=5" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 5. Get User Statistics

```bash
curl -X GET "http://localhost:8080/api/users/stats" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Response:**
```json
{
  "success": true,
  "data": {
    "total": 150,
    "active": 120,
    "inactive": 20,
    "suspended": 10
  }
}
```

### 6. Update User

```bash
curl -X PUT "http://localhost:8080/api/users/update?id=5" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "phone": "+1234567890",
    "gender": "male"
  }'
```

### 7. Activate User

```bash
curl -X POST "http://localhost:8080/api/users/activate?id=5" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 8. Deactivate User

```bash
curl -X POST "http://localhost:8080/api/users/deactivate?id=5" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 9. Suspend User

```bash
curl -X POST "http://localhost:8080/api/users/suspend?id=5" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

## 🏪 Seller Management

### 1. Get All Sellers

```bash
curl -X GET "http://localhost:8080/api/sellers?page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 2. Get Sellers by Status

```bash
# Get approved sellers
curl -X GET "http://localhost:8080/api/sellers?status=approved&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"

# Get suspended sellers
curl -X GET "http://localhost:8080/api/sellers?status=suspended&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 3. Get Pending Seller Approvals

```bash
curl -X GET "http://localhost:8080/api/sellers/pending?page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 4. Search Sellers

```bash
curl -X GET "http://localhost:8080/api/sellers/search?q=electronics&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 5. Get Seller by ID

```bash
curl -X GET "http://localhost:8080/api/sellers/get?id=3" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 6. Get Seller Statistics

```bash
curl -X GET "http://localhost:8080/api/sellers/stats" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Response:**
```json
{
  "success": true,
  "data": {
    "total": 50,
    "pending": 5,
    "approved": 40,
    "rejected": 3,
    "suspended": 2
  }
}
```

### 7. Get Top Sellers

```bash
curl -X GET "http://localhost:8080/api/sellers/top?limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 8. Approve Seller

```bash
curl -X POST "http://localhost:8080/api/sellers/approve?id=3" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 9. Reject Seller

```bash
curl -X POST "http://localhost:8080/api/sellers/reject?id=3" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "reason": "Incomplete business documentation"
  }'
```

### 10. Suspend Seller

```bash
curl -X POST "http://localhost:8080/api/sellers/suspend?id=3" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "reason": "Multiple customer complaints"
  }'
```

### 11. Reactivate Seller

```bash
curl -X POST "http://localhost:8080/api/sellers/reactivate?id=3" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

## 📦 Product Management

### 1. Get All Products

```bash
curl -X GET "http://localhost:8080/api/products?page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 2. Get Products by Status

```bash
# Get approved products
curl -X GET "http://localhost:8080/api/products?status=approved&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"

# Get pending products
curl -X GET "http://localhost:8080/api/products?status=pending&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 3. Get Pending Product Approvals

```bash
curl -X GET "http://localhost:8080/api/products/pending?page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 4. Search Products

```bash
curl -X GET "http://localhost:8080/api/products/search?q=laptop&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 5. Get Product by ID

```bash
curl -X GET "http://localhost:8080/api/products/get?id=15" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 6. Get Product Statistics

```bash
curl -X GET "http://localhost:8080/api/products/stats" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Response:**
```json
{
  "success": true,
  "data": {
    "total": 500,
    "pending": 25,
    "approved": 450,
    "rejected": 15,
    "low_stock": 30
  }
}
```

### 7. Approve Product

```bash
curl -X POST "http://localhost:8080/api/products/approve?id=15" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 8. Reject Product

```bash
curl -X POST "http://localhost:8080/api/products/reject?id=15" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "reason": "Product images do not match description"
  }'
```

### 9. Delete Product

```bash
curl -X DELETE "http://localhost:8080/api/products/delete?id=15" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

## 📋 Order Management

### 1. Get All Orders

```bash
curl -X GET "http://localhost:8080/api/orders?page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```
response 
   {
                                "id": 55,
                                "product_id": 40,
                                "image_url": "https://res.cloudinary.com/ddylnmsou/image/upload/v1763053011/products/pexels-minan1398-793012_%281%29_1763053011.webp",
                                "alt_text": "Product image 3",
                                "is_primary": false,
                                "sort_order": 3,
                                "type": "gallery",
                                "created_at": "2025-11-13T22:56:53.742714+06:00",
                                "updated_at": "2025-11-13T22:56:53.742714+06:00",
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
                    "seller": {
                        "id": 62,
                        "role_id": 2,
                        "user_id": 128,
                        "business_email": "hasan@gmail.com",
                        "phone": "01300000000",
                        "store_name": "hasan plant",
                        "store_slug": "hasan-plant",
                        "store_description": "",
                        "store_logo": "https://res.cloudinary.com/ddylnmsou/image/upload/v1762372086/store-logos/Black_Purple_and_White_Futuristic_and_Simple_Galaxy_Themed_Desktop_Wallpaper_1762372082.png",
                        "store_banner": "https://res.cloudinary.com/ddylnmsou/image/upload/v1762372283/store-banners/I_am_nothing......_1762372280.jpg",
                        "website": "N/A",
                        "business_type": "individual",
                        "tax_number": "",
                        "business_license": "",
                        "total_sales": 0,
                        "total_reviews": 0,
                        "total_earnings": 0,
                        "total_orders": 0,
                        "average_rating": 0,
                        "commission": 10,
                        "address": "janpur bankpara",
                        "city": "Sirajganj",
                        "state": "Rajshahi",
                        "country": "Bangladesh",
                        "postal_code": "5700",
                        "status": "approved",
                        "approval_status": "approved",
                        "is_active": true,
                        "is_verified": true,
                        "is_approved": true,
                        "approved_at": "2025-11-05T14:57:17.249727+06:00",
                        "rejected_at": null,
                        "rejection_reason": "",
                        "is_profile_complete": true,
                        "has_business_info": true,
                        "has_address": true,
                        "has_payment_method": true,
                        "can_add_products": true,
                        "missing_fields": null,
                        "next_step": "approved",
                        "approved_by": 5,
                        "created_at": "2025-11-05T14:53:02.386088+06:00",
                        "updated_at": "2025-11-06T01:51:23.419743+06:00",
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
                    }
                }
            ],
            "order_history": null
        }
    ],
    "page": 1,
    "limit": 8,
    "total": 22
}

### 2. Get Orders by Status

```bash
# Get pending orders
curl -X GET "http://localhost:8080/api/orders?status=pending&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
response: {
    "success": true,
    "data": [
        {
            "id": 9,
            "order_number": "ORD-1761068589-43",
            "buyer_id": 43,
            "seller_id": 0,
            "status": "delivered",
            "payment_status": "pending",
            "payment_method": "cash_on_delivery",
            "subtotal": 362.97,
            "shipping_cost": 10,
            "tax_amount": 0,
            "discount_amount": 2,
            "total": 370.97,
            "commission_rate": 0,
            "commission_amount": 0,
            "seller_earnings": 0,
            "shipping_address": "123 Main St, City",
            "billing_address": "123 Main St, City",
            "customer_email": "buyer@example.com",
            "customer_phone": "+1234567890",
            "cancellation_reason": "",
            "cancelled_by": "",
            "refund_amount": 0,
            "refund_reason": "",
            "tracking_number": "",
            "shipped_at": "2025-11-04T00:01:20.491543+06:00",
            "delivered_at": "2025-11-04T00:03:13.522344+06:00",
            "order_date": "0001-01-01T00:00:00Z",
            "confirmed_at": null,
            "notes": "Please deliver before 5 PM",
            "created_at": "2025-10-21T23:43:09.811323+06:00",
            "updated_at": "2025-11-04T00:03:13.52408+06:00",
            "deleted_at": null,
            "buyer": {
                "id": 43,
                "role_id": 3,
                "user_id": 76,
                "phone": "",
                "profile_picture": "",
                "profile_picture_url": "",
                "profile_picture_public_id": "",
                "status": "inactive",
                "default_address": "",
                "default_address_id": null,
                "is_active": true,
                "email_verified": false,
                "last_order_at": null,
                "total_orders_count": 0,
                "total_spent": 0,
                "created_at": "2025-10-21T23:35:51.532003+06:00",
                "updated_at": "2025-11-30T15:36:03.790023+06:00",
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
                "reg_user": {
                    "id": 76,
                    "first_name": "morzina",
                    "last_name": "khatun",
                    "email": "morzinakhatunn@gmail.com",
                    "password": "$2a$10$uQ8yYHHKiEKnrkf2WEUazueSjGdujUCXgP0zXvlQFfoDEsw8SMOF.",
                    "gender": "",
                    "date_of_birth": null,
                    "role": "buyer",
                    "phone": "N/A",
                    "avatar": "",
                    "is_active": true,
                    "is_verified": false,
                    "email_verified": false,
                    "profile_photo": "",
                    "two_factor_enabled": false,
                    "last_login_at": null,
                    "updated_at": "2025-10-21T23:35:51.529742+06:00",
                    "deleted_at": null,
                    "created_at": "2025-10-21T23:35:51.529742+06:00"
                },
                "orders": null,
                "cart_items": null,
                "reviews": null,
                "wishlists": null,
                "addresses": null
            },
            "reviews": null,
            "order_items": [
                {
                    "id": 12,
                    "order_id": 9,
                    "product_id": 15,
                    "seller_id": 31,
                    "product_name": "rose",
                    "product_sku": "rose-PREMIUM-001",
                    "quantity": 3,
                    "price": 120.99,
                    "total": 362.97,
                    "commission": 54.45,
                    "seller_earning": 308.52,
                    "status": "delivered",
                    "shipped_at": "2025-11-04T00:01:20.497715+06:00",
                    "delivered_at": "2025-11-04T00:03:13.531001+06:00",
                    "created_at": "2025-10-21T23:43:09.817873+06:00",
                    "order": {
                        "id": 0,
                        "order_number": "",
                        "buyer_id": 0,
                        "seller_id": 0,
                        "status": "",
                        "payment_status": "",
                        "payment_method": "",
                        "subtotal": 0,
                        "shipping_cost": 0,
                        "tax_amount": 0,
                        "discount_amount": 0,
                        "total": 0,
                        "commission_rate": 0,
                        "commission_amount": 0,
                        "seller_earnings": 0,
                        "shipping_address": "",
                        "billing_address": "",
                        "customer_email": "",
                        "customer_phone": "",
                        "cancellation_reason": "",
                        "cancelled_by": "",
                        "refund_amount": 0,
                        "refund_reason": "",
                        "tracking_number": "",
                        "shipped_at": null,
                        "delivered_at": null,
                        "order_date": "0001-01-01T00:00:00Z",
                        "confirmed_at": null,
                        "notes": "",
                        "created_at": "0001-01-01T00:00:00Z",
                        "updated_at": "0001-01-01T00:00:00Z",
                        "deleted_at": null,
                        "buyer": {
                            "id": 0,
                            "role_id": 0,
                            "user_id": 0,
                            "phone": "",
                            "profile_picture": "",
                            "profile_picture_url": "",
                            "profile_picture_public_id": "",
                            "status": "",
                            "default_address": "",
                            "default_address_id": null,
                            "is_active": false,
                            "email_verified": false,
                            "last_order_at": null,
                            "total_orders_count": 0,
                            "total_spent": 0,
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
                            "reg_user": null,
                            "orders": null,
                            "cart_items": null,
                            "reviews": null,
                            "wishlists": null,
                            "addresses": null
                        },
                        "reviews": null,
                        "order_items": null,
                        "order_history": null
                    },
                    "product": {
                        "id": 15,
                        "seller_id": 31,
                        "name": "rose",
                        "slug": "rose-3",
                        "description": "High-quality pine apple tree perfect for your garden",
                        "sku": "rose-PREMIUM-001",
                        "category_id": 1,
                        "price": 120.99,
                        "discount price": 12.99,
                        "discount_percent": 10.04,
                        "height": "5 feet",
                        "age": "3 years",
                        "tree_type": "fruit",
                        "pot_size": "extra large",
                        "scientific_name": "Mangifera indica",
                        "common_names": "Mango, Aam, Manga",
                        "quantity": 18,
                        "min_quantity": 1,
                        "weight": 8.5,
                        "is_active": true,
                        "is_approved": false,
                        "is_featured": false,
                        "approved_by": null,
                        "approved_at": null,
                        "approval_status": "pending",
                        "rejection_reason": "",
                        "created_at": "2025-10-14T23:04:51.83042+06:00",
                        "updated_at": "2025-10-21T23:43:09.822658+06:00",
                        "deleted_at": null,
                        "meta_title": "Premium Mango Tree - Best Quality",
                        "meta_description": "Buy premium mango trees online. High-quality, 3-year-old trees ready for your garden.",
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
                        "images": [
                            {
                                "id": 23,
                                "product_id": 15,
                                "image_url": "https://example.com/mango-tree-1.jpg",
                                "alt_text": "Mango tree full view",
                                "is_primary": true,
                                "sort_order": 1,
                                "type": "gallery",
                                "created_at": "2025-10-14T23:04:51.834837+06:00",
                                "updated_at": "2025-10-14T23:04:51.834837+06:00",
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
                            },
                            {
                                "id": 24,
                                "product_id": 15,
                                "image_url": "https://example.com/mango-tree-2.jpg",
                                "alt_text": "Mango tree leaves close-up",
                                "is_primary": false,
                                "sort_order": 2,
                                "type": "gallery",
                                "created_at": "2025-10-14T23:04:51.838044+06:00",
                                "updated_at": "2025-10-14T23:04:51.838044+06:00",
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
                    "seller": {
                        "id": 31,
                        "role_id": 2,
                        "user_id": 65,
                        "business_email": "hannansheikh@gmail.com",
                        "phone": "+018152097",
                        "store_name": "sobuj family store",
                        "store_slug": "sobuj-family-store",
                        "store_description": "",
                        "store_logo": "",
                        "store_banner": "",
                        "website": "N/A",
                        "business_type": "nursery",
                        "tax_number": "",
                        "business_license": "",
                        "total_sales": 0,
                        "total_reviews": 0,
                        "total_earnings": 0,
                        "total_orders": 0,
                        "average_rating": 0,
                        "commission": 10,
                        "address": "sirajganj",
                        "city": "sirajganj",
                        "state": "rajshahi Division",
                        "country": "Bangladesh",
                        "postal_code": "5700",
                        "status": "pending",
                        "approval_status": "pending",
                        "is_active": true,
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
                        "next_step": "wait_approval",
                        "approved_by": null,
                        "created_at": "2025-10-14T22:59:10.058136+06:00",
                        "updated_at": "2025-10-14T23:01:03.657898+06:00",
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
                    }
                }
            ],
            "order_history": null
        }
    ],
    "page": 1,
    "limit": 10,
    "total": 1
}
# Get delivered orders
curl -X GET "http://localhost:8080/api/orders?status=delivered&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"

# Get cancelled orders
curl -X GET "http://localhost:8080/api/orders?status=cancelled&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 3. Search Orders

```bash
# Search by order number
curl -X GET "http://localhost:8080/api/orders/search?q=ORD-12345&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"

# Search by buyer email
curl -X GET "http://localhost:8080/api/orders/search?q=buyer@email.com&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 4. Get Order by ID

```bash
curl -X GET "http://localhost:8080/api/orders/get?id=25" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 5. Get Order Statistics

```bash
curl -X GET "http://localhost:8080/api/orders/stats" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Response:**
```json
{
  "success": true,
  "data": {
    "total": 1500,
    "pending": 50,
    "confirmed": 100,
    "processing": 80,
    "shipped": 200,
    "delivered": 1000,
    "cancelled": 70,
    "total_revenue": 150000.50
  }
}
```
### update order stats now updated order stats give you montly 
```bash
curl -X GET "http://localhost:8080/api/orders/stats" \
  -H
  ```
  **response:**
```json
{
  "total": 120,
  "total_revenue": 54320.50,
  "statuses": {
    "pending": 15,
    "confirmed": 10,
    "processing": 20,
    "shipped": 30,
    "delivered": 40,
    "cancelled": 5
  },
  "monthly_report": [
    { "month": "2025-06", "revenue": 3200.00, "orders": 12 },
    { "month": "2025-07", "revenue": 5500.00, "orders": 20 },
    { "month": "2025-08", "revenue": 8800.00, "orders": 25 },
    { "month": "2025-09", "revenue": 6700.00, "orders": 22 },
    { "month": "2025-10", "revenue": 9500.00, "orders": 28 },
    { "month": "2025-11", "revenue": 7620.50, "orders": 13 }
  ]
}
```
### 6. Get Recent Orders

```bash
curl -X GET "http://localhost:8080/api/orders/recent?limit=20" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 7. Update Order Status

```bash
# Confirm order
curl -X PUT "http://localhost:8080/api/orders/update-status?id=25" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "confirmed"
  }'

# Mark as shipped
curl -X PUT "http://localhost:8080/api/orders/update-status?id=25" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "shipped"
  }'

# Mark as delivered
curl -X PUT "http://localhost:8080/api/orders/update-status?id=25" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "delivered"
  }'
```

### 8. Cancel Order

```bash
curl -X POST "http://localhost:8080/api/orders/cancel?id=25" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "reason": "Customer requested cancellation"
  }'
```

---

## 📊 Dashboard

### Get Complete Dashboard Statistics

```bash
curl -X GET "http://localhost:8080/api/dashboard/stats" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Response:**
```json
{
    "status": "success",
    "message": "Dashboard statistics retrieved successfully",
    "data": {
        "orders": {
            "monthly_report": [
                {
                    "month": "2025-10",
                    "orders": 9,
                    "revenue": 2066.05
                },
                {
                    "month": "2025-11",
                    "orders": 13,
                    "revenue": 130370.96
                }
            ],
            "statuses": {
                "cancelled": 7,
                "confirmed": 0,
                "delivered": 3,
                "pending": 11,
                "processing": 0,
                "shipped": 1
            },
            "total": 22,
            "total_revenue": 2164.55
        },
        "products": {
            "approved": 2,
            "low_stock": 0,
            "pending": 16,
            "rejected": 0,
            "total": 19
        },
        "sellers": {
            "approved": 6,
            "pending": 36,
            "rejected": 3,
            "suspended": 1,
            "total": 45
        },
        "users": {
            "active": 33,
            "inactive": 8,
            "suspended": 1,
            "total": 42
        }
    }
}
```bash
curl -X GET "localhost:8080/api/products/top"
---

    {
                    "id": 58,
                    "product_id": 42,
                    "image_url": "https://res.cloudinary.com/ddylnmsou/image/upload/v1763284627/products/2b8nnhtktwbPCcBrDiKH-1200-80_1763284626.webp",
                    "alt_text": "Product image 2",
                    "is_primary": false,
                    "sort_order": 2,
                    "type": "gallery",
                    "created_at": "2025-11-16T15:17:08.38171+06:00",
                    "updated_at": "2025-11-16T15:17:08.38171+06:00",
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
        }
    ]
}
```bash
curl -X GET "localhost:8080/api/sellers/top?limit=10"
```
**response:**
{
            "id": 62,
            "role_id": 2,
            "user_id": 128,
            "business_email": "hasan@gmail.com",
            "phone": "01300000000",
            "store_name": "hasan plant",
            "store_slug": "hasan-plant",
            "store_description": "",
            "store_logo": "https://res.cloudinary.com/ddylnmsou/image/upload/v1762372086/store-logos/Black_Purple_and_White_Futuristic_and_Simple_Galaxy_Themed_Desktop_Wallpaper_1762372082.png",
            "store_banner": "https://res.cloudinary.com/ddylnmsou/image/upload/v1762372283/store-banners/I_am_nothing......_1762372280.jpg",
            "website": "N/A",
            "business_type": "individual",
            "tax_number": "",
            "business_license": "",
            "total_sales": 0,
            "total_reviews": 0,
            "total_earnings": 0,
            "total_orders": 0,
            "average_rating": 0,
            "commission": 10,
            "address": "janpur bankpara",
            "city": "Sirajganj",
            "state": "Rajshahi",
            "country": "Bangladesh",
            "postal_code": "5700",
            "status": "approved",
            "approval_status": "approved",
            "is_active": true,
            "is_verified": true,
            "is_approved": true,
            "approved_at": "2025-11-05T14:57:17.249727+06:00",
            "rejected_at": null,
            "rejection_reason": "",
            "is_profile_complete": true,
            "has_business_info": true,
            "has_address": true,
            "has_payment_method": true,
            "can_add_products": true,
            "missing_fields": null,
            "next_step": "approved",
            "approved_by": 5,
            "created_at": "2025-11-05T14:53:02.386088+06:00",
            "updated_at": "2025-11-06T01:51:23.419743+06:00",
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
            "reg_user": {
                "id": 128,
                "first_name": "Hasan",
                "last_name": "Sheikh",
                "email": "hasan@gmail.com",
                "password": "$2a$10$vth3RBPMVe13BAypMI3GnOwcRZ9wen3ETQAEs4mGs8JpwHO26O922",
                "gender": "",
                "date_of_birth": null,
                "role": "seller",
                "phone": "N/A",
                "avatar": "",
                "is_active": true,
                "is_verified": true,
                "email_verified": false,
                "profile_photo": "https://res.cloudinary.com/ddylnmsou/image/upload/v1762362717/seller-profiles/Remove_background_project_1762362714.jpg",
                "two_factor_enabled": false,
                "last_login_at": null,
                "updated_at": "2025-11-10T22:23:22.491659+06:00",
                "deleted_at": null,
                "created_at": "0001-01-01T06:00:00+06:00"
            },
            "payment_methods": null,
            "seller_categories": null
        }
    ]
}
```bash
curl -X GET "localhost:api/orders/recent?limit=10"
```
**response:**
{
    "status": "success",
    "message": "Recent orders retrieved successfully",
    "data": [
        {
            "id": 22,
            "order_number": "ORD-1764136023-62",
            "buyer_id": 62,
            "seller_id": 0,
            "status": "delivered",
            "payment_status": "pending",
            "payment_method": "cash_on_delivery",
            "subtotal": 1800,
            "shipping_cost": 120,
            "tax_amount": 0,
            "discount_amount": 600,
            "total": 1320,
            "commission_rate": 0,
            "commission_amount": 0,
            "seller_earnings": 0,
            "shipping_address": "ss road, sirajgonj, rajshahi, 5700, Bangladesh",
            "billing_address": "ss road, sirajgonj, rajshahi, 5700, Bangladesh",
            "customer_email": "hijoltomal@gmail.com",
            "customer_phone": "01899999999",
            "cancellation_reason": "",
            "cancelled_by": "",
            "refund_amount": 0,
            "refund_reason": "",
            "tracking_number": "",
            "shipped_at": "2025-12-01T13:10:16.264739+06:00",
            "delivered_at": "2025-12-01T13:10:20.610351+06:00",
            "order_date": "0001-01-01T06:00:00+06:00",
            "confirmed_at": null,
            "notes": "",
            "created_at": "2025-11-26T11:47:03.970076+06:00",
            "updated_at": "2025-12-01T13:10:20.611065+06:00",
            "deleted_at": null,
            "buyer": {
                "id": 62,
                "role_id": 3,
                "user_id": 136,
                "phone": "",
                "profile_picture": "",
                "profile_picture_url": "",
                "profile_picture_public_id": "",
                "status": "active",
                "default_address": "",
                "default_address_id": null,
                "is_active": true,
                "email_verified": false,
                "last_order_at": null,
                "total_orders_count": 0,
                "total_spent": 0,
                "created_at": "2025-11-19T10:36:16.930152+06:00",
                "updated_at": "2025-11-19T10:36:16.930152+06:00",
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
                "reg_user": {
                    "id": 136,
                    "first_name": "hijol",
                    "last_name": "tomal",
                    "email": "hijoltomal@gmail.com",
                    "password": "$2a$10$SwzBbO3.x7mu0vVNF/gP8urZkDcE5wki8.6KJEducYi.Wc91zyvK6",
                    "gender": "",
                    "date_of_birth": null,
                    "role": "buyer",
                    "phone": "N/A",
                    "avatar": "",
                    "is_active": true,
                    "is_verified": false,
                    "email_verified": false,
                    "profile_photo": "",
                    "two_factor_enabled": false,
                    "last_login_at": null,
                    "updated_at": "2025-11-19T10:36:16.893386+06:00",
                    "deleted_at": null,
                    "created_at": "2025-11-19T10:36:16.893386+06:00"
                },
                "orders": null,
                "cart_items": null,
                "reviews": null,
                "wishlists": null,
                "addresses": null
            },
            "reviews": null,
            "order_items": [
                {
                    "id": 30,
                    "order_id": 22,
                    "product_id": 41,
                    "seller_id": 62,
                    "product_name": "sundori",
                    "product_sku": "SUNDORI-7721",
                    "quantity": 3,
                    "price": 600,
                    "total": 1800,
                    "commission": 270,
                    "seller_earning": 1530,
                    "status": "delivered",
                    "shipped_at": "2025-12-01T13:10:16.270138+06:00",
                    "delivered_at": "2025-12-01T13:10:20.614744+06:00",
                    "created_at": "2025-11-26T11:47:04.021576+06:00",
                    "order": {
                        "id": 0,
                        "order_number": "",
                        "buyer_id": 0,
                        "seller_id": 0,
                        "status": "",
                        "payment_status": "",
                        "payment_method": "",
                        "subtotal": 0,
                        "shipping_cost": 0,
                        "tax_amount": 0,
                        "discount_amount": 0,
                        "total": 0,
                        "commission_rate": 0,
                        "commission_amount": 0,
                        "seller_earnings": 0,
                        "shipping_address": "",
                        "billing_address": "",
                        "customer_email": "",
                        "customer_phone": "",
                        "cancellation_reason": "",
                        "cancelled_by": "",
                        "refund_amount": 0,
                        "refund_reason": "",
                        "tracking_number": "",
                        "shipped_at": null,
                        "delivered_at": null,
                        "order_date": "0001-01-01T00:00:00Z",
                        "confirmed_at": null,
                        "notes": "",
                        "created_at": "0001-01-01T00:00:00Z",
                        "updated_at": "0001-01-01T00:00:00Z",
                        "deleted_at": null,
                        "buyer": {
                            "id": 0,
                            "role_id": 0,
                            "user_id": 0,
                            "phone": "",
                            "profile_picture": "",
                            "profile_picture_url": "",
                            "profile_picture_public_id": "",
                            "status": "",
                            "default_address": "",
                            "default_address_id": null,
                            "is_active": false,
                            "email_verified": false,
                            "last_order_at": null,
                            "total_orders_count": 0,
                            "total_spent": 0,
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
                            "reg_user": null,
                            "orders": null,
                            "cart_items": null,
                            "reviews": null,
                            "wishlists": null,
                            "addresses": null
                        },
                        "reviews": null,
                        "order_items": null,
                        "order_history": null
                    },
                    "product": {
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
                        "quantity": 14,
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
                        "updated_at": "2025-11-26T11:47:04.024061+06:00",
                        "deleted_at": null,
                        "meta_title": "",
                        "meta_description": "",
                        "view_count": 13,
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
                    }
                }
            ],
            "order_history": null
        }
## 🏥 Health Check

```bash
curl -X GET "http://localhost:8080/api/health"
```

**Response:**
```
OK
```

---

## 📝 Response Format

### Success Response

```json
{
  "success": true,
  "message": "Operation successful",
  "data": { ... }
}
```

### Paginated Response

```json
{
  "success": true,
  "data": [ ... ],
  "page": 1,
  "limit": 10,
  "total": 100
}
```

### Error Response

```json
{
  "success": false,
  "error": "Error message here"
}
```

---

## 🔑 Common Status Codes

| Code | Meaning |
|------|---------|
| 200 | Success |
| 201 | Created |
| 400 | Bad Request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not Found |
| 500 | Internal Server Error |

---

## 💡 Tips

1. **Save your JWT token**: After login, save the token to use in all protected endpoints
2. **Pagination**: Use `page` and `limit` query parameters for large datasets
3. **Filtering**: Use `status` query parameter to filter results
4. **Search**: Use `q` query parameter for search functionality
5. **Error Handling**: Always check the response status code and error message

---

## 🧪 Testing Workflow

### Complete Testing Sequence

1. **Login**
   ```bash
   # Get JWT token
   curl -X POST http://localhost:8080/login ...
   ```

2. **Test User Management**
   ```bash
   # Get all users
   curl -X GET http://localhost:8080/api/users ...
   
   # Suspend a user
   curl -X POST http://localhost:8080/api/users/suspend?id=5 ...
   ```

3. **Test Seller Management**
   ```bash
   # Get pending sellers
   curl -X GET http://localhost:8080/api/sellers/pending ...
   
   # Approve a seller
   curl -X POST http://localhost:8080/api/sellers/approve?id=3 ...
   ```

4. **Test Product Management**
   ```bash
   # Get pending products
   curl -X GET http://localhost:8080/api/products/pending ...
   
   # Approve a product
   curl -X POST http://localhost:8080/api/products/approve?id=15 ...
   ```

5. **Test Order Management**
   ```bash
   # Get all orders
   curl -X GET http://localhost:8080/api/orders ...
   
   # Update order status
   curl -X PUT http://localhost:8080/api/orders/update-status?id=25 ...
   ```

6. **Check Dashboard**
   ```bash
   # Get complete statistics
   curl -X GET http://localhost:8080/api/dashboard/stats ...
   ```