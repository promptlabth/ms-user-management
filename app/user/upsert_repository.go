package user

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *UserRepository) Transactional(fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}

func (r *UserRepository) UpsertUser(ctx context.Context, tx *gorm.DB, userEntity *UserEntity) error {
	if userEntity == nil {
		return nil
	}

	db := r.db
	if tx != nil {
		db = tx
	}

	if err := db.Debug().WithContext(ctx).Clauses(
		clause.OnConflict{
			Columns: []clause.Column{{Name: "firebase_id"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"name",
				"email",
				"profilepic",
				"access_token",
			}),
		},
		clause.Returning{
			Columns: []clause.Column{
				{Name: "firebase_id"},
				{Name: "name"},
				{Name: "email"},
				{Name: "profilepic"},
				{Name: "platform"},
				{Name: "access_token"},
				{Name: "stripe_id"},
				{Name: "plan_id"},
			},
		},
	).Create(userEntity).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetPlanByType(ctx context.Context, planType string) (*PlanEntity, error) {
	var planEntity PlanEntity
	if err := r.db.Where("plan_type = ?", planType).First(&planEntity).Error; err != nil {
		return nil, err
	}
	return &planEntity, nil
}

func (r *UserRepository) CreateUserBalance(ctx context.Context, tx *gorm.DB, userBalanceEntity *UserBalanceMessage) error {
	if userBalanceEntity == nil {
		return nil
	}

	db := r.db
	if tx != nil {
		db = tx
	}

	if err := db.WithContext(ctx).Clauses(
		clause.OnConflict{
			Columns: []clause.Column{{
				Name: "firebase_id",
			}},
			DoNothing: true,
		},
	).Create(userBalanceEntity).Error; err != nil {
		// if have error and error is not duplicate key
		return err
	}
	return nil
}

func (r *UserRepository) GetBalanceMessage(ctx context.Context, firebase string) (*UserBalanceMessage, error) {
	var userBalance UserBalanceMessage

	if err := r.db.WithContext(ctx).Where("firebase_id = ?", firebase).First(&userBalance).Error; err != nil {
		return nil, err
	}
	return &userBalance, nil
}

func (r *UserRepository) GetPlanById(ctx context.Context, planId int64) (*PlanEntity, error) {
	var planEntity PlanEntity
	if err := r.db.Where("id = ?", planId).First(&planEntity).Error; err != nil {
		return nil, err
	}
	return &planEntity, nil
}

func (r *UserRepository) GetUserByFirebaseId(ctx context.Context, firebaseId string) (*UserEntity, error) {
	var user UserEntity

	if err := r.db.WithContext(ctx).Where("firebase_id = ?", firebaseId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
