package util

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, Data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(Data)
}
