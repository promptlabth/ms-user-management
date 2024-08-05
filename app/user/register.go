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

func (r *UserRegis) CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	msg := CreateUserReqDomain{
		FirebaseId:  req.FirebaseId,
		Name:        req.Name,
		Email:       req.Email,
		ProfilePic:  req.ProfilePic,
		Platform:    req.Platform,
		AccessToken: req.AccessToken,
		StripeId:    req.StripeId,
	}
	err := r.userService.CreateUser(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &CreateUserRes{
		Status:  "200",
		Message: "success to save",
	}, nil
}

func (r *UserRegis) mustEmbedUnimplementedUserServiceServer() {}
