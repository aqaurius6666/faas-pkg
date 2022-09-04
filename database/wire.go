//go:build wireinject
// +build wireinject

package database

import "github.com/google/wire"

type Option struct {
	Dsn DSN
}

func NewDatabase(o Option) (Repository, error) {
	wire.Build(
		wire.FieldsOf(&o, "Dsn"),
		RepositorySet,
	)
	return nil, nil
}
