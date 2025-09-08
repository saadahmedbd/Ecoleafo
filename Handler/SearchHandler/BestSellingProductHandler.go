package searchhandler

import (
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func BestSellingProduct(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	result := Config.DB.Raw(`
		SELECT p.*, SUM(oi.quantity) AS total_sold
		FROM products p
		JOIN order_items oi ON oi.product_id = p.id
		GROUP BY p.id
		ORDER BY total_sold DESC
		LIMIT 10
	`).Scan(&products)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	util.SendData(w, products, 200)
}
