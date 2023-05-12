package pagination

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type PaginationRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Page   int `json:"page"`
}

type PaginationResponse struct {
	NumPages     int   `json:"numPages"`
	ResultsArray []int `json:"resultsArray"`
}

const total = 20

func HandleLimitOffsetPage(w http.ResponseWriter, r *http.Request) {
	queryLimit := r.FormValue("limit")
	queryPage := r.FormValue("page")

	var pagination PaginationRequest
	hasBody := true
	if err := json.NewDecoder(r.Body).Decode(&pagination); err != nil {
		hasBody = false
	}
	limit, err := getValue(queryLimit, hasBody, pagination.Limit, w)
	if err != nil {
		return
	}
	page, err := getValue(queryPage, hasBody, pagination.Page, w)
	if err != nil {
		return
	}

	start := (page - 1) * limit
	if start > total {
		w.WriteHeader(404)
	}

	res := PaginationResponse{
		NumPages:     total / limit,
		ResultsArray: make([]int, 0),
	}

	for i := start; i < total && len(res.ResultsArray) < limit; i++ {
		res.ResultsArray = append(res.ResultsArray, i)
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(500)
	}
}

func HandleLimitOffsetOffset(w http.ResponseWriter, r *http.Request) {
	queryLimit := r.FormValue("limit")
	queryOffset := r.FormValue("offset")

	var pagination PaginationRequest
	hasBody := true
	if err := json.NewDecoder(r.Body).Decode(&pagination); err != nil {
		hasBody = false
	}
	limit, err := getValue(queryLimit, hasBody, pagination.Limit, w)
	if err != nil {
		return
	}
	offset, err := getValue(queryOffset, hasBody, pagination.Offset, w)
	if err != nil {
		return
	}

	if offset > total {
		w.WriteHeader(404)
	}

	res := PaginationResponse{
		NumPages:     total / limit,
		ResultsArray: make([]int, 0),
	}

	for i := offset; i < total && len(res.ResultsArray) < limit; i++ {
		res.ResultsArray = append(res.ResultsArray, i)
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(500)
	}
}

func getValue(queryValue string, hasBody bool, paginationValue int, w http.ResponseWriter) (int, error) {
	if hasBody && queryValue == "" {
		return paginationValue, nil
	} else {
		value, err := strconv.Atoi(queryValue)
		if err != nil {
			w.WriteHeader(400)
			return 0, err
		}
		return value, nil
	}
}
