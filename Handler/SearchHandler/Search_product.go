package searchhandler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

func Search_Product(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(r.URL.Query().Get("q"))
	category := r.URL.Query().Get("category_id")
	priceMin := r.URL.Query().Get("min")
	priceMax := r.URL.Query().Get("max")

	var sqlQuery strings.Builder
	var args []interface{}

	sqlQuery.WriteString(`
		SELECT id, name, description, category_id, price, quantity
		FROM products
		WHERE 1=1
	`)

	// Full-text search
	if query != "" {
		sqlQuery.WriteString(" AND search_vector @@ plainto_tsquery(?)")
		args = append(args, query)

	}

	// Filter by category
	if category != "" {
		sqlQuery.WriteString(" AND category_id = ?")
		categoryID, err := strconv.Atoi(category)
		if err != nil {
			http.Error(w, "Invalid category_id", http.StatusBadRequest)
			return
		}
		args = append(args, categoryID)

	}

	// Filter by price range
	if priceMin != "" && priceMax != "" {
		sqlQuery.WriteString(" AND price BETWEEN ? AND ?")
		PriceMax, err := strconv.ParseFloat(priceMax, 64)
		if err != nil {
			http.Error(w, "invalid price max", http.StatusBadRequest)
			return
		}
		PriceMin, err := strconv.ParseFloat(priceMin, 64)
		if err != nil {
			http.Error(w, "invalid price max", http.StatusBadRequest)
			return
		}
		args = append(args, PriceMin, PriceMax)

	}

	// Debugging
	// fmt.Println("SQL:", sqlQuery.String())
	// fmt.Println("Args:", args)

	var products []models.Product
	result := Config.DB.Debug().Raw(sqlQuery.String(), args...).Scan(&products)
	if result.Error != nil {
		http.Error(w, fmt.Sprintf("Search failed: %v", result.Error), http.StatusInternalServerError)
		return
	}

	util.SendData(w, products, 200)
}
