package database

import (
	"github.com/aqaurius6666/faas-pkg/entity"
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	RepositorySet = wire.NewSet(NewGormRepository, wire.Bind(new(Repository), new(*RepositoryGormImpl)))
)

type DSN string

type Repository interface {
	Migrate() error
	UserRepository
}

type RepositoryGormImpl struct {
	db *gorm.DB
}

func NewGormRepository(dsn DSN) (*RepositoryGormImpl, error) {
	db, err := gorm.Open(postgres.Open(string(dsn)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &RepositoryGormImpl{
		db,
	}, err
}

func (r *RepositoryGormImpl) Migrate() error {
	return r.db.AutoMigrate(&entity.User{})
}
