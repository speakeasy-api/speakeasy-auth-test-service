package pagination

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type LimitOffsetRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Page   int `json:"page"`
}

type CursorRequest struct {
	Cursor int `json:"cursor"`
}

type PaginationResponse struct {
	NumPages    int   `json:"numPages"`
	ResultArray []int `json:"resultArray"`
}

const total = 20

func HandleLimitOffsetPage(w http.ResponseWriter, r *http.Request) {
	queryLimit := r.FormValue("limit")
	queryPage := r.FormValue("page")

	var pagination LimitOffsetRequest
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

	res := PaginationResponse{
		NumPages:    total / limit,
		ResultArray: make([]int, 0),
	}

	for i := start; i < total && len(res.ResultArray) < limit; i++ {
		res.ResultArray = append(res.ResultArray, i)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(500)
	}
}

func HandleLimitOffsetOffset(w http.ResponseWriter, r *http.Request) {
	queryLimit := r.FormValue("limit")
	queryOffset := r.FormValue("offset")

	var pagination LimitOffsetRequest
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

	res := PaginationResponse{
		NumPages:    total / limit,
		ResultArray: make([]int, 0),
	}

	for i := offset; i < total && len(res.ResultArray) < limit; i++ {
		res.ResultArray = append(res.ResultArray, i)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(500)
	}
}

func HandleCursor(w http.ResponseWriter, r *http.Request) {
	queryCursor := r.FormValue("cursor")

	var pagination CursorRequest
	hasBody := true
	if err := json.NewDecoder(r.Body).Decode(&pagination); err != nil {
		hasBody = false
	}

	cursor, err := getValue(queryCursor, hasBody, pagination.Cursor, w)
	if err != nil {
		return
	}

	res := PaginationResponse{
		NumPages:    0,
		ResultArray: make([]int, 0),
	}

	for i := cursor + 1; i < total && len(res.ResultArray) < 15; i++ {
		res.ResultArray = append(res.ResultArray, i)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(500)
	}
}

func getValue(queryValue string, hasBody bool, paginationValue int, w http.ResponseWriter) (int, error) {
	if hasBody {
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
