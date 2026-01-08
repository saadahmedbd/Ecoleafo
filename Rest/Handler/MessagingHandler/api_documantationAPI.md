# Messaging System API Documentation

## Base URL
```
http://localhost:8080/api/v1/messaging
```

## Authentication
All endpoints require JWT authentication via Bearer token in the Authorization header.

```http
Authorization: Bearer <your_jwt_token>
```

---

## Endpoints Overview

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/conversations` | Get all user's conversations |
| `POST` | `/conversations` | Create a new conversation |
| `GET` | `/conversations/{id}` | Get conversation details |
| `PUT` | `/conversations/{id}/read` | Mark conversation messages as read |
| `POST` | `/messages` | Send a message |
| `GET` | `/messages` | Get messages (with polling support) |
| `GET` | `/unread-count` | Get total unread message count |

---

## 1. Get All Conversations

Get a list of all conversations for the authenticated user.

### Request

```http
GET /api/v1/messaging/conversations
Authorization: Bearer <token>
```

### Response

**Status Code:** `200 OK`

```json
{
  "success": true,
  "message": "Conversations retrieved successfully",
  "data": {
    "conversations": [
      {
        "id": 1,
        "other_user": {
          "id": 5,
          "name": "Green Garden Store",
          "email": "seller@example.com",
          "user_type": "seller",
          "profile_picture": "https://example.com/logo.jpg",
          "store_name": "Green Garden Store"
        },
        "last_message": "Is this product still available?",
        "last_message_at": "2024-12-06T10:30:00Z",
        "unread_count": 2,
        "context": {
          "type": "product",
          "id": 123,
          "name": "Oak Tree 5ft",
          "image": "https://example.com/product.jpg"
        },
        "created_at": "2024-12-05T08:00:00Z"
      },
      {
        "id": 2,
        "other_user": {
          "id": 8,
          "name": "Tree Paradise",
          "email": "treeparadise@example.com",
          "user_type": "seller",
          "profile_picture": "",
          "store_name": "Tree Paradise"
        },
        "last_message": "Your order has been shipped",
        "last_message_at": "2024-12-05T14:20:00Z",
        "unread_count": 0,
        "context": {
          "type": "order",
          "id": 456,
          "name": "ORD-2024-001234",
          "image": ""
        },
        "created_at": "2024-12-04T12:00:00Z"
      }
    ],
    "total_unread": 2
  }
}
```

### Response Fields

| Field | Type | Description |
|-------|------|-------------|
| `id` | integer | Conversation ID |
| `other_user` | object | Information about the other participant |
| `other_user.id` | integer | User's role-specific ID (Buyer.ID, User.ID, Admin.ID) |
| `other_user.user_type` | string | Type: `buyer`, `seller`, or `admin` |
| `other_user.store_name` | string | Store name (for sellers only) |
| `last_message` | string | Preview of the last message |
| `last_message_at` | timestamp | When the last message was sent |
| `unread_count` | integer | Number of unread messages in this conversation |
| `context` | object | Context information (product/order) |
| `context.type` | string | Context type: `product`, `order`, or `support` |

---

## 2. Create Conversation

Start a new conversation with another user (typically a seller).

### Request

```http
POST /api/v1/messaging/conversations
Authorization: Bearer <token>
Content-Type: application/json
```

```json
{
  "recipient_id": 5,
  "recipient_type": "seller",
  "context_type": "product",
  "context_id": 123,
  "initial_message": "Is this product still available? I'm interested in purchasing."
}
```

### Request Body

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `recipient_id` | integer | Yes | The role-specific ID of the recipient (User.ID for seller, Buyer.ID for buyer) |
| `recipient_type` | string | Yes | Type of recipient: `buyer`, `seller`, or `admin` |
| `context_type` | string | No | Context type: `product`, `order`, or `support` |
| `context_id` | integer | No | ID of the product or order (required if context_type is set) |
| `initial_message` | string | Yes | First message (1-1000 characters) |

### Response

**Status Code:** `201 Created`

```json
{
  "success": true,
  "message": "Conversation created successfully",
  "data": {
    "id": 3,
    "other_user": {
      "id": 5,
      "name": "Green Garden Store",
      "email": "seller@example.com",
      "user_type": "seller",
      "profile_picture": "https://example.com/logo.jpg",
      "store_name": "Green Garden Store"
    },
    "last_message": "Is this product still available? I'm interested in purchasing.",
    "last_message_at": "2024-12-06T11:00:00Z",
    "unread_count": 0,
    "context": {
      "type": "product",
      "id": 123,
      "name": "Oak Tree 5ft",
      "image": "https://example.com/product.jpg"
    },
    "created_at": "2024-12-06T11:00:00Z"
  }
}
```

### Error Responses

**Status Code:** `400 Bad Request`
```json
{
  "success": false,
  "error": "Invalid request body"
}
```

**Status Code:** `404 Not Found`
```json
{
  "success": false,
  "error": "Recipient not found"
}
```

---

## 3. Get Conversation Details

Get detailed information about a specific conversation.

### Request

```http
GET /api/v1/messaging/conversations/{id}
Authorization: Bearer <token>
```

### Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `id` | integer | Conversation ID |

### Response

**Status Code:** `200 OK`

```json
{
  "success": true,
  "message": "Conversation details retrieved successfully",
  "data": {
    "id": 1,
    "other_user": {
      "id": 5,
      "name": "Green Garden Store",
      "email": "seller@example.com",
      "user_type": "seller",
      "profile_picture": "https://example.com/logo.jpg",
      "store_name": "Green Garden Store"
    },
    "last_message": "Yes, it's available. We can ship today.",
    "last_message_at": "2024-12-06T10:35:00Z",
    "unread_count": 1,
    "context": {
      "type": "product",
      "id": 123,
      "name": "Oak Tree 5ft",
      "image": "https://example.com/product.jpg"
    },
    "created_at": "2024-12-05T08:00:00Z"
  }
}
```

### Error Responses

**Status Code:** `404 Not Found`
```json
{
  "success": false,
  "error": "conversation not found"
}
```

**Status Code:** `403 Forbidden`
```json
{
  "success": false,
  "error": "unauthorized: not a participant"
}
```

---

## 4. Send Message

Send a message in an existing conversation.

### Request

```http
POST /api/v1/messaging/messages
Authorization: Bearer <token>
Content-Type: application/json
```

```json
{
  "conversation_id": 1,
  "message_text": "Great! Can you ship it to New York?"
}
```

### Request Body

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `conversation_id` | integer | Yes | ID of the conversation |
| `message_text` | string | Yes | Message content (1-5000 characters) |

### Response

**Status Code:** `201 Created`

```json
{
  "success": true,
  "message": "Message sent successfully",
  "data": {
    "id": 42,
    "sender_id": 3,
    "sender_name": "John Doe",
    "sender_type": "buyer",
    "message_text": "Great! Can you ship it to New York?",
    "is_read": false,
    "read_at": null,
    "created_at": "2024-12-06T11:05:00Z",
    "is_mine": true
  }
}
```

### Response Fields

| Field | Type | Description |
|-------|------|-------------|
| `id` | integer | Message ID |
| `sender_id` | integer | Role-specific ID of sender |
| `sender_name` | string | Display name of sender |
| `sender_type` | string | Type: `buyer`, `seller`, or `admin` |
| `message_text` | string | Message content |
| `is_read` | boolean | Whether message has been read |
| `read_at` | timestamp | When message was read (null if unread) |
| `created_at` | timestamp | When message was sent |
| `is_mine` | boolean | Whether current user is the sender |

### Error Responses

**Status Code:** `404 Not Found`
```json
{
  "success": false,
  "error": "conversation not found"
}
```

**Status Code:** `403 Forbidden`
```json
{
  "success": false,
  "error": "unauthorized: not a participant"
}
```

---

## 5. Get Messages (Polling)

Get messages from a conversation. Supports polling by using the `since` parameter.

### Request

```http
GET /api/v1/messaging/messages?conversation_id={id}&since={timestamp}&limit={limit}
Authorization: Bearer <token>
```

### Query Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `conversation_id` | integer | Yes | Conversation ID |
| `since` | string | No | RFC3339 timestamp - only return messages after this time |
| `limit` | integer | No | Maximum number of messages (1-100, default: 50) |

### Examples

**Get initial messages:**
```http
GET /api/v1/messaging/messages?conversation_id=1&limit=50
```

**Poll for new messages (every 5 seconds):**
```http
GET /api/v1/messaging/messages?conversation_id=1&since=2024-12-06T11:05:00Z
```

### Response

**Status Code:** `200 OK`

```json
{
  "success": true,
  "message": "Messages retrieved successfully",
  "data": {
    "messages": [
      {
        "id": 38,
        "sender_id": 3,
        "sender_name": "John Doe",
        "sender_type": "buyer",
        "message_text": "Is this product still available?",
        "is_read": true,
        "read_at": "2024-12-06T10:31:00Z",
        "created_at": "2024-12-06T10:30:00Z",
        "is_mine": true
      },
      {
        "id": 39,
        "sender_id": 5,
        "sender_name": "Green Garden Store",
        "sender_type": "seller",
        "message_text": "Yes, it's available. We can ship today.",
        "is_read": false,
        "read_at": null,
        "created_at": "2024-12-06T10:35:00Z",
        "is_mine": false
      },
      {
        "id": 42,
        "sender_id": 3,
        "sender_name": "John Doe",
        "sender_type": "buyer",
        "message_text": "Great! Can you ship it to New York?",
        "is_read": false,
        "read_at": null,
        "created_at": "2024-12-06T11:05:00Z",
        "is_mine": true
      }
    ],
    "has_more": false,
    "conversation": {
      "id": 1,
      "other_user": {
        "id": 5,
        "name": "Green Garden Store",
        "user_type": "seller"
      }
    }
  }
}
```

### Polling Strategy

For real-time updates, implement polling:

```javascript
// Initial load
GET /messages?conversation_id=1&limit=50

// Save timestamp of last message
const lastMessageTime = "2024-12-06T11:05:00Z";

// Poll every 5 seconds for new messages
setInterval(() => {
  GET /messages?conversation_id=1&since=2024-12-06T11:05:00Z
}, 5000);
```

---

## 6. Mark Messages as Read

Mark all unread messages in a conversation as read.

### Request

```http
PUT /api/v1/messaging/conversations/{id}/read
Authorization: Bearer <token>
```

### Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `id` | integer | Conversation ID |

### Response

**Status Code:** `200 OK`

```json
{
  "success": true,
  "message": "Messages marked as read",
  "data": null
}
```

### Notes

- This marks ALL unread messages in the conversation as read
- Only marks messages where the current user is NOT the sender
- Updates the `is_read` field and sets `read_at` timestamp
- Resets the unread count for the current user

---

## 7. Get Unread Count

Get the total number of unread messages across all conversations.

### Request

```http
GET /api/v1/messaging/unread-count
Authorization: Bearer <token>
```

### Response

**Status Code:** `200 OK`

```json
{
  "success": true,
  "message": "Unread count retrieved successfully",
  "data": {
    "total_unread": 5
  }
}
```

### Polling for Badge Updates

Poll this endpoint every 10 seconds to update notification badges:

```javascript
setInterval(() => {
  GET /api/v1/messaging/unread-count
  // Update badge UI with total_unread value
}, 10000);
```

---

## Error Responses

### Common Error Status Codes

| Status Code | Description |
|-------------|-------------|
| `400` | Bad Request - Invalid input data |
| `401` | Unauthorized - Missing or invalid JWT token |
| `403` | Forbidden - User not authorized for this action |
| `404` | Not Found - Resource doesn't exist |
| `500` | Internal Server Error - Server-side error |

### Error Response Format

```json
{
  "success": false,
  "error": "Error message describing what went wrong"
}
```

### Example Error Responses

**Invalid Token:**
```json
{
  "error": "Invalid or expired token: signature is invalid"
}
```

**Missing Fields:**
```json
{
  "success": false,
  "error": "Key: 'CreateConversationRequest.recipient_id' Error:Field validation for 'recipient_id' failed on the 'required' tag"
}
```

**Not a Participant:**
```json
{
  "success": false,
  "error": "unauthorized: not a participant"
}
```

---

## Data Models

### Conversation Object

```json
{
  "id": 1,
  "other_user": {
    "id": 5,
    "name": "Green Garden Store",
    "email": "seller@example.com",
    "user_type": "seller",
    "profile_picture": "https://example.com/logo.jpg",
    "store_name": "Green Garden Store"
  },
  "last_message": "Message text",
  "last_message_at": "2024-12-06T10:30:00Z",
  "unread_count": 2,
  "context": {
    "type": "product",
    "id": 123,
    "name": "Oak Tree 5ft",
    "image": "https://example.com/product.jpg"
  },
  "created_at": "2024-12-05T08:00:00Z"
}
```

### Message Object

```json
{
  "id": 42,
  "sender_id": 3,
  "sender_name": "John Doe",
  "sender_type": "buyer",
  "message_text": "Message content",
  "is_read": false,
  "read_at": null,
  "created_at": "2024-12-06T11:05:00Z",
  "is_mine": true
}
```

---

## Usage Examples

### Example 1: Buyer Messages Seller About Product

```javascript
// 1. Buyer clicks "Message Seller" on product page
POST /api/v1/messaging/conversations
{
  "recipient_id": 5,           // Seller's User.ID
  "recipient_type": "seller",
  "context_type": "product",
  "context_id": 123,           // Product.ID
  "initial_message": "Is this product available in larger sizes?"
}

// 2. Navigate to messages page and poll for response
GET /api/v1/messaging/messages?conversation_id=1

// 3. Poll every 5 seconds for new messages
setInterval(() => {
  GET /api/v1/messaging/messages?conversation_id=1&since=2024-12-06T11:05:00Z
}, 5000);

// 4. Buyer sends another message
POST /api/v1/messaging/messages
{
  "conversation_id": 1,
  "message_text": "I need it by next week."
}
```

### Example 2: Seller Contacts Buyer About Order

```javascript
// 1. Seller creates conversation from order page
POST /api/v1/messaging/conversations
{
  "recipient_id": 3,           // Buyer's Buyer.ID
  "recipient_type": "buyer",
  "context_type": "order",
  "context_id": 456,           // Order.ID
  "initial_message": "Your order has been prepared and will ship tomorrow."
}

// 2. Get all seller's conversations
GET /api/v1/messaging/conversations

// 3. Update unread badge every 10 seconds
setInterval(() => {
  GET /api/v1/messaging/unread-count
}, 10000);
```

### Example 3: Buyer Contacts Support

```javascript
// 1. Start support conversation
POST /api/v1/messaging/conversations
{
  "recipient_id": 1,           // Admin.ID
  "recipient_type": "admin",
  "context_type": "support",
  "initial_message": "I need help with my account settings."
}
```

---

## Rate Limiting

**Recommended polling intervals:**
- Unread count: 10 seconds
- Active conversation messages: 5 seconds
- Conversation list: 10 seconds

**Best practices:**
- Only poll when user is active (use Page Visibility API)
- Stop polling when app is in background
- Use exponential backoff on errors

---

## WebSocket Alternative (Future)

For true real-time messaging, consider implementing WebSockets:

```javascript
// Future enhancement
const ws = new WebSocket('ws://localhost:8080/api/v1/messaging/ws');

ws.onmessage = (event) => {
  const message = JSON.parse(event.data);
  // Update UI immediately
};
```

---

## Testing with cURL

### Get Conversations
```bash
curl -X GET http://localhost:8080/api/v1/messaging/conversations \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

### Create Conversation
```bash
curl -X POST http://localhost:8080/api/v1/messaging/conversations \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "recipient_id": 5,
    "recipient_type": "seller",
    "context_type": "product",
    "context_id": 123,
    "initial_message": "Is this available?"
  }'
```

### Send Message
```bash
curl -X POST http://localhost:8080/api/v1/messaging/messages \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "conversation_id": 1,
    "message_text": "Thank you!"
  }'
```

### Get Messages
```bash
curl -X GET "http://localhost:8080/api/v1/messaging/messages?conversation_id=1&limit=50" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### Mark as Read
```bash
curl -X PUT http://localhost:8080/api/v1/messaging/conversations/1/read \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### Get Unread Count
```bash
curl -X GET http://localhost:8080/api/v1/messaging/unread-count \
  -H "Authorization: Bearer YOUR_TOKEN"
```

---

## Database Schema Reference

### conversations table
```sql
CREATE TABLE conversations (
    id BIGINT PRIMARY KEY,
    participant1_id BIGINT NOT NULL,
    participant1_type VARCHAR(20) NOT NULL,
    participant2_id BIGINT NOT NULL,
    participant2_type VARCHAR(20) NOT NULL,
    context_type VARCHAR(20),
    context_id BIGINT,
    last_message TEXT,
    last_message_at TIMESTAMP,
    participant1_unread_count INT DEFAULT 0,
    participant2_unread_count INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### messages table
```sql
CREATE TABLE messages (
    id BIGINT PRIMARY KEY,
    conversation_id BIGINT NOT NULL,
    sender_id BIGINT NOT NULL,
    sender_type VARCHAR(20) NOT NULL,
    message_text TEXT NOT NULL,
    is_read BOOLEAN DEFAULT FALSE,
    read_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

## Support

For issues or questions:
- Check error messages in response
- Verify JWT token is valid
- Ensure user has proper permissions
- Check database for conversation/message existence

---

**API Version:** 1.0  
**Last Updated:** December 6, 2024  
**Base URL:** `http://localhost:8080/api/v1/messaging`