package dtos

type CreateCompanyDTO struct {
	Name string `json:"name" validate:"required"`
}

type CreateMessageDTO struct {
	Message string `json:"message"`
}
