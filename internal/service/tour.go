package service

import (
	"context"
	"tour/internal/biz"

	pb "tour/api/tour/v1"
)

type TourService struct {
	pb.UnimplementedTourServer
	uc *biz.TourUsecase
}

func NewTourService(uc *biz.TourUsecase) *TourService {
	return &TourService{uc: uc}
}

func (s *TourService) SayHello(ctx context.Context, req *pb.HiRequest) (*pb.HiReply, error) {
	return &pb.HiReply{}, nil
}
func (s *TourService) Ai(ctx context.Context, req *pb.AiRequest) (*pb.AiReply, error) {
	resp, err := s.uc.GetAiResp(ctx, &biz.Tour{Msg: req.Msg})
	if err != nil {
		return nil, err
	}
	return &pb.AiReply{Msg: resp}, nil
}
