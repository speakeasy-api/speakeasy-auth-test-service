package requestbody

import (
	"encoding/json"
	"github.com/speakeasy-api/speakeasy-auth-test-service/internal/utils"
	"io"
	"net/http"
)

func HandleRequestBody(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := json.NewEncoder(w).Encode(body); err != nil {
		utils.HandleError(w, err)
	}
}
