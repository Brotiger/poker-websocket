package model

type Message struct {
	Type   string `json:"type"`
	Header struct {
		AccessToken string `json:"access_token"`
		JoinToken   string `json:"join_token"`
	} `json:"header"`
}
