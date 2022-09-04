package model

import (
	"github.com/aqaurius6666/faas-pkg/database"
	"github.com/google/wire"
)

var (
	ModelSet = wire.NewSet(wire.Struct(new(ModelImpl), "*"), wire.Bind(new(Model), new(*ModelImpl)))
)

type Model interface {
	User
}

type ModelImpl struct {
	Repository database.Repository
}
