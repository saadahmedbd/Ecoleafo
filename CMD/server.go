package CMD

import (
	"github.com/saadahmedbd/Treestore/Config"
	rest "github.com/saadahmedbd/Treestore/Rest"
	buyerhandler "github.com/saadahmedbd/Treestore/Rest/Handler/BuyerHandler"
	cartitemhandler "github.com/saadahmedbd/Treestore/Rest/Handler/CartItemHandler"
	categoryHandler "github.com/saadahmedbd/Treestore/Rest/Handler/CategoryHandler"
	orderhandler "github.com/saadahmedbd/Treestore/Rest/Handler/OrderHandler"
	orderitemhandler "github.com/saadahmedbd/Treestore/Rest/Handler/OrderItemHandler"
	producthandler "github.com/saadahmedbd/Treestore/Rest/Handler/ProductHandler"
	profilehandler "github.com/saadahmedbd/Treestore/Rest/Handler/ProfileHandler"
	reviewhandler "github.com/saadahmedbd/Treestore/Rest/Handler/ReviewHandler"
	rolehandler "github.com/saadahmedbd/Treestore/Rest/Handler/RoleHandler"
	searchhandler "github.com/saadahmedbd/Treestore/Rest/Handler/SearchHandler"
	userhandler "github.com/saadahmedbd/Treestore/Rest/Handler/UserHandler"
)

func Server() {
	cnf := Config.GetConfig()
	buyerhandler := buyerhandler.NewHandler()
	cartitemhandler := cartitemhandler.NewHandler()
	categoryHandler := categoryHandler.NewHandler()
	orderhandler := orderhandler.NewHandler()
	orderitemhandler := orderitemhandler.NewHandler()
	producthandler := producthandler.NewHandler()
	profilehandler := profilehandler.NewHandler()
	reviewhandler := reviewhandler.NewHandler()
	rolehandler := rolehandler.NewHandler()
	searchhandler := searchhandler.NewHandler()
	userhandler := userhandler.NewHandler()

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
	)
	server.Start(cnf)

}
