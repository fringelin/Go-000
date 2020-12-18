// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/google/wire"
	"reader/internal/book/biz"
	"reader/internal/book/data"
	"reader/internal/book/service"
)

//go:generate wire
func NewBookService() (*service.BookService, func(), error) {
	panic(wire.Build(
		wire.NewSet(data.NewDB),
		wire.NewSet(data.NewBookRepo, wire.Bind(new(biz.BookRepo), new(*data.BookRepo))),
		wire.NewSet(biz.NewBookUseCase),
		service.NewBookService,
	))
}
