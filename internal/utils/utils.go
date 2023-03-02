package utils

import (
	"encoding/json"
	"errors"
	"github.com/speakeasy-api/speakeasy-auth-test-service/pkg/models"
	"log"
	"net/http"
)

var authError = errors.New("invalid auth")

func HandleError(w http.ResponseWriter, err error) {
	log.Println(err)

	data, marshalErr := json.Marshal(models.ErrorResponse{
		Error: models.ErrorMessage{
			Message: err.Error(),
		},
	})
	if marshalErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if errors.Is(err, authError) {
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, _ = w.Write(data)
}
