package biz

import (
	"context"
)

type Book struct {
	ID   int64
	Name string
}

type BookRepo interface {
	SaveBook(context.Context, *Book) error
}

func NewBookUseCase(repo BookRepo) *BookUseCase {
	return &BookUseCase{repo: repo}
}

type BookUseCase struct {
	repo BookRepo
}

func (uc *BookUseCase) Add(ctx context.Context, book *Book) error {
	return uc.repo.SaveBook(ctx, book)
}
