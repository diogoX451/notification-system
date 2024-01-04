package dtos

type CreateGroupDTO struct {
	Name string `json:"name" validate:"required"`
}

type CreateMessageDTO struct {
	Message string `json:"message" validate:"required"`
}
