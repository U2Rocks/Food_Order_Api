package models

// this struct is used to send JSON messages back to clients
type ResponseObject struct {
	Statuscode int    `json:"statuscode"`
	Message    string `json:"message"`
}
