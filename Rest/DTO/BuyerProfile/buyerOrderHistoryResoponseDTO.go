package buyerprofile

type BuyerOrderHistoryResponse struct {
	Orders     []OrderInfo    `json:"orders"`
	Pagination PaginationInfo `json:"pagination"`
}
