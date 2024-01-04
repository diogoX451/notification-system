package services

import (
	dtos "notification-system/internal/dtos/group"
	"notification-system/internal/entitys"
	"notification-system/internal/repositories"
)

type IGroupService interface {
	Create(group *dtos.CreateGroupDTO) (*dtos.CreateMessageDTO, error)
}

type GroupService struct {
	repo repositories.IGroupRepository
}

func NewGroupService(repo repositories.IGroupRepository) *GroupService {
	return &GroupService{repo}
}

func (s *GroupService) Create(group *dtos.CreateGroupDTO) (*dtos.CreateMessageDTO, error) {

	groupEntity := entitys.Groups{
		Name: group.Name,
	}

	_, err := s.repo.Create(&groupEntity)

	if err != nil {
		return nil, err
	}
	return &dtos.CreateMessageDTO{
		Message: "Group created successfully",
	}, nil
}
