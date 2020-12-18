// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/google/wire"
	"reader/internal/account/biz"
	"reader/internal/account/data"
	"reader/internal/account/service"
)

//go:generate wire
func NewAccountService() (*service.AccountService, func(), error) {
	panic(wire.Build(
		wire.NewSet(data.NewDB),
		wire.NewSet(data.NewAccountRepo, wire.Bind(new(biz.AccountRepo), new(*data.AccountRepo))),
		wire.NewSet(biz.NewAccountUseCase),
		service.NewAccountService,
	))
}
