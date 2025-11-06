package sellerdashboard

// DashboardFilter - Query parameters for filtering
type DashboardFilter struct {
	Period    string `json:"period"`     // today, week, month, year
	StartDate string `json:"start_date"` // YYYY-MM-DD
	EndDate   string `json:"end_date"`   // YYYY-MM-DD
	Limit     int    `json:"limit"`
}
