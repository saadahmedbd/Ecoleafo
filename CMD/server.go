package CMD

import (
	"github.com/saadahmedbd/Treestore/Config"
	rest "github.com/saadahmedbd/Treestore/Rest"
	buyerhandler "github.com/saadahmedbd/Treestore/Rest/Handler/BuyerHandler"
	cartitemhandler "github.com/saadahmedbd/Treestore/Rest/Handler/CartItemHandler"
)

func Server() {
	cnf := Config.GetConfig()
	buyerhandler := buyerhandler.NewHandler()
	cartitemhandler := cartitemhandler.NewHandler()
	server := rest.NewServer(buyerhandler, cartitemhandler)
	server.Start(cnf)

}
