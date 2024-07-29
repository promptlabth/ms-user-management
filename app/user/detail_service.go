package user

import (
	context "context"

	"github.com/promptlabth/ms-user-management/logger"
)

func (s *UserSerivce) GetUserById(ctx context.Context, firebase_id string) (*GetUserByIdRes, error) {
	res, err := s.userRepository.GetDetailUserById(ctx, firebase_id)
	if err != nil {
		return nil, err
	}
	logger.Info(ctx, res.PlanType)
	return &GetUserByIdRes{
		Id:          res.ID,
		FirebaseId:  res.FirebaseId,
		Name:        res.Name,
		Email:       res.Email,
		ProfilePic:  res.ProfilePic,
		Platform:    res.Platform,
		AccessToken: res.AccessToken,
		StripeId:    res.StripeId,
		Plan: &PlanDetail{
			Id:          res.PlanId,
			PlanType:    res.PlanType,
			MaxMessages: res.MaxMessages,
			ProductId:   res.ProductId,
		},
		BalanceMessage: res.BalanceMessage,
	}, nil
}
