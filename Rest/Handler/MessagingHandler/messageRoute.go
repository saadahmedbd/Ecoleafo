package messaginghandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *MessagingHandler) RegisterMessageRoute(mux *http.ServeMux) {
	mux.Handle("GET /api/v1/messaging/conversations",
		middleware.Chain(http.HandlerFunc(h.GetConversations),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
		))
	mux.Handle("POST /api/v1/messaging/conversations",
		middleware.Chain(http.HandlerFunc(h.CreateConversation),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
		))
	mux.Handle("GET /api/v1/messaging/conversations/{id}",
		middleware.Chain(http.HandlerFunc(h.GetConversationDetails),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
		))
	mux.Handle("PUT /api/v1/messaging/conversations/{id}/read",
		middleware.Chain(http.HandlerFunc(h.MarkAsRead),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
		))
	mux.Handle("POST /api/v1/messaging/messages",
		middleware.Chain(http.HandlerFunc(h.SendMessage),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
		))
	mux.Handle("GET /api/v1/messaging/messages",
		middleware.Chain(http.HandlerFunc(h.GetMessages),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
		))

	mux.Handle("GET /api/v1/messaging/unread-count",
		middleware.Chain(http.HandlerFunc(h.GetUnreadCount),
			middleware.Logger,
			middleware.Cors,
			middleware.AuthenticateJWT,
		))
}
