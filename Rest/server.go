package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	authhandler "github.com/saadahmedbd/Treestore/Rest/Handler/AuthHandler"
	buyerhandler "github.com/saadahmedbd/Treestore/Rest/Handler/BuyerHandler"
	buyerprofilehandler "github.com/saadahmedbd/Treestore/Rest/Handler/BuyerProfileHandler"
	cartitemhandler "github.com/saadahmedbd/Treestore/Rest/Handler/CartItemHandler"
	categoryHandler "github.com/saadahmedbd/Treestore/Rest/Handler/CategoryHandler"
	guestcarthandler "github.com/saadahmedbd/Treestore/Rest/Handler/GuestcartHandler"
	orderhandler "github.com/saadahmedbd/Treestore/Rest/Handler/OrderHandler"
	orderitemhandler "github.com/saadahmedbd/Treestore/Rest/Handler/OrderItemHandler"
	producthandler "github.com/saadahmedbd/Treestore/Rest/Handler/ProductHandler"
	profilehandler "github.com/saadahmedbd/Treestore/Rest/Handler/ProfileHandler"
	reviewhandler "github.com/saadahmedbd/Treestore/Rest/Handler/ReviewHandler"
	rolehandler "github.com/saadahmedbd/Treestore/Rest/Handler/RoleHandler"
	searchhandler "github.com/saadahmedbd/Treestore/Rest/Handler/SearchHandler"
	sellerproflehandler "github.com/saadahmedbd/Treestore/Rest/Handler/SellerProfleHandler"
	userhandler "github.com/saadahmedbd/Treestore/Rest/Handler/UserHandler"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

// crete dependency
type Server struct {
	buyerHandler         *buyerhandler.Handler
	cartitemhandler      *cartitemhandler.Handler
	categoryHandler      *categoryHandler.Handler
	orderhandler         *orderhandler.Handler
	orderitemhandler     *orderitemhandler.Handler
	producthandler       *producthandler.Handler
	profilehandler       *profilehandler.Handler
	reviewhandler        *reviewhandler.Handler
	rolehandler          *rolehandler.Handler
	searchhandler        *searchhandler.Handler
	userhandler          *userhandler.Handler
	authhandler          *authhandler.Handler
	sellerprofilehandler *sellerproflehandler.SellerProfileHandler
	buyerprofilehandler  *buyerprofilehandler.Buyerprofilehandler
	guestcarthandler     *guestcarthandler.GuestcartHandler
}

func NewServer(
	buyerHandler *buyerhandler.Handler,
	cartitemHandler *cartitemhandler.Handler,
	categoryHandler *categoryHandler.Handler,
	orderHandler *orderhandler.Handler,
	orderitemHandler *orderitemhandler.Handler,
	productHandler *producthandler.Handler,
	profileHandler *profilehandler.Handler,
	reviewHandler *reviewhandler.Handler,
	rolehandler *rolehandler.Handler,
	searchHandler *searchhandler.Handler,
	userHandler *userhandler.Handler,
	authhandler *authhandler.Handler,
	sellerProfileHandler *sellerproflehandler.SellerProfileHandler,
	buyerProfileHandler *buyerprofilehandler.Buyerprofilehandler,
	guestcartHandler *guestcarthandler.GuestcartHandler,

) *Server {
	return &Server{
		buyerHandler:         buyerHandler,
		cartitemhandler:      cartitemHandler,
		categoryHandler:      categoryHandler,
		orderhandler:         orderHandler,
		orderitemhandler:     orderitemHandler,
		producthandler:       productHandler,
		profilehandler:       profileHandler,
		reviewhandler:        reviewHandler,
		rolehandler:          rolehandler,
		searchhandler:        searchHandler,
		userhandler:          userHandler,
		authhandler:          authhandler,
		sellerprofilehandler: sellerProfileHandler,
		buyerprofilehandler:  buyerProfileHandler,
		guestcarthandler:     guestcartHandler,
	}

}

func (server *Server) Start(cnf Config.Config) {
	mux := http.NewServeMux()
	// router

	server.buyerHandler.BuyerRoute(mux)
	server.cartitemhandler.CartItemRoute(mux)
	server.categoryHandler.CategoryRoute(mux)
	server.orderhandler.OrderRoute(mux)
	server.orderitemhandler.OrderItemRoute(mux)
	server.producthandler.ProductRoute(mux)
	server.profilehandler.Profile(mux)
	server.reviewhandler.ReviewRoute(mux)
	server.rolehandler.RoleRoute(mux)
	server.searchhandler.Search_ProductRoute(mux)
	server.userhandler.UserRoute(mux)
	server.authhandler.AuthRouth(mux)
	server.sellerprofilehandler.SellerRoute(mux)
	server.buyerprofilehandler.BuyerRoute(mux)
	server.guestcarthandler.GuestCartRoute(mux)

	// Handle all cors
	handler := middleware.Cors(mux)

	address := ":" + strconv.Itoa(int(cnf.HttpPort))

	fmt.Println("server is running", address)

	err := http.ListenAndServe(address, handler)
	if err != nil {
		fmt.Println("Server already used...", err)
		os.Exit(1)
	}
}
