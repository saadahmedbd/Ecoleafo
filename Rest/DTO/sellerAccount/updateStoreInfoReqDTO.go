package selleraccount

type UpdateStoreInfoRequest struct {
	StoreName   string `json:"store_name"`
	StoreDesc   string `json:"store_desc"`
	StoreLogo   string `json:"store_logo"`
	StoreBanner string `json:"store_banner"`
}
