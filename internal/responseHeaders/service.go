package responseHeaders

import (
	"encoding/json"
	"github.com/speakeasy-api/speakeasy-auth-test-service/internal/utils"
	"io"
	"net/http"
)

func HandleVendorJsonResponseHeaders(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	var obj interface{}

	err = json.Unmarshal([]byte("{\"name\":\"Panda\"}"), &obj)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/vnd.api+json; charset=utf-8")

	if err := json.NewEncoder(w).Encode(obj); err != nil {
		utils.HandleError(w, err)
	}
}
