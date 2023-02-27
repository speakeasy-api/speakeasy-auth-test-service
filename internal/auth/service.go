package auth

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/speakeasy-api/speakeasy-auth-test-service/pkg/models"
)

func HandleAuth(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		handleError(w, err)
		return
	}

	var req models.AuthRequest
	if err := json.Unmarshal(body, &req); err != nil {
		handleError(w, err)
		return
	}

	if err := checkAuth(req, r); err != nil {
		handleError(w, err)
		return
	}
}

func handleError(w http.ResponseWriter, err error) {
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
