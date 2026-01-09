package searchhandler

import (
	"net/http"
	"strings"

	"github.com/saadahmedbd/Treestore/Config"
	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
)

type AutocompleteResponse struct {
	Suggestions []models.Product `json:"suggestions"`
}

func (h *Handler) Autocomplete(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(r.URL.Query().Get("q"))
	
	if query == "" {
		util.SendData(w, AutocompleteResponse{Suggestions: []models.Product{}}, 200)
		return
	}

	searchPattern := "%" + strings.ToLower(query) + "%"
	startPattern := strings.ToLower(query) + "%"

	type IDResult struct {
		ID uint
	}
	var idResults []IDResult
	if err := Config.DB.Raw(`
		SELECT id
		FROM products
		WHERE LOWER(name) LIKE ? OR LOWER(description) LIKE ?
		ORDER BY 
			CASE 
				WHEN LOWER(name) LIKE ? THEN 1
				WHEN LOWER(name) LIKE ? THEN 2
				ELSE 3
			END
		LIMIT 10
	`, searchPattern, searchPattern, startPattern, searchPattern).Scan(&idResults).Error; err != nil {
		util.SendData(w, AutocompleteResponse{Suggestions: []models.Product{}}, 200)
		return
	}

	var products []models.Product
	if len(idResults) > 0 {
		var productIDs []uint
		for _, r := range idResults {
			productIDs = append(productIDs, r.ID)
		}
		Config.DB.Preload("Images").Where("id IN ?", productIDs).Find(&products)
	}

	util.SendData(w, AutocompleteResponse{Suggestions: products}, 200)
}
