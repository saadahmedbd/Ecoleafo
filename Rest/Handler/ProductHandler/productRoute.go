package producthandler

import (
	"net/http"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

func (h *Handler) ProductRoute(mux *http.ServeMux) {
	mux.Handle("GET /api/products", middleware.Chain(http.HandlerFunc(h.GetProduct),
		middleware.Logger,
	))
	mux.Handle("GET /api/products/{productId}", middleware.Chain(http.HandlerFunc(h.GetProductById),
		middleware.Logger,
	))
	mux.Handle("GET /api/seller/{sellerID}/products", middleware.Chain(http.HandlerFunc(h.GetPublicSellerProducts),
		middleware.Logger,
	))
	//protected
	mux.Handle("POST /api/addproducts", middleware.Chain(http.HandlerFunc(h.CreateProduct),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))

	mux.Handle("PUT /api/updateproducts/{productId}", middleware.Chain(http.HandlerFunc(h.UpdateProduct),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("DELETE /api/deleteproducts/{productId}", middleware.Chain(http.HandlerFunc(h.DeleteProduct),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	mux.Handle("GET /api/seller/products", middleware.Chain(http.HandlerFunc(h.GetSellerProduct),
		middleware.AuthenticateJWT,
		middleware.Logger,
	))
	// mux.Handle("POST /api/products/{productId}/images", middleware.Chain(http.HandlerFunc(h.AddProductImage),
	// 	middleware.AuthenticateJWT,
	// 	middleware.Logger,
	// ))
	//upload images
	mux.Handle("POST /api/products/{ImageId}/images", middleware.Chain(http.HandlerFunc(h.UploadImage),
		middleware.AuthenticateJWT,
		middleware.Cors,
		middleware.Logger,
	))
	mux.Handle("POST /api/products/{ImageId}/images/multiple", middleware.Chain(http.HandlerFunc(h.UploadMultipleProductImages),
		middleware.AuthenticateJWT,
		middleware.Cors,
		middleware.Logger,
	))
	mux.Handle("POST /api/products/{ImageId}/images/url", middleware.Chain(http.HandlerFunc(h.AddProductImageByURL),
		middleware.AuthenticateJWT,
		middleware.Cors,
		middleware.Logger,
	))
	mux.Handle("POST /api/product/{productId}/images/{imageId}", middleware.Chain(http.HandlerFunc(h.DeleteProductImage),

		middleware.AuthenticateJWT,
		middleware.Cors,
		middleware.Logger,
	))

}
