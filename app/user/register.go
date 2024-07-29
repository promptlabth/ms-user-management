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

func (r *UserRegis) mustEmbedUnimplementedUserServiceServer() {}
