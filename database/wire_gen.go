// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package database

// Injectors from wire.go:

func NewDatabase(o Option) (Repository, error) {
	dsn := o.Dsn
	repositoryGormImpl, err := NewGormRepository(dsn)
	if err != nil {
		return nil, err
	}
	return repositoryGormImpl, nil
}

// wire.go:

type Option struct {
	Dsn DSN
}
