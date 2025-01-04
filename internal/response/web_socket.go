package response

type Header struct {
	Code  int    `json:"code"`
	Event string `json:"event"`
}

type Respons struct {
	Header Header `json:"header"`
	Body   any    `json:"body,omitempty"`
}
