package user

import (
	context "context"

	"github.com/promptlabth/ms-user-management/logger"
	userProto "github.com/promptlabth/proto-lib/user"
)

func (s *UserSerivce) GetUserById(ctx context.Context, firebase_id string) (*userProto.GetUserByIdRes, error) {
	res, err := s.userRepository.GetDetailUserById(ctx, firebase_id)
	if err != nil {
		return nil, err
	}
	logger.Info(ctx, res.PlanType)
	return &userProto.GetUserByIdRes{
		Id:          res.ID,
		FirebaseId:  res.FirebaseId,
		Name:        res.Name,
		Email:       res.Email,
		ProfilePic:  res.ProfilePic,
		Platform:    res.Platform,
		AccessToken: res.AccessToken,
		StripeId:    res.StripeId,
		Plan: &userProto.PlanDetail{
			Id:          res.PlanId,
			PlanType:    res.PlanType,
			MaxMessages: res.MaxMessages,
			ProductId:   res.ProductId,
		},
		BalanceMessage: res.BalanceMessage,
	}, nil
}
