The Complete Journey:
1️⃣ Initial Login (Day 1)
bash# Client sends
POST /api/auth/login
{ "email": "user@example.com", "password": "pass123" }

# Server responds
{
  "access_token": "eyJhbGc...",     # Expires in 15 minutes
  "refresh_token": "Xy9kF2...",     # Expires in 7 days
  "expires_in": 900                 # seconds
}
2️⃣ Using the App (Day 1-7)
bash# Every API call
GET /api/users
Authorization: Bearer eyJhbGc... (access token)

# Works fine for 15 minutes ✅
3️⃣ Access Token Expires (After 15 min)
bash# API call fails
GET /api/users
Authorization: Bearer eyJhbGc... (expired)

# Server responds
401 Unauthorized
{ "error": "Token expired" }
4️⃣ Client Auto-Refreshes
bash# Client automatically calls (no user interaction needed)
POST /api/auth/refresh
{ "refresh_token": "Xy9kF2..." }

# Server process:
# 1. Check DB: SELECT * FROM refresh_tokens WHERE token='Xy9kF2...'
# 2. Validate: is_revoked=false AND expires_at > NOW()
# 3. Get user info
# 4. Revoke old token: UPDATE refresh_tokens SET is_revoked=true
# 5. Create NEW access token (15 min)
# 6. Create NEW refresh token (7 days)
# 7. Save new refresh token to DB

# Server responds
{
  "access_token": "eyJNEW...",      # Fresh 15 min token
  "refresh_token": "AbC789...",     # New 7 day token
  "expires_in": 900
}
```

**5️⃣ Continue Seamlessly**
- Client saves new tokens
- Retries the original failed request with new access token
- User never noticed anything!

### **Why Token Rotation?**

**Without Rotation (Bad):**
```
Login → Refresh Token: ABC123
Day 1: Use ABC123 ✅
Day 2: Use ABC123 ✅
Day 7: Use ABC123 ✅

⚠️ If stolen on Day 1, attacker has 7 days of access!
```

**With Rotation (Good):**
```
Login → Refresh Token: ABC123
Hour 1: Refresh with ABC123 → Get NEW token DEF456, ABC123 revoked ❌
Hour 2: Refresh with DEF456 → Get NEW token GHI789, DEF456 revoked ❌

✅ If ABC123 stolen, it only works once, then it's revoked!
Client-Side Implementation Example:
javascript// Axios interceptor example
axios.interceptors.response.use(
  response => response,
  async error => {
    if (error.response?.status === 401) {
      // Access token expired
      const refreshToken = localStorage.getItem('refresh_token');
      
      try {
        // Get new tokens
        const { data } = await axios.post('/api/auth/refresh', {
          refresh_token: refreshToken
        });
        
        // Save new tokens
        localStorage.setItem('access_token', data.access_token);
        localStorage.setItem('refresh_token', data.refresh_token);
        
        // Retry original request with new token
        error.config.headers.Authorization = `Bearer ${data.access_token}`;
        return axios.request(error.config);
      } catch (refreshError) {
        // Refresh failed - user needs to login again
        window.location.href = '/login';
      }
    }
    return Promise.reject(error);
  }
);
Security Benefits:

Short-lived access tokens - Even if stolen, only valid 15 minutes
Refresh tokens in DB - Can be revoked immediately
Token rotation - Each refresh invalidates the previous token
Logout all devices - Revoke all refresh tokens for a user

This is why refresh tokens are much more secure than long-lived JWTs! 🔐
How auth/me Works:
Flow:

Client sends request with Authorization: Bearer <access_token> header
Middleware/Handler extracts token from header
Verifies token using VerifyJwt()
Extracts user_id from claims
Fetches user from database
Returns user info (without password)

Example Request:
bashcurl -H "Authorization: Bearer eyJhbGc..." http://localhost:8080/api/auth/me
Response:
json{
  "id": 1,
  "email": "user@example.com",
  "first_name": "John",
  "last_name": "Doe",
  "role": "admin"
}
How logout-all Works:
Flow:

Client sends request with access token (to identify who is logging out)
Extract user_id from the access token
Mark ALL refresh tokens for that user as is_revoked = true in database
User is logged out from all devices/sessions

Why it works:

When user tries to refresh on any device, the system checks: is_revoked = false
All revoked tokens fail validation
User must login again on all devices

Example:
go// User logs in on 3 devices -> 3 refresh tokens in DB
// Device 1: refresh_token_abc (user_id: 1, is_revoked: false)
// Device 2: refresh_token_def (user_id: 1, is_revoked: false)  
// Device 3: refresh_token_ghi (user_id: 1, is_revoked: false)

// User calls /logout-all from Device 1
// System updates ALL tokens: is_revoked = true

// Now when Device 2 or Device 3 try to refresh:
// Token validation fails because is_revoked = true
// They must login again
```

**Visual Flow:**
```
Login (Device 1)     Login (Device 2)     Login (Device 3)
      ↓                    ↓                    ↓
  Token ABC            Token DEF            Token GHI
      ↓                    ↓                    ↓
         Device 1 calls /logout-all
                      ↓
         Revoke ALL tokens (ABC, DEF, GHI)
                      ↓
         All devices must re-authenticate
Key Point: The access tokens still work until they expire (15 min), but users can't get NEW access tokens because all refresh tokens are revoked. For immediate logout, you'd need a token blacklist or reduce access token TTL.