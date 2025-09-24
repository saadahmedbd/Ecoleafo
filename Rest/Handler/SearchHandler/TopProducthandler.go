package searchhandler

import (
	"net/http"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func (h *Handler) TopProduct(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	err := Config.DB.Preload("Seller").Preload("Category").Preload("CartItems").Preload("OrderItems").Preload("Reviews").Raw(`
		SELECT p.*, COALESCE(AVG(r.rating), 0) AS avg_rating
		FROM products p
		LEFT JOIN reviews r ON r.product_id = p.id
		GROUP BY p.id
		ORDER BY avg_rating DESC
		LIMIT 10
	`).Scan(&products).Error
	if err != nil {
		http.Error(w, "not found", http.StatusInternalServerError)
		return
	}
	util.SendData(w, products, 200)
}
