package models

type ErrorMessage struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error ErrorMessage `json:"error"`
}

type HeaderAuth struct {
	HeaderName    string `json:"headerName"`
	ExpectedValue string `json:"expectedValue"`
}

type BasicAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthRequest struct {
	HeaderAuth []HeaderAuth `json:"headerAuth,omitempty"`
	BasicAuth  *BasicAuth   `json:"basicAuth,omitempty"`
}

type BodyRequest struct {
	ArrObjValue []any `json:"arrObjValue"`
}
