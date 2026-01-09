package auditlogservice

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"time"

	auditlogdto "github.com/saadahmedbd/Treestore/Rest/DTO/AuditLogDTO"
)

// Export logs
func (s *auditLogService) ExportLogs(req auditlogdto.ExportAuditLogsRequest) ([]byte, string, error) {
	logs, _, err := s.auditRepo.GetAll(req.Filters)
	if err != nil {
		return nil, "", err
	}

	var data []byte
	var filename string

	if req.Format == "json" {
		data, _ = json.MarshalIndent(logs, "", "  ")
		filename = fmt.Sprintf("audit_logs_%s.json", time.Now().Format("20060102"))
	} else {
		// CSV format
		var csvData [][]string
		csvData = append(csvData, []string{
			"ID", "Date", "Actor", "Actor Type", "Action", "Entity Type",
			"Entity ID", "Description", "IP Address", "Status", "Severity",
		})

		for _, log := range logs {
			csvData = append(csvData, []string{
				fmt.Sprintf("%d", log.ID),
				log.CreatedAt.Format("2006-01-02 15:04:05"),
				log.ActorName,
				log.ActorType,
				log.Action,
				log.EntityType,
				fmt.Sprintf("%d", *log.EntityID),
				log.Description,
				log.IPAddress,
				log.Status,
				log.Severity,
			})
		}

		var buf bytes.Buffer
		writer := csv.NewWriter(&buf)
		writer.WriteAll(csvData)
		writer.Flush()
		data = buf.Bytes()
		filename = fmt.Sprintf("audit_logs_%s.csv", time.Now().Format("20060102"))
	}

	return data, filename, nil
}
