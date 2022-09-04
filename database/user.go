package database

import (
	"context"

	"github.com/aqaurius6666/faas-pkg/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser(context.Context, *entity.UserQuery) (*entity.User, error)
	InsertUser(context.Context, *entity.User) (*entity.User, error)
}

func applyUserQuery(db *gorm.DB, query *entity.UserQuery) *gorm.DB {
	if query == nil {
		return db
	}
	if query.ID != 0 {
		db = db.Where(&entity.User{ID: query.ID})
	}
	return db
}

func (r *RepositoryGormImpl) GetUser(ctx context.Context, query *entity.UserQuery) (*entity.User, error) {
	user := new(entity.User)
	db := applyUserQuery(r.db, query)
	if err := db.WithContext(ctx).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *RepositoryGormImpl) InsertUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
