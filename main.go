package main

import (
	"github.com/saadahmedbd/Treestore/CMD"
	"github.com/saadahmedbd/Treestore/Config"
)

func main() {
	// Run Auto Migration
	Config.Connect()

	// jwt, err := authhandler.CreateJwt(
	// 	"my-secret", authhandler.Payload{
	// 		Sub:       23,
	// 		FirstName: "Saad",
	// 		LastName:  "Ahmed",
	// 		Exp:       1234567890,
	// 		Iat:       1234567890,
	// 	})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(jwt)

	CMD.Server()

}
