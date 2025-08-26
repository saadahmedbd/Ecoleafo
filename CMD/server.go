package CMD

import (
	"fmt"
	"net/http"

	buyerhandler "github.com/saadahmedbd/Treestore/Handler/BuyerHandler"
	cartitemhandler "github.com/saadahmedbd/Treestore/Handler/CartItemHandler"
	categoryHandler "github.com/saadahmedbd/Treestore/Handler/CategoryHandler"
	orderhandler "github.com/saadahmedbd/Treestore/Handler/OrderHandler"
	orderitemhandler "github.com/saadahmedbd/Treestore/Handler/OrderItemHandler"
	producthandler "github.com/saadahmedbd/Treestore/Handler/ProductHandler"
	reviewhandler "github.com/saadahmedbd/Treestore/Handler/ReviewHandler"
	rolehandler "github.com/saadahmedbd/Treestore/Handler/RoleHandler"
	userhandler "github.com/saadahmedbd/Treestore/Handler/UserHandler"

	routehandler "github.com/saadahmedbd/Treestore/RouteHandler"
)

func Server() {
	mux := http.NewServeMux()

	// router
	mux.HandleFunc("/home", userhandler.Home)
	// user(seller) route
	mux.Handle("GET /getseller", http.HandlerFunc(userhandler.GetSeller))                     //Getseller route
	mux.Handle("POST /createseller", http.HandlerFunc(userhandler.CreateSeller))              // create seller route
	mux.Handle("GET /getseller/{sellerId}", http.HandlerFunc(userhandler.GetSellerById))      // get seller by id route
	mux.Handle("PUT /updateseller/{sellerId}", http.HandlerFunc(userhandler.UpdateSeller))    //update seller route
	mux.Handle("DELETE /deleteseller/{sellerId}", http.HandlerFunc(userhandler.DeleteSeller)) //update seller
	// buyer route

	mux.Handle("GET /getbuyer", http.HandlerFunc(buyerhandler.GetBuyer))
	mux.Handle("POST /createbuyer", http.HandlerFunc(buyerhandler.CreateBuyer))
	mux.Handle("GET /getbuyer/{buyerId}", http.HandlerFunc(buyerhandler.GetBuyerById))
	mux.Handle("PUT /updatebuyer/{buyerId}", http.HandlerFunc(buyerhandler.UpdateBuyer))
	mux.Handle("DELETE /deletebuyer/{buyerId}", http.HandlerFunc(buyerhandler.DeleteBuyer))

	//product route
	mux.Handle("GET /getproduct", http.HandlerFunc(producthandler.GetProduct))
	mux.Handle("POST /createproduct", http.HandlerFunc(producthandler.CreateProduct))
	mux.Handle("GET /getproduct/{productId}", http.HandlerFunc(producthandler.GetProductById))
	mux.Handle("PUT /updateproduct/{productId}", http.HandlerFunc(producthandler.UpdateProduct))
	mux.Handle("DELETE /deleteproduct/{productId}", http.HandlerFunc(producthandler.DeleteProduct))

	//category route
	mux.Handle("GET /getcategory", http.HandlerFunc(categoryHandler.GetCategory))
	mux.Handle("POST /createcategory", http.HandlerFunc(categoryHandler.CreateCategory))
	mux.Handle("GET /getcategory/{categoryId}", http.HandlerFunc(categoryHandler.GetCategoryById))
	mux.Handle("PUT /updatecategory/{categoryId}", http.HandlerFunc(categoryHandler.UpdateCategory))
	mux.Handle("DELETE /deletecategory/{categoryId}", http.HandlerFunc(categoryHandler.DeleteCategory))

	// handle cartitem route
	mux.Handle("GET /getcartitem", http.HandlerFunc(cartitemhandler.GetCartItem))
	mux.Handle("GET /getcartitem/{cartitemId}", http.HandlerFunc(cartitemhandler.GetcartItemById))
	mux.Handle("POST /createcartitem", http.HandlerFunc(cartitemhandler.CreatecartItem))
	mux.Handle("PUT /updatecartitem/{cartitemId}", http.HandlerFunc(cartitemhandler.UpdateCartItem))
	mux.Handle("DELETE /deletecartitem/{cartitemId}", http.HandlerFunc(cartitemhandler.DeleteCartItem))

	// handle order route
	mux.Handle("GET /getorder", http.HandlerFunc(orderhandler.GetOrder))
	mux.Handle("GET /getorder/{orderId}", http.HandlerFunc(orderhandler.GetOrderById))
	mux.Handle("POST /createorder", http.HandlerFunc(orderhandler.CreateOrder))
	mux.Handle("PUT /updateorder/{orderId}", http.HandlerFunc(orderhandler.UpdateOrder))
	mux.Handle("DELETE /deleteorder/{orderId}", http.HandlerFunc(orderhandler.DeleteOrder))

	//handle order item route
	mux.Handle("GET /getorderitem", http.HandlerFunc(orderitemhandler.GetOrderItem))
	mux.Handle("GET /getorderitem/{orderitemId}", http.HandlerFunc(orderitemhandler.GetOrderById))
	mux.Handle("POST /createorderitem", http.HandlerFunc(orderitemhandler.CreateOrderItem))
	mux.Handle("POST /updateorderitem/{orderitemId}", http.HandlerFunc(orderitemhandler.UpdateOrderItem))
	mux.Handle("DELETE /deleteorderitem/{orderitemId}", http.HandlerFunc(orderitemhandler.DeleteOrderItem))

	// handle review route
	mux.Handle("GET /getreview", http.HandlerFunc(reviewhandler.GetReview))
	mux.Handle("GET /getreview/{reviewId}", http.HandlerFunc(reviewhandler.GetReviewById))
	mux.Handle("POST /createreview", http.HandlerFunc(reviewhandler.CreateReview))
	mux.Handle("PUT /updatereview/{reviewId}", http.HandlerFunc(reviewhandler.UpdateReview))
	mux.Handle("GET /deletereview/{reviewId}", http.HandlerFunc(reviewhandler.DeleteReview))

	//role route
	mux.Handle("GET /getrole", http.HandlerFunc(rolehandler.GetRole))

	// Handle all cors
	globalRoute := routehandler.GlobalHandler(mux)

	fmt.Println("3000 port server is running")
	err := http.ListenAndServe(":3000", globalRoute)
	if err != nil {
		fmt.Println("Server already used...", err)
	}
}
