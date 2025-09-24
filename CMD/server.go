package CMD

import (
	"github.com/saadahmedbd/Treestore/Config"
	rest "github.com/saadahmedbd/Treestore/Rest"
	buyerhandler "github.com/saadahmedbd/Treestore/Rest/Handler/BuyerHandler"
	cartitemhandler "github.com/saadahmedbd/Treestore/Rest/Handler/CartItemHandler"
	categoryHandler "github.com/saadahmedbd/Treestore/Rest/Handler/CategoryHandler"
	orderhandler "github.com/saadahmedbd/Treestore/Rest/Handler/OrderHandler"
	orderitemhandler "github.com/saadahmedbd/Treestore/Rest/Handler/OrderItemHandler"
)

func Server() {
	cnf := Config.GetConfig()
	buyerhandler := buyerhandler.NewHandler()
	cartitemhandler := cartitemhandler.NewHandler()
	categoryHandler := categoryHandler.NewHandler()
	orderhandler := orderhandler.NewHandler()
	orderitemhandler := orderitemhandler.NewHandler()
	server := rest.NewServer(buyerhandler,
		cartitemhandler,
		categoryHandler,
		orderhandler,
		orderitemhandler,
	)
	server.Start(cnf)

}
