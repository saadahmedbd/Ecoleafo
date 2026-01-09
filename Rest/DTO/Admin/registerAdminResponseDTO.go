package admin

type RegisterAdminResponse struct {
	Admin *AdminResponse `json:"admin"`
	Token string               `json:"token"`
}