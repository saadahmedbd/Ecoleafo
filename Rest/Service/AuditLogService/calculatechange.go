package auditlogservice

import "fmt"

// Helper: Calculate changes between old and new values
func (s *auditLogService) calculateChanges(oldValues, newValues map[string]interface{}) map[string]interface{} {
	changes := make(map[string]interface{})

	for key, newVal := range newValues {
		if oldVal, exists := oldValues[key]; exists {
			if fmt.Sprintf("%v", oldVal) != fmt.Sprintf("%v", newVal) {
				changes[key] = map[string]interface{}{
					"from": oldVal,
					"to":   newVal,
				}
			}
		} else {
			changes[key] = map[string]interface{}{
				"from": nil,
				"to":   newVal,
			}
		}
	}

	return changes
}
