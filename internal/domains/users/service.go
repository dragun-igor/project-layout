package users

import (
	"context"
	"fmt"

	"github.com/dragun-igor/project-layout/internal/entities"
)

type Creator interface {
	Insert(ctx context.Context, user entities.User) (entities.User, error)
}

type Service struct {
	userCreator Creator
}

func NewService(userCreator Creator) *Service {
	return &Service{userCreator: userCreator}
}

func (s *Service) Create(ctx context.Context, user entities.User) (entities.User, error) {
	createdUser, err := s.userCreator.Insert(ctx, user)
	if err != nil {
		return entities.User{}, fmt.Errorf("Create: %w", err)
	}
	return createdUser, nil
}
