//go:build wireinject
// +build wireinject

package model

import (
	"github.com/aqaurius6666/faas-pkg/database"
	"github.com/google/wire"
)

type Option struct {
	Repository database.Repository
}

func NewModel(o Option) (Model, error) {
	wire.Build(
		wire.FieldsOf(&o, "Repository"),
		ModelSet,
	)
	return nil, nil
}
