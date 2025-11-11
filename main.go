package main

import (
	"log"

	"github.com/saadahmedbd/Treestore/CMD"
	"github.com/saadahmedbd/Treestore/Config"
	"github.com/saadahmedbd/Treestore/Database"
	util "github.com/saadahmedbd/Treestore/Util"
)

func main() {
	// Run Auto Migration
	Config.Connect()

	// Seed super admin
	database.SeedSuperAdminFromEnv(Config.DB)
	// Initialize Cloudinary service
	log.Println("Initializing Cloudinary...")
	if err := Config.InitCloudinary(); err != nil {
		log.Fatal("Failed to initialize Cloudinary:", err)
	}
	util.RunTokenCleanup(Config.DB)
	CMD.Server()

}
