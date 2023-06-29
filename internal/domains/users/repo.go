package users

import (
	"context"

	"github.com/dragun-igor/project-layout/internal/entities"
)

type Repo struct{}

func (r *Repo) Insert(ctx context.Context, user entities.User) (entities.User, error) {
	//Inserting and returning
	return entities.User{}, nil
}
