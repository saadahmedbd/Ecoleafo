package CMD

import (
	"github.com/saadahmedbd/Treestore/Config"
	rest "github.com/saadahmedbd/Treestore/Rest"
)

func Server() {
	cnf := Config.GetConfig()
	rest.Start(cnf)

}
