package carthandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *CartHandler) CartItemRoute(mux *http.ServeMux) {
	mux.Handle("POST /api/cart", middleware.Chain(http.HandlerFunc(h.AddToCart),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))

	mux.Handle("GET /api/get/cart", middleware.Chain(http.HandlerFunc(h.GetCart),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("GET /api/cart/summary", middleware.Chain(http.HandlerFunc(h.GetCartSummary),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("PUT /api/cart/{productId}", middleware.Chain(http.HandlerFunc(h.UpdateCartItem),
		middleware.AuthenticateJWT,
		middleware.Cors,
		middleware.Logger,
	))
	mux.Handle("DELETE /api/cart/{productId}", middleware.Chain(http.HandlerFunc(h.RemoveFromCart),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("DELETE /api/cart", middleware.Chain(http.HandlerFunc(h.ClearCart),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("POST /api/cart/{productId}/increment", middleware.Chain(http.HandlerFunc(h.IncrementQuantity),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("POST /api/cart/{productId}/decrement", middleware.Chain(http.HandlerFunc(h.DecrementQuantity),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("POST /api/cart/bulk-update", middleware.Chain(http.HandlerFunc(h.BulkUpdateCart),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("POST /api/cart/{productId}/save-later", middleware.Chain(http.HandlerFunc(h.SaveForLater),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("GET /api/cart/saved-for-later", middleware.Chain(http.HandlerFunc(h.GetSavedForLater),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("POST /api/cart/{productId}/move-to-cart", middleware.Chain(http.HandlerFunc(h.MoveToCart),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("GET /api/cart/validate", middleware.Chain(http.HandlerFunc(h.ValidateCart),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("GET /api/cart/count", middleware.Chain(http.HandlerFunc(h.GetCartItemCount),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("GET /api/cart/total", middleware.Chain(http.HandlerFunc(h.GetCartTotal),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("POST /api/wishlist", middleware.Chain(http.HandlerFunc(h.AddToWishlist),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("GET /api/wishlist", middleware.Chain(http.HandlerFunc(h.GetWishlist),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("POST /api/cart/{productId}/move-to-wishlist", middleware.Chain(http.HandlerFunc(h.MoveToWishlist),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("POST /api/wishlist/{productId}/move-to-cart", middleware.Chain(http.HandlerFunc(h.MoveFromWishlistToCart),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("DELETE /api/wishlist/{productId}", middleware.Chain(http.HandlerFunc(h.RemoveFromWishlist),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
}
