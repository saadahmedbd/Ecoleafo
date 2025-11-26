package auditmiddleware

import (
	"strings"

	"github.com/saadahmedbd/Treestore/constants"
)

func (m *AuditMiddleware) determineAction(method, path string) string {
	switch method {
	case "POST":
		if strings.Contains(path, "login") {
			return constants.ActionLogin
		}
		return "create"
	case "PUT", "PATCH":
		return "update"
	case "DELETE":
		return "delete"
	default:
		return "view"
	}
}
