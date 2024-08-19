package user

import (
	context "context"

	"github.com/promptlabth/ms-user-management/logger"
	userProto "github.com/promptlabth/proto-lib/user"
)

type UserRegis struct {
	userService *UserSerivce
	userProto.UserServiceServer
}

func NewRegister(userService *UserSerivce) *UserRegis {
	return &UserRegis{
		userService: userService,
	}
}

func (r *UserRegis) GetDetailUser(ctx context.Context, req *userProto.GetUserByIdReq) (*userProto.GetUserByIdRes, error) {
	res, err := r.userService.GetUserById(ctx, req.FirebaseId)
	if err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}
	return res, nil
}

func (r *UserRegis) UpsertUser(ctx context.Context, req *userProto.UpsertUserReq) (*userProto.UpsertUserRes, error) {
	msg := UpsertUserReqDomain{
		FirebaseId:  req.FirebaseId,
		Name:        req.Name,
		Email:       req.Email,
		ProfilePic:  req.ProfilePic,
		Platform:    req.Platform,
		AccessToken: req.AccessToken,
	}
	res, err := r.userService.CreateUser(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &userProto.UpsertUserRes{
		Status:  "200",
		Message: "success to save",
		UserDetail: &userProto.UpsertUserDataRes{
			FirebaseId:     res.UserDetail.FirebaseId,
			Name:           res.UserDetail.Name,
			Email:          res.UserDetail.Email,
			ProfilePic:     res.UserDetail.ProfilePic,
			Platform:       res.UserDetail.Platform,
			AccessToken:    res.UserDetail.AccessToken,
			BalanceMessage: res.UserDetail.Balance,
			StripeId:       res.UserDetail.StripeId,
		},
		PlanDetail: &userProto.LoginPlanDetailRes{
			PlanType:    res.PlanDetail.PlanType,
			MaxMessages: res.PlanDetail.MaxMessage,
		},
	}, nil
}
