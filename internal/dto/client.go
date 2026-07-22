package dto

// dto para criação, atualização e patch de clientes

type CreateClientInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required"`
	Document string `json:"document" binding:"required"`
}

type UpdateClientInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required"`
	Document string `json:"document" binding:"required"`
}

type PatchClientInput struct {
	Name     *string `json:"name" binding:"omitempty,min=1"`
	Email    *string `json:"email" binding:"omitempty,email"`
	Phone    *string `json:"phone" binding:"omitempty,min=1"`
	Document *string `json:"document" binding:"omitempty,min=1"`
}
