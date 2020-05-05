package models

type SecureMessageRequest struct {
	Message              string `json:"message" validate:"required"`
	SelfDestruct         int    `json:"self_destruct" validate:"oneof=0 1 3 6 12 24 48 72 168 720"`
	Password             string `json:"password" validate:"omitempty,min=8,max=24"`
	DestructAfterOpening bool   `json:"destruct_after_opening" validate:"omitempty"`
}

type SecureMessageResponse struct {
	Id  string `json:"id"`
	URL string `json:"url"`
}

type SecureMessage struct {
	Message              []byte
	KeyId                string
	Password             string
	DestructAfterOpening bool
}
