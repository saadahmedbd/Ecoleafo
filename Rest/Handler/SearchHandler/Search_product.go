package searchhandler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

// PaginationMetadata holds metadata for the paginated response
type PaginationMetadata struct {
	Total       int64 `json:"total"`
	PerPage     int   `json:"per_page"`
	CurrentPage int   `json:"current_page"`
	TotalPages  int   `json:"total_pages"`
}

// SearchResponse wraps the products and pagination metadata
type SearchResponse struct {
	Products []models.Product   `json:"products"`
	Meta     PaginationMetadata `json:"meta"`
}

func (h *Handler) Search_Product(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(r.URL.Query().Get("q"))
	category := r.URL.Query().Get("category_id")
	priceMin := r.URL.Query().Get("min")
	priceMax := r.URL.Query().Get("max")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}
	perpage, _ := strconv.Atoi(r.URL.Query().Get("perpage"))
	if perpage < 1 {
		perpage = 10
	}

	var sqlQuery strings.Builder
	var countQuery strings.Builder
	var args []interface{}

	sqlQuery.WriteString(`
		SELECT id, name, description, category_id, price, quantity
		FROM products
		WHERE 1=1
	`)
	//COUNT TOTAL PRODUCT
	countQuery.WriteString(`
			SELECT COUNT (*)
			FROM products
			WHERE 1=1
		`)

	// Full-text search
	if query != "" {
		sqlQuery.WriteString(" AND search_vector @@ plainto_tsquery(?)")
		countQuery.WriteString(" AND search_vector @@ plainto_tsquery(?)")
		args = append(args, query)

	}

	// Filter by category
	if category != "" {
		sqlQuery.WriteString(" AND category_id = ?")
		countQuery.WriteString(" AND category_id = ?")
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
		countQuery.WriteString(" AND price BETWEEN ? AND ?")
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

	// Debugging: Log the queries and args
	log.Printf("Count Query: %s, Args: %v", countQuery.String(), args)
	log.Printf("Data Query: %s, Args: %v", sqlQuery.String(), args)

	//GET TOTAL COUNT
	var total int64
	if err := Config.DB.Raw(countQuery.String(), args...).Scan(&total).Error; err != nil {
		http.Error(w, fmt.Sprintf("Count query failed: %v", err), http.StatusInternalServerError)
		return
	}
	// Calculate offset and apply pagination
	offset := (page - 1) * perpage
	sqlQuery.WriteString(fmt.Sprintf(" LIMIT %d OFFSET %d", perpage, offset))

	//execute product query
	var products []models.Product
	result := Config.DB.Debug().Raw(sqlQuery.String(), args...).Scan(&products)
	if result.Error != nil {
		http.Error(w, fmt.Sprintf("Search failed: %v", result.Error), http.StatusInternalServerError)
		return
	}

	// Calculate total pages
	totalPages := int((total + int64(perpage) - 1) / int64(perpage))
	// Prepare response
	response := SearchResponse{
		Products: products,
		Meta: PaginationMetadata{
			Total:       total,
			PerPage:     perpage,
			CurrentPage: page,
			TotalPages:  totalPages,
		},
	}

	util.SendData(w, response, 200)
}
