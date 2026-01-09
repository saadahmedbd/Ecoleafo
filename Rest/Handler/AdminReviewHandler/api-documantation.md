
-- Create indexes for performance
CREATE INDEX IF NOT EXISTS idx_reviews_status ON reviews(status);
CREATE INDEX IF NOT EXISTS idx_reviews_is_reported ON reviews(is_reported);
CREATE INDEX IF NOT EXISTS idx_reviews_product_status ON reviews(product_id, status);

-- Create review_reports table
CREATE TABLE IF NOT EXISTS review_reports (
    id SERIAL PRIMARY KEY,
    review_id INTEGER NOT NULL REFERENCES reviews(id) ON DELETE CASCADE,
    reporter_id INTEGER NOT NULL,
    reporter_type VARCHAR(20) NOT NULL,
    reason VARCHAR(100) NOT NULL,
    details TEXT,
    status VARCHAR(20) DEFAULT 'pending',
    reviewed_by INTEGER REFERENCES admins(id),
    reviewed_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_review_reports_review ON review_reports(review_id);
CREATE INDEX IF NOT EXISTS idx_review_reports_status ON review_reports(status);
CREATE INDEX IF NOT EXISTS idx_review_reports_reporter ON review_reports(reporter_id);

-- Update existing reviews to 'approved' status (for backward compatibility)
UPDATE reviews SET status = 'approved' WHERE status IS NULL;
*/

// ============================================================================
// STEP 8: API DOCUMENTATION
// ============================================================================

/*
=============================================================================
REVIEW MODERATION API DOCUMENTATION
=============================================================================

BASE URL: http://localhost:8080
AUTHENTICATION: JWT Bearer token required
HEADERS: Authorization: Bearer {token}

=============================================================================
ADMIN ENDPOINTS
=============================================================================

1. GET ALL REVIEWS
   GET /api/admin/reviews?page=1&limit=20&status=pending&rating=1&is_reported=true&search=keyword

   Query Parameters:
   - page: Page number (default: 1)
   - limit: Items per page (default: 20)
   - status: pending, approved, rejected
   - rating: 1-5
   - is_reported: true/false
   - search: Search in title and comment

   Response:
   {
     "status": "success",
     "message": "Reviews retrieved successfully",
     "data": [
       {
         "id": 1,
         "product_id": 10,
         "product_name": "iPhone 15",
         "product_image": "https://...",
         "buyer_id": 5,
         "buyer_name": "John Doe",
         "buyer_avatar": "https://...",
         "seller_id": 3,
         "seller_name": "Tech Store",
         "rating": 1,
         "title": "Terrible product",
         "comment": "Received fake product",
         "status": "pending",
         "is_verified_purchase": true,
         "helpful_count": 0,
         "report_count": 3,
         "is_reported": true,
         "created_at": "2025-11-19T10:00:00Z",
         "images": ["https://..."]
       }
     ],
     "pagination": {
       "page": 1,
       "limit": 20,
       "total": 150,
       "total_pages": 8
     }
   }

2. GET PENDING REVIEWS
   GET /api/admin/reviews/pending?page=1&limit=20

   Returns: Reviews with status="pending"

3. GET REPORTED REVIEWS
   GET /api/admin/reviews/reported?page=1&limit=20

   Returns: Reviews with is_reported=true

4. APPROVE/REJECT REVIEW
   POST /api/admin/reviews/moderate?id=5

   Request Body:
   {
     "action": "approve",  // or "reject"
     "reason": "Spam content" // Required only for reject
   }

   Response:
   {
     "status": "success",
     "message": "Review approved successfully"
   }

5. DELETE REVIEW
   DELETE /api/admin/reviews/delete?id=5

   Response:
   {
     "status": "success",
     "message": "Review deleted successfully"
   }

6. GET REVIEW STATISTICS
   GET /api/admin/reviews/stats

   Response:
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

=============================================================================
REVIEW REPORTS
=============================================================================

7. REPORT A REVIEW (Buyer/Seller)
   POST /api/reviews/report?id=5

   Request Body:
   {
     "reason": "spam",  // spam, inappropriate, fake, offensive, other
     "details": "This is clearly a fake review"
   }

   Response:
   {
     "status": "success",
     "message": "Review reported successfully"
   }

8. GET ALL REPORTS (Admin)
   GET /api/admin/reviews/reports?status=pending&page=1&limit=20

   Query Parameters:
   - status: pending, reviewed, dismissed

   Response:
   {
     "status": "success",
     "data": [
       {
         "id": 10,
         "review_id": 5,
         "review": { /* full review object */ },
         "reporter_type": "buyer",
         "reason": "spam",
         "details": "Fake review",
         "status": "pending",
         "created_at": "2025-11-19T10:00:00Z"
       }
     ]
   }

9. REVIEW REPORT ACTION (Admin)
   POST /api/admin/reviews/reports/action?id=10&action=reviewed

   Query Parameters:
   - id: Report ID
   - action: reviewed or dismissed

   Response:
   {
     "status": "success",
     "message": "Report marked as reviewed"
   }

=============================================================================
WORKFLOW EXAMPLES
=============================================================================

1. BUYER WRITES REVIEW → Status: pending
2. ADMIN REVIEWS → Approves or Rejects
3. IF APPROVED → Review visible to public, updates product rating
4. IF REJECTED → Hidden from public, reason sent to buyer

REPORT WORKFLOW:
1. USER REPORTS REVIEW → Creates ReviewReport
2. REVIEW.is_reported = true, report_count++
3. ADMIN REVIEWS REPORT → Can approve/reject/delete the review
4. ADMIN MARKS REPORT → Status: reviewed or dismissed

=============================================================================
TESTING EXAMPLES (cURL)
=============================================================================

# Get pending reviews
curl -X GET "http://localhost:8080/api/admin/reviews/pending?page=1&limit=20" \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN"

# Approve review
curl -X POST "http://localhost:8080/api/admin/reviews/moderate?id=5" \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"action": "approve"}'

# Reject review
curl -X POST "http://localhost:8080/api/admin/reviews/moderate?id=5" \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"action": "reject", "reason": "Spam content"}'

# Report review (buyer/seller)
curl -X POST "http://localhost:8080/api/reviews/report?id=5" \
  -H "Authorization: Bearer YOUR_USER_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"reason": "spam", "details": "This is fake"}'

# Get review stats
curl -X GET "http://localhost:8080/api/admin/reviews/stats" \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN"
