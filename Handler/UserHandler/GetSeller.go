package userhandler

import (
	"fmt"
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is Home page")
}
// get seller function
func GetSeller(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Please provide valid Method", http.StatusBadRequest)
		return
	}
	var Sellers []models.User
	db := Config.DB
	db.Find(&Sellers)
	util.SendData(w, Sellers, 200)

}
