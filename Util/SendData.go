package util

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, Data interface{}, statusCode int) {

	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(Data)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
