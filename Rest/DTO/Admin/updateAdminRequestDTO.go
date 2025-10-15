package admin

type UpdateAdminRequest struct {
	FullName   string `json:"full_name"`
	Phone      string `json:"phone"`
	Department string `json:"department"`
}
