package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Tour struct {
	Msg string
}

type TourRepo interface {
	Ai(ctx context.Context, t *Tour) (string, error)
}

type TourUsecase struct {
	repo TourRepo
	log  *log.Helper
}

func NewTourUsecase(repo TourRepo, logger log.Logger) *TourUsecase {
	return &TourUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *TourUsecase) GetAiResp(ctx context.Context, t *Tour) (string, error) {
	uc.log.WithContext(ctx).Infof("GetAiResp: %v", t.Msg)
	return uc.repo.Ai(ctx, t)
}
