package data

import (
	"context"
	"github.com/pkg/errors"
	"reader/internal/book/biz"
	"reader/internal/book/data/ent"
)

var _ biz.BookRepo = (*BookRepo)(nil)

type BookRepo struct {
	db *ent.Client
}

func NewBookRepo(db *ent.Client) *BookRepo {
	return &BookRepo{db: db}
}

func (r *BookRepo) SaveBook(ctx context.Context, book *biz.Book) error {
	_, err := r.db.Book.Create().SetID(book.ID).SetName(book.Name).Save(ctx)
	if err != nil {
		return errors.Wrap(err, "create book failed")
	}
	return nil
}
