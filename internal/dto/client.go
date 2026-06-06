package dto

type CreateClientInput struct {
	Name string `json:"name"`

	Email string `json:"email"`

	Phone string `json:"phone"`

	Document string `json:"document"`
}
