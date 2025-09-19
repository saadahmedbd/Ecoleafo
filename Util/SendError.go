package util

import (
	"encoding/json"
	"net/http"
)

func SendError(w http.ResponseWriter, mge string, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(mge)

}
