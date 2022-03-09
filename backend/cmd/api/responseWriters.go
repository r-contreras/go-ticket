package api

import (
	"encoding/json"
	"net/http"
)

func (*Application) WriteJsonResponse(wrap string, object interface{}, status int, w http.ResponseWriter) error {
	wrapper := make(map[string]interface{})

	wrapper[wrap] = object

	json, err := json.Marshal(wrapper)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json)
	return nil
}

func (app *Application) WriteBadRequest(w http.ResponseWriter) error {
	return app.WriteJsonResponse("error", "Bad Request", http.StatusBadRequest, w)
}
