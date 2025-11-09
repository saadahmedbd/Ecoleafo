package inventoryhandler

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"time"

	inventory "github.com/saadahmedbd/Treestore/Rest/DTO/Inventory"
	util "github.com/saadahmedbd/Treestore/Util"
	"github.com/saadahmedbd/Treestore/constants"
)

// ExportInventory - GET /api/seller/inventory/export
func (h *Inventoryhandler) ExportInventory(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value(constants.ContextKeyUserID)
	userRoleVal := r.Context().Value(constants.ContextKeyRole)
	if userIDVal == nil || userRoleVal == nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	sellerID, ok := userIDVal.(uint)
	if !ok {
		http.Error(w, `{"error":"invalid user_id type"}`, http.StatusInternalServerError)
		return
	}

	// Get all inventory
	inventory, err := h.inventoryservice.GetInventory(sellerID, inventory.InventoryFilter{
		Limit: 10000, // Get all
	})
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Set CSV headers
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=inventory_%s.csv", time.Now().Format("20060102")))

	// Create CSV writer
	writer := csv.NewWriter(w)
	defer writer.Flush()

	// Write header
	writer.Write([]string{
		"ID", "Name", "SKU", "Category", "Price", "Stock",
		"Low Stock Threshold", "Stock Value", "Status", "Last Updated",
	})

	// Write data
	for _, item := range inventory.Items {
		status := "In Stock"
		if item.Stock == 0 {
			status = "Out of Stock"
		} else if item.Stock <= item.LowStockThreshold {
			status = "Low Stock"
		}

		writer.Write([]string{
			fmt.Sprintf("%d", item.ID),
			item.Name,
			item.SKU,
			item.CategoryName,
			fmt.Sprintf("%.2f", item.Price),
			fmt.Sprintf("%d", item.Stock),
			fmt.Sprintf("%d", item.LowStockThreshold),
			fmt.Sprintf("%.2f", item.StockValue),
			status,
			item.LastUpdated.Format("2006-01-02 15:04:05"),
		})
	}
}
