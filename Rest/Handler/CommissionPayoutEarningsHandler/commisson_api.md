# COMMISSION & EARNINGS API DOCUMENTATION

## BASE URL
http://localhost:8080

## AUTHENTICATION
All endpoints require JWT Bearer token
Header: `Authorization: Bearer {token}`

---

## COMMISSION SETTINGS

### 1. Get Commission Settings
**GET** `/api/admin/commission/settings`

**Response:**
```json
{
  "status": "success",
  "message": "Commission settings retrieved successfully",
  "data": {
    "id": 1,
    "default_rate": 10.00,
    "description": "Default platform commission rate",
    "is_active": true,
    "updated_by": 1,
    "updated_at": "2025-11-19T10:30:00Z"
  }
}
```

### 2. Update Commission Settings (Super Admin Only)
**PUT** `/api/admin/commission/settings`

**Request Body:**
```json
{
  "default_rate": 12.00,
  "description": "Updated commission rate"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Commission settings updated successfully"
}
```

---

## SELLER EARNINGS

### 3. Get Seller Earnings
**GET** `/api/admin/sellers/earnings?seller_id=5`

**Response:**
```json
{
  "status": "success",
  "message": "Seller earnings retrieved successfully",
  "data": {
    "seller_id": 5,
    "seller_name": "John Doe",
    "store_name": "Tech Store",
    "total_orders": 120,
    "completed_orders": 115,
    "gross_sales": 25000.00,
    "total_commission": 2500.00,
    "net_earnings": 22500.00,
    "total_withdrawn": 15000.00,
    "available_balance": 7500.00,
    "pending_clearance": 0.00,
    "last_updated": "2025-11-19T10:30:00Z"
  }
}
```

### 4. Get All Seller Earnings
**GET** `/api/admin/sellers/earnings/all?page=1&limit=20`

**Response:**
```json
{
  "status": "success",
  "message": "Seller earnings retrieved successfully",
  "data": [
    {
      "seller_id": 5,
      "seller_name": "John Doe",
      "store_name": "Tech Store",
      "total_orders": 120,
      "gross_sales": 25000.00,
      "net_earnings": 22500.00,
      "available_balance": 7500.00
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 20,
    "total": 50,
    "total_pages": 3
  }
}
```

---

## PAYOUTS

### 5. Request Payout (Seller)
**POST** `/api/seller/payouts/request`

**Request Body:**
```json
{
  "amount": 5000.00,
  "payment_method": "bank_transfer",
  "bank_name": "Bank of America",
  "account_number": "1234567890",
  "account_name": "John Doe",
  "request_note": "Monthly withdrawal"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Payout request submitted successfully"
}
```

**Validation Rules:**
- Minimum payout: $50.00
- Available balance must be sufficient
- One pending request at a time

### 6. Get Pending Payouts (Admin)
**GET** `/api/admin/payouts/pending?page=1&limit=20`

**Response:**
```json
{
  "status": "success",
  "message": "Pending payouts retrieved successfully",
  "data": [
    {
      "id": 10,
      "seller_id": 5,
      "seller_name": "Tech Store",
      "amount": 5000.00,
      "payment_method": "bank_transfer",
      "status": "pending",
      "requested_at": "2025-11-18T10:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 20,
    "total": 5
  }
}
```

### 7. Process Payout (Admin)
**POST** `/api/admin/payouts/process?id=10`

**Request Body:**
```json
{
  "transaction_id": "TXN123456789",
  "admin_note": "Transferred via bank on 2025-11-19"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Payout processed successfully"
}
```

### 8. Reject Payout (Admin)
**POST** `/api/admin/payouts/reject?id=10`

**Request Body:**
```json
{
  "rejection_reason": "Incorrect bank account details"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Payout rejected successfully"
}
```

### 9. Get Payout History
**GET** `/api/admin/payouts/history?seller_id=5&page=1&limit=20`

**Query Parameters:**
- `seller_id` (optional): Filter by specific seller
- `page`: Page number (default: 1)
- `limit`: Items per page (default: 20)

**Response:**
```json
{
  "status": "success",
  "message": "Payout history retrieved successfully",
  "data": [
    {
      "id": 10,
      "seller_id": 5,
      "seller_name": "Tech Store",
      "amount": 5000.00,
      "payment_method": "bank_transfer",
      "status": "completed",
      "transaction_id": "TXN123456789",
      "requested_at": "2025-11-18T10:00:00Z",
      "processed_at": "2025-11-19T14:30:00Z",
      "completed_at": "2025-11-19T14:30:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 20,
    "total": 15
  }
}
```

---

## ANALYTICS

### 10. Get Platform Earnings Overview
**GET** `/api/admin/earnings/overview`

**Response:**
```json
{
  "status": "success",
  "message": "Platform earnings overview retrieved successfully",
  "data": {
    "total_gross_sales": 500000.00,
    "total_commission_earned": 50000.00,
    "total_seller_earnings": 450000.00,
    "pending_payouts": 25000.00,
    "completed_payouts": 425000.00,
    "total_orders": 1500,
    "completed_orders": 1450
  }
}
```

### 11. Get Monthly Revenue
**GET** `/api/admin/revenue/monthly?year=2025`

**Response:**
```json
{
  "status": "success",
  "message": "Monthly revenue retrieved successfully",
  "data": [
    {
      "month": "2025-01",
      "gross_revenue": 45000.00,
      "commission_earned": 4500.00,
      "total_orders": 150
    },
    {
      "month": "2025-02",
      "gross_revenue": 52000.00,
      "commission_earned": 5200.00,
      "total_orders": 175
    }
  ]
}
```

### 12. Get Top Sellers by Revenue
**GET** `/api/admin/sellers/top?limit=10`

**Response:**
```json
{
  "status": "success",
  "message": "Top sellers retrieved successfully",
  "data": [
    {
      "seller_id": 5,
      "seller_name": "Tech Store",
      "store_name": "Tech Store",
      "total_orders": 120,
      "gross_sales": 25000.00,
      "commission_generated": 2500.00
    },
    {
      "seller_id": 8,
      "seller_name": "Garden Paradise",
      "store_name": "Garden Paradise",
      "total_orders": 95,
      "gross_sales": 18500.00,
      "commission_generated": 1850.00
    }
  ]
}
```

---

## PAYOUT STATUS FLOW
```
pending → processing → completed
pending → rejected
```

---

## COMMISSION STATUS FLOW
```
pending → cleared → paid_out
```

**Status Descriptions:**
- **pending**: Order just completed, commission calculated but not yet available for withdrawal
- **cleared**: Commission is cleared and available for seller to withdraw (typically after refund period)
- **paid_out**: Commission has been paid out to seller

---

## ERROR RESPONSES

**400 Bad Request**
```json
{
  "status": "error",
  "message": "Insufficient balance for withdrawal"
}
```

**401 Unauthorized**
```json
{
  "status": "error",
  "message": "Unauthorized access"
}
```

**403 Forbidden**
```json
{
  "status": "error",
  "message": "Insufficient permissions"
}
```

**404 Not Found**
```json
{
  "status": "error",
  "message": "Payout not found"
}
```

**500 Internal Server Error**
```json
{
  "status": "error",
  "message": "Failed to process payout",
  "error": "database connection error"
}
```

---

## BUSINESS LOGIC

### Commission Calculation
When an order is marked as **delivered**:
1. Get commission rate (seller's custom rate or default platform rate)
2. Calculate: `commission_amount = total_amount × (rate / 100)`
3. Calculate: `seller_earnings = total_amount - commission_amount`
4. Create `OrderCommission` record with status "pending"
5. Update `SellerEarningsSummary`

### Commission Clearance
- **Option 1 (Immediate)**: Set status to "cleared" immediately
- **Option 2 (Delayed)**: Use cron job to clear after 7-14 days (refund protection period)

### Payout Processing
1. Seller requests payout
2. System validates available balance
3. Admin reviews and approves
4. Admin enters transaction ID
5. System deducts from available balance
6. System adds to total_withdrawn
7. Payout status changes to "completed"

---

## TESTING EXAMPLES (cURL)

### Get Commission Settings
```bash
curl -X GET http://localhost:8080/api/admin/commission/settings \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Get Seller Earnings
```bash
curl -X GET "http://localhost:8080/api/admin/sellers/earnings?seller_id=5" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Request Payout (Seller)
```bash
curl -X POST http://localhost:8080/api/seller/payout/request \
  -H "Authorization: Bearer YOUR_SELLER_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "amount": 5000.00,
    "payment_method": "bank_transfer",
    "bank_name": "Bank of America",
    "account_number": "1234567890",
    "account_name": "John Doe",
    "request_note": "Monthly withdrawal"
  }'
```

### Process Payout (Admin)
```bash
curl -X POST "http://localhost:8080/api/admin/payouts/process?id=10" \
  -H "Authorization: Bearer YOUR_ADMIN_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "transaction_id": "TXN123456789",
    "admin_note": "Transferred via bank"
  }'
```

### Get Platform Earnings Overview
```bash
curl -X GET http://localhost:8080/api/admin/earnings/overview \
  -H "Authorization: Bearer YOUR_ADMIN_JWT_TOKEN"
```

---

## NOTES

1. **Minimum Payout**: $50.00 (configurable)
2. **Commission Rate**: Can be set per seller or use default platform rate
3. **Clearance Period**: 7-14 days recommended for refund protection
4. **One Pending Request**: Sellers can only have one pending payout at a time
5. **Transaction Tracking**: All payouts must have transaction_id for audit trail