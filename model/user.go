package model

import (
	"context"

	"github.com/aqaurius6666/faas-pkg/entity"
)

type User interface {
	GetUserById(ctx context.Context, id int64) (*entity.User, error)
	InsertUser(ctx context.Context, user *entity.User) (*entity.User, error)
}

func (s *ModelImpl) GetUserById(ctx context.Context, id int64) (*entity.User, error) {
	query := &entity.UserQuery{}
	query.ID = id
	return s.Repository.GetUser(ctx, query)
}

func (s *ModelImpl) InsertUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return s.Repository.InsertUser(ctx, user)
}
