package user

import (
	context "context"

	"github.com/promptlabth/ms-user-management/logger"
)

type UserRegis struct {
	userService *UserSerivce
}

func NewRegister(userService *UserSerivce) *UserRegis {
	return &UserRegis{
		userService: userService,
	}
}

func (r *UserRegis) GetDetailUser(ctx context.Context, req *GetUserByIdReq) (*GetUserByIdRes, error) {
	res, err := r.userService.GetUserById(ctx, req.FirebaseId)
	if err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}
	return res, nil
}

func (r *UserRegis) UpsertUser(ctx context.Context, req *UpsertUserReq) (*UpsertUserRes, error) {
	msg := UpsertUserReqDomain{
		FirebaseId:  req.FirebaseId,
		Name:        req.Name,
		Email:       req.Email,
		ProfilePic:  req.ProfilePic,
		Platform:    req.Platform,
		AccessToken: req.AccessToken,
		StripeId:    req.StripeId,
	}
	res, err := r.userService.CreateUser(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &UpsertUserRes{
		Status:  "200",
		Message: "success to save",
		UserDetail: &UpsertUserDataRes{
			FirebaseId:     res.UserDetail.FirebaseId,
			Name:           res.UserDetail.Name,
			Email:          res.UserDetail.Email,
			ProfilePic:     res.UserDetail.ProfilePic,
			Platform:       res.UserDetail.Platform,
			AccessToken:    res.UserDetail.AccessToken,
			StripeId:       res.UserDetail.StripeId,
			BalanceMessage: res.UserDetail.Balance,
		},
		PlanDetail: &LoginPlanDetailRes{
			PlanType:    res.PlanDetail.PlanType,
			MaxMessages: res.PlanDetail.MaxMessage,
		},
	}, nil
}

func (r *UserRegis) mustEmbedUnimplementedUserServiceServer() {}
