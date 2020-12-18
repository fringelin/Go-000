package service

import (
	"context"
	pb "reader/api/account"
	"reader/internal/account/biz"
)

type AccountService struct {
	pb.UnimplementedAccountServer
	auc *biz.AccountUseCase
}

func NewAccountService(auc *biz.AccountUseCase) *AccountService {
	return &AccountService{auc: auc}
}

func (s *AccountService) AddAccount(ctx context.Context, in *pb.AccountReq) (*pb.AccountReply, error) {
	account := &biz.Account{
		ID:   in.GetId(),
		Name: in.GetName(),
	}
	err := s.auc.Add(ctx, account)
	return &pb.AccountReply{}, err
}
