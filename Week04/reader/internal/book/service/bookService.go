package service

import (
	"context"
	pb "reader/api/book"
	"reader/internal/book/biz"
)

type BookService struct {
	pb.UnimplementedBookServer
	buc *biz.BookUseCase
}

func NewBookService(buc *biz.BookUseCase) *BookService {
	return &BookService{buc: buc}
}

func (s *BookService) AddBook(ctx context.Context, in *pb.BookReq) (*pb.BookReply, error) {
	account := &biz.Book{
		ID:   in.GetId(),
		Name: in.GetName(),
	}
	err := s.buc.Add(ctx, account)
	return &pb.BookReply{}, err
}
