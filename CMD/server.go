package CMD

import (
	"github.com/saadahmedbd/Treestore/Config"
	rest "github.com/saadahmedbd/Treestore/Rest"
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
	selleraccounthandler "github.com/saadahmedbd/Treestore/Rest/Handler/sellerAccountHandler"
	repository "github.com/saadahmedbd/Treestore/Rest/Repository"
	buyercompleterepo "github.com/saadahmedbd/Treestore/Rest/Repository/BuyerCompleteRepo"
	buyerProfilerepo "github.com/saadahmedbd/Treestore/Rest/Repository/BuyerProfileRepo"
	guestcartrepo "github.com/saadahmedbd/Treestore/Rest/Repository/GuestCartRepo"
	selleraccountrepo "github.com/saadahmedbd/Treestore/Rest/Repository/SellerAccountRepo"
	productservice "github.com/saadahmedbd/Treestore/Rest/Service/ProductService"
	selleraccountservice "github.com/saadahmedbd/Treestore/Rest/Service/SellerAccountService"
	buyerservice "github.com/saadahmedbd/Treestore/Rest/Service/buyerService"
	guestcartservice "github.com/saadahmedbd/Treestore/Rest/Service/guestCartService"
	sellerprofileservice "github.com/saadahmedbd/Treestore/Rest/Service/sellerProfileService"
)

func Server() {
	cnf := Config.GetConfig()

	//repo
	repository := repository.NewSellerRepostory(Config.DB)
	buyerProfilerepo := buyerProfilerepo.NewBuyerRepository(Config.DB)
	guestcartrepo := guestcartrepo.NewGuestCartRepository(Config.DB)
	buyercompleterepo := buyercompleterepo.NewProfileCompleteRepository(Config.DB)
	selleraccountrepo := selleraccountrepo.NewSellerRegistrationRepository(Config.DB)

	//service
	productservice := productservice.NewProductService(Config.DB)
	sellerprofileservice := sellerprofileservice.NewSellerService(repository)
	buyerProfileservice := buyerservice.NewBuyerService(buyerProfilerepo)
	guestcartservice := guestcartservice.NewGuestCartService(guestcartrepo, buyercompleterepo, buyerProfilerepo)
	selleracoubtService := selleraccountservice.NewSellerRegistrationService(selleraccountrepo)

	//handler
	buyerhandler := buyerhandler.NewHandler()
	cartitemhandler := cartitemhandler.NewHandler()
	categoryHandler := categoryHandler.NewHandler()
	orderhandler := orderhandler.NewHandler()
	orderitemhandler := orderitemhandler.NewHandler()
	producthandler := producthandler.NewHandler(productservice)
	profilehandler := profilehandler.NewHandler()
	reviewhandler := reviewhandler.NewHandler()
	rolehandler := rolehandler.NewHandler()
	searchhandler := searchhandler.NewHandler()
	userhandler := userhandler.NewHandler()
	authhandler := authhandler.NewHandler(authhandler.NewProductService(Config.DB))
	sellerprofilehandler := sellerproflehandler.NewSellerProfileHandler(sellerprofileservice)
	buyerprofilehandler := buyerprofilehandler.NewBuyerProfileHandler(buyerProfileservice)
	guestcarthandler := guestcarthandler.NewGuestCartHandler(guestcartservice)
	selleraccounthandler := selleraccounthandler.NewSellerRegistrationHandler(selleracoubtService)

	server := rest.NewServer(buyerhandler,
		cartitemhandler,
		categoryHandler,
		orderhandler,
		orderitemhandler,
		producthandler,
		profilehandler,
		reviewhandler,
		rolehandler,
		searchhandler,
		userhandler,
		authhandler,
		sellerprofilehandler,
		buyerprofilehandler,
		guestcarthandler,
		selleraccounthandler,
	)
	server.Start(cnf)

}
