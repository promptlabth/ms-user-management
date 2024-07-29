package user

import context "context"

type UserServicer interface {
	GetUserById()
}

type userRepository interface {
	GetDetailUserById(ctx context.Context, id string) (*UserDetailInquirey, error)
}
