package main

import (
	"github.com/saadahmedbd/Treestore/CMD"
	"github.com/saadahmedbd/Treestore/Config"
	"github.com/saadahmedbd/Treestore/Database"
)

func main() {
	// Run Auto Migration
	Config.Connect()

	// Seed super admin
	database.SeedSuperAdminFromEnv(Config.DB)
	// Initialize Cloudinary service
	Config.InitializeCloudinary()
	CMD.Server()

}
