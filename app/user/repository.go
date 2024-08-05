package user

import "gorm.io/gorm"

type UserRepository struct {
	db *gorm.DB
}

// verify interface compliance
var _ userRepository = (*UserRepository)(nil)

func NewRepostiroy(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
