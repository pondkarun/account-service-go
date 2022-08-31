package Respond

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	mar, _ := json.Marshal(data)
	output := string(mar)
	fmt.Fprint(w, output)
}
