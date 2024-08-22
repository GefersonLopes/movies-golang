package middlewares

import (
	"encoding/json"
	"net/http"
)

func HandleErros(err error, w http.ResponseWriter) {
	if err != nil {
		w.Header().Set("Content-Type", "application/json")

		if validationErr, ok := err.(*ValidationError); ok {
			response := map[string]interface{}{
				"message": validationErr.Message,
				"status":  validationErr.StatusCode,
			}
			w.WriteHeader(validationErr.StatusCode)
			json.NewEncoder(w).Encode(response)    
		} else {
		
			response := map[string]interface{}{
				"message": err.Error(),
				"status":  http.StatusInternalServerError,
			}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
		}
		return
	}
}
