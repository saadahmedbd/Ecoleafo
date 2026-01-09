package order

import "time"

type OrderHistoryResponse struct {
	ID        uint      `json:"id"`
	Status    string    `json:"status"`
	Comment   string    `json:"comment"`
	UpdatedBy string    `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
}
