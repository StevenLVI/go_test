package response

import (
	"encoding/json"
	"net/http"
)

// Estandarizar respuesta de error de las API
type ErrorMessage struct {
	Message string `json:"message"`
}

type Map map[string]interface{}

// Estandarizar respuesta JSON para las API.
func JSON(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) error {
	if data == nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(statusCode)
		return nil
	}

	j, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(j)
	return nil
}

// Retornar error en formato JSON
func HTTPError(w http.ResponseWriter, r *http.Request, statusCode int, message string) error {
	msg := ErrorMessage{
		Message: message,
	}

	return JSON(w, r, statusCode, msg)
}
