package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/saadahmedbd/Treestore/Config"
	adminhandler "github.com/saadahmedbd/Treestore/Rest/Handler/AdminHandler"
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
	selleraccountsettinghandler "github.com/saadahmedbd/Treestore/Rest/Handler/SellerAccountSettingHandler"
	sellerproflehandler "github.com/saadahmedbd/Treestore/Rest/Handler/SellerProfleHandler"
	userhandler "github.com/saadahmedbd/Treestore/Rest/Handler/UserHandler"
	carthandler "github.com/saadahmedbd/Treestore/Rest/Handler/cartHandler"
	selleraccounthandler "github.com/saadahmedbd/Treestore/Rest/Handler/sellerAccountHandler"

	middleware "github.com/saadahmedbd/Treestore/Rest/Middleware"
)

// crete dependency
type Server struct {
	buyerHandler                *buyerhandler.Handler
	cartitemhandler             *cartitemhandler.Handler
	orderitemhandler            *orderitemhandler.Handler
	producthandler              *producthandler.Handler
	profilehandler              *profilehandler.Handler
	reviewhandler               *reviewhandler.Handler
	rolehandler                 *rolehandler.Handler
	searchhandler               *searchhandler.Handler
	userhandler                 *userhandler.Handler
	authhandler                 *authhandler.Handler
	sellerprofilehandler        *sellerproflehandler.SellerProfileHandler
	buyerprofilehandler         *buyerprofilehandler.Buyerprofilehandler
	guestcarthandler            *guestcarthandler.GuestcartHandler
	selleraccounthandler        *selleraccounthandler.SellerRegistrationHandler
	carthandler                 *carthandler.CartHandler
	orderhandler                *orderhandler.OrderHandler
	adminhandler                *adminhandler.Adminhandler
	categoryHandler             *categoryHandler.CategoryHandler
	selleraccountsettinghandler *selleraccountsettinghandler.Selleraccountsettinghandler
}

func NewServer(
	buyerHandler *buyerhandler.Handler,
	cartitemHandler *cartitemhandler.Handler,
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
	sellerAccountHandler *selleraccounthandler.SellerRegistrationHandler,
	cartHandler *carthandler.CartHandler,
	orderHandler *orderhandler.OrderHandler,
	adminHadler *adminhandler.Adminhandler,
	categoryHandler *categoryHandler.CategoryHandler,
	selleraccountsettinghandler *selleraccountsettinghandler.Selleraccountsettinghandler,

) *Server {
	return &Server{
		buyerHandler:                buyerHandler,
		cartitemhandler:             cartitemHandler,
		orderitemhandler:            orderitemHandler,
		producthandler:              productHandler,
		profilehandler:              profileHandler,
		reviewhandler:               reviewHandler,
		rolehandler:                 rolehandler,
		searchhandler:               searchHandler,
		userhandler:                 userHandler,
		authhandler:                 authhandler,
		sellerprofilehandler:        sellerProfileHandler,
		buyerprofilehandler:         buyerProfileHandler,
		guestcarthandler:            guestcartHandler,
		selleraccounthandler:        sellerAccountHandler,
		carthandler:                 cartHandler,
		orderhandler:                orderHandler,
		adminhandler:                adminHadler,
		categoryHandler:             categoryHandler,
		selleraccountsettinghandler: selleraccountsettinghandler,
	}

}

func (server *Server) Start(cnf Config.Config) {
	mux := http.NewServeMux()
	// router

	server.buyerHandler.BuyerRoute(mux)
	server.cartitemhandler.CartItemRoute(mux)
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
	server.selleraccounthandler.SellerAccountRoute(mux)
	server.carthandler.CartItemRoute(mux)
	server.orderhandler.OrderRoute(mux)
	server.adminhandler.AdminRoute(mux)
	server.categoryHandler.CategoryRoute(mux)
	server.selleraccountsettinghandler.RegisterSellerAccountSettingHandler(mux)

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
