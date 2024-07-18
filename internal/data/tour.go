package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"tour/internal/biz"
	"tour/pkg/utils"
)

type tourRepo struct {
	data *Data
	log  *log.Helper
}

func NewTourRepo(data *Data, logger log.Logger) biz.TourRepo {
	return &tourRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (t *tourRepo) Ai(ctx context.Context, to *biz.Tour) (string, error) {
	resp := utils.Gpt(to.Msg, t.data.Ai.Key, t.data.Ai.Model, t.data.Ai.Url)
	return resp, nil
}
