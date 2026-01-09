package cartitem

type CartMessage struct {
	Type    string `json:"type"` // info, warning, error, success
	Message string `json:"message"`
	Action  string `json:"action,omitempty"` // Optional action button
}
