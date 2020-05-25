package models

type NoteCreatedResponse struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
	Data    `json:"Data"`
	Error   []string `json:"Error"`
}

type Data struct {
	Id  string `json:"id"`
	URL string `json:"url"`
}
