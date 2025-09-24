package CMD

import (
	"github.com/saadahmedbd/Treestore/Config"
	rest "github.com/saadahmedbd/Treestore/Rest"
	buyerhandler "github.com/saadahmedbd/Treestore/Rest/Handler/BuyerHandler"
)

func Server() {
	cnf := Config.GetConfig()
	buyerhandler := buyerhandler.NewHandler()
	server := rest.NewServer(buyerhandler)
	server.Start(cnf)

}
