package models

//NoteCreatedResponse API response struct
type NoteCreatedResponse struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
	Data    `json:"Data"`
	Error   []string `json:"Error"`
}

//Data API data struct (holds the main data)
type Data struct {
	Id  string `json:"id"`
	URL string `json:"url"`
}
