package echo

type StringResponse struct {
	Response *string `json:"response,omitempty"`
}

func NewStringResponse() *StringResponse {
	return &StringResponse{}
}

func NewStringResponseFill(response *string) *StringResponse {
	return &StringResponse{Response: response}
}
