package biz

import (
	"context"
)

type Account struct {
	ID   int64
	Name string
}

type AccountRepo interface {
	SaveAccount(context.Context, *Account) error
}

func NewAccountUseCase(repo AccountRepo) *AccountUseCase {
	return &AccountUseCase{repo: repo}
}

type AccountUseCase struct {
	repo AccountRepo
}

func (uc *AccountUseCase) Add(ctx context.Context, account *Account) error {
	return uc.repo.SaveAccount(ctx, account)
}
