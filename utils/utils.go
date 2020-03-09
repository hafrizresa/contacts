package utils

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, msg string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": msg}
}

func Response(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-type", "Application/json")
	json.NewEncoder(w).Encode(data)
}
