package CMD

import (
	"github.com/saadahmedbd/Treestore/Config"
	rest "github.com/saadahmedbd/Treestore/Rest"
	authhandler "github.com/saadahmedbd/Treestore/Rest/Handler/AuthHandler"
	buyerhandler "github.com/saadahmedbd/Treestore/Rest/Handler/BuyerHandler"
	buyerprofilehandler "github.com/saadahmedbd/Treestore/Rest/Handler/BuyerProfileHandler"
	cartitemhandler "github.com/saadahmedbd/Treestore/Rest/Handler/CartItemHandler"
	categoryHandler "github.com/saadahmedbd/Treestore/Rest/Handler/CategoryHandler"
	orderhandler "github.com/saadahmedbd/Treestore/Rest/Handler/OrderHandler"
	orderitemhandler "github.com/saadahmedbd/Treestore/Rest/Handler/OrderItemHandler"
	producthandler "github.com/saadahmedbd/Treestore/Rest/Handler/ProductHandler"
	profilehandler "github.com/saadahmedbd/Treestore/Rest/Handler/ProfileHandler"
	reviewhandler "github.com/saadahmedbd/Treestore/Rest/Handler/ReviewHandler"
	rolehandler "github.com/saadahmedbd/Treestore/Rest/Handler/RoleHandler"
	searchhandler "github.com/saadahmedbd/Treestore/Rest/Handler/SearchHandler"
	sellerproflehandler "github.com/saadahmedbd/Treestore/Rest/Handler/SellerProfleHandler"
	userhandler "github.com/saadahmedbd/Treestore/Rest/Handler/UserHandler"
	repository "github.com/saadahmedbd/Treestore/Rest/Repository"
	buyerProfilerepo "github.com/saadahmedbd/Treestore/Rest/Repository/BuyerProfileRepo"
	productservice "github.com/saadahmedbd/Treestore/Rest/Service/ProductService"
	buyerservice "github.com/saadahmedbd/Treestore/Rest/Service/buyerService"
	sellerprofileservice "github.com/saadahmedbd/Treestore/Rest/Service/sellerProfileService"
)

func Server() {
	cnf := Config.GetConfig()

	//repo
	repository := repository.NewSellerRepostory(Config.DB)
	buyerProfilerepo := buyerProfilerepo.NewBuyerRepository(Config.DB)

	//service
	productservice := productservice.NewProductService(Config.DB)
	sellerprofileservice := sellerprofileservice.NewSellerService(repository)
	buyerProfileservice := buyerservice.NewBuyerService(buyerProfilerepo)

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
	)
	server.Start(cnf)

}
