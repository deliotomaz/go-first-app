package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseBadRequest(w http.ResponseWriter) {
	responseWithJSON(w, 400, map[string]interface{}{"error": "Bad Request", "success": false})
}
func ResponseWithError(w http.ResponseWriter, code int, message string) {
	responseWithJSON(w, code, map[string]interface{}{"error": message, "success": false})
}
func ResponseOK(w http.ResponseWriter, returned interface{}) {

	responseWithJSON(w, 200, map[string]interface{}{"result": returned, "success": true})

}
func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
