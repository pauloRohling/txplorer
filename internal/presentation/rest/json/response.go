package json

type HttpErrorResponse struct {
	Err      error  `json:"-"`
	Type     string `json:"type,omitempty"`
	Title    string `json:"title,omitempty"`
	Status   int    `json:"status,omitempty"`
	Detail   string `json:"detail,omitempty"`
	Instance string `json:"instance,omitempty"`
}

func (error *HttpErrorResponse) Error() string {
	return error.Err.Error()
}
