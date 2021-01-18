package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func getPayloads(r *http.Request, payload interface{}) (interface{}, error) {

	err := json.NewDecoder(r.Body).Decode(&payload)
	defer r.Body.Close()

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return payload, nil
}

// JSONResponse - converts response to json type
func JSONResponse(w http.ResponseWriter, data []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}
