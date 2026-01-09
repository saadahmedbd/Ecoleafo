package cartitem

type BulkUpdateCartRequest struct {
	Item []struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	} `json:"items"`
}
