package main

import (
	"github.com/saadahmedbd/Treestore/CMD"
	"github.com/saadahmedbd/Treestore/Config"
)

func main() {
	// Run Auto Migration
	Config.Connect()

	CMD.Server()

}
