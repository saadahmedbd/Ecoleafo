package reviewdto

// PaginationInfo - DTO for pagination information
type PaginationInfo struct {
	Page       int   `json:"page"`
	PerPage    int   `json:"per_page"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

// ReviewListResponse - DTO for paginated review list
type ReviewListResponse struct {
	Reviews    []ReviewResponse `json:"reviews"`
	Pagination PaginationInfo   `json:"pagination"`
}
