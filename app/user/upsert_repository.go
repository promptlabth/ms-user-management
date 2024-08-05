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

	if err := db.WithContext(ctx).Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "firebase_id"}},
			UpdateAll: true,
		},
	).Create(userEntity).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetPlanByType(ctx context.Context, planType string) (*PlanEntity, error) {
	var planEntity PlanEntity
	if err := r.db.Raw(`SELECT
	id,
	'planType' as plan_type,
	'maxMessages' as mex_messages,
	product_id
	FROM plans
	WHERE plan_type = ?
	`, freePlanType).Scan(&planEntity).Error; err != nil {
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

	if err := db.WithContext(ctx).Create(userBalanceEntity).Error; err != nil && err != gorm.ErrDuplicatedKey {
		// if have error and error is not duplicate key
		return err
	}
	return nil
}
