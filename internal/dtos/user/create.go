package dtos

type CreateUserDTO struct {
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	GroupID   int64  `json:"group_id" validate:"required"`
	CompanyID int64  `json:"company_id" validate:"required"`
}

type CreateMessage struct {
	Message string `json:"message" validate:"required"`
}
