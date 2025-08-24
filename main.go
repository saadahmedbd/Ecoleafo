package main

import (
	"github.com/saadahmedbd/Treestore/CMD"
	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func main() {
	// Run Auto Migration
	Config.Connect()

	CMD.Server()

	// insert user(seller data)
	Pass := "saadahmed"
	password, err := util.HashPassword(Pass)
	if err != nil {
		panic(err)
	}
	user1 := models.User{
		ID:         4,
		RoleID:     2,
		FirstName:  "saad",
		LastName:   "ahmed",
		Email:      "saadahmed@gmail.com",
		Password:   password,
		Phone:      "012345677",
		StoreName:  "saad astore",
		StoreDesc:  "digital product",
		IsActive:   true,
		IsVerified: true,
	}

	if err := models.InsertData(Config.DB, user1); err != nil {
		panic(err)
	}

}
