package auditlogdto

// Export request
type ExportAuditLogsRequest struct {
	Format    string        `json:"format"` // csv, json
	StartDate string        `json:"start_date"`
	EndDate   string        `json:"end_date"`
	Filters   AuditLogQuery `json:"filters"`
}
