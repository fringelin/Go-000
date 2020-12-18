package data

import (
	"context"
	"github.com/pkg/errors"
	"reader/internal/account/biz"
	"reader/internal/account/data/ent"
)

var _ biz.AccountRepo = (*AccountRepo)(nil)

type AccountRepo struct {
	db *ent.Client
}

func NewAccountRepo(db *ent.Client) *AccountRepo {
	return &AccountRepo{db: db}
}

func (r *AccountRepo) SaveAccount(ctx context.Context, account *biz.Account) error {
	_, err := r.db.Account.Create().SetID(account.ID).SetName(account.Name).Save(ctx)
	if err != nil {
		return errors.Wrap(err, "create account failed")
	}
	return nil
}
