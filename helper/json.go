package helper

import (
	"encoding/json"
	"net/http"
)

func ReadRequestBody(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	if err != nil {
		panic(err)
	}

}

func WriteResponseBody(write http.ResponseWriter, response interface{}) {
	write.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(write)
	err := encoder.Encode(response)
	if err != nil {
		http.Error(write, "Error", http.StatusInternalServerError)
	}
}
