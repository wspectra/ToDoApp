package utils

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

type ResponseStruct struct {
	Status  string
	Message string
}

func HttpResponseWriter(w http.ResponseWriter, answer string, code int) {
	responseSrruct := ResponseStruct{}
	responseSrruct.Message = answer
	if code == 200 {
		responseSrruct.Status = "success"
	} else {
		responseSrruct.Status = "fail"
		log.Error().Msg(answer)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(responseSrruct)
}
