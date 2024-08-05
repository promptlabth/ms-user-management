package user

import (
	"context"

	"gorm.io/gorm"
)

func (s *UserSerivce) CreateUser(ctx context.Context, createUser CreateUserReqDomain) error {
	planFree, err := s.userRepository.GetPlanByType(ctx, freePlanType)
	if err != nil {
		return err
	}

	err = s.userRepository.Transactional(func(tx *gorm.DB) error {
		user := UserEntity{
			FirebaseId:  createUser.FirebaseId,
			Name:        createUser.Name,
			Email:       createUser.Email,
			ProfilePic:  createUser.ProfilePic,
			Platform:    createUser.Platform,
			AccessToken: createUser.AccessToken,
			StripeId:    createUser.StripeId,
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
		return err
	}
	return nil
}
