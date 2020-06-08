package models

//SecureMessageRequest model for requesting the API
type SecureMessageRequest struct {
	Note                 string `json:"message"`
	SelfDestruct         int    `json:"self_destruct"`
	DestructAfterOpening bool   `json:"destruct_after_opening"`
}
