package data

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"reader/internal/book/data/ent"
)

func NewDB() (*ent.Client, func(), error) {
	client, err := ent.Open("sqlite3", "sqlite3/reader:book?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return nil, nil, err
	}
	f := func() {
		client.Close()
	}
	// Run the auto migration tool.
	err = client.Schema.Create(context.Background())

	return client, f, err
}
