package user

import (
	"context"

	"gorm.io/gorm"
)

func (s *UserSerivce) CreateUser(ctx context.Context, createUser UpsertUserReqDomain) (*UpsertUserResponseDomain, error) {

	planFree, err := s.userRepository.GetPlanByType(ctx, freePlanType)
	if err != nil {
		return nil, err
	}

	var user UserEntity
	err = s.userRepository.Transactional(func(tx *gorm.DB) error {
		user = UserEntity{
			FirebaseId:  createUser.FirebaseId,
			Name:        createUser.Name,
			Email:       createUser.Email,
			ProfilePic:  createUser.ProfilePic,
			Platform:    createUser.Platform,
			AccessToken: createUser.AccessToken,
			StripeId:    nil,
			PlanId:      planFree.ID,
		}
		if err := s.userRepository.UpsertUser(ctx, tx, &user); err != nil {
			return err
		}

		userBalance := UserBalanceMessage{
			FirebaseId:     createUser.FirebaseId,
			BalanceMessage: 0,
		}
		if err := s.userRepository.CreateUserBalance(ctx, tx, &userBalance); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	// // get user detail from firebase id
	// user, err := s.userRepository.GetUserByFirebaseId(ctx, createUser.FirebaseId)
	// if err != nil {
	// 	return nil, err
	// }

	// get detail of user plan
	plan, err := s.userRepository.GetPlanById(ctx, user.PlanId)
	if err != nil {
		return nil, err
	}

	// get balance of user
	userBalance, err := s.userRepository.GetBalanceMessage(ctx, user.FirebaseId)
	if err != nil {
		return nil, err
	}

	return &UpsertUserResponseDomain{
		UserDetail: UserDetailRespDomain{
			FirebaseId:  user.FirebaseId,
			Name:        user.Name,
			Email:       user.Email,
			ProfilePic:  user.ProfilePic,
			Platform:    user.Platform,
			AccessToken: user.AccessToken,
			Balance:     userBalance.BalanceMessage,
		},
		PlanDetail: PlanDetailRespDomain{
			PlanType:   plan.PlanType,
			MaxMessage: plan.MaxMessages,
		},
	}, nil
}
