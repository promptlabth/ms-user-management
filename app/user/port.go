//go:generate bash -c "mockgen -source=port.go -package=$(go list -f '{{.Name}}') -destination=mock_port_test.go"
package user

import (
	context "context"

	"gorm.io/gorm"
)

type UserServicer interface {
	GetUserById()
}

type userRepository interface {
	Transactional(fn func(tx *gorm.DB) error) error
	GetDetailUserById(ctx context.Context, id string) (*UserDetailInquirey, error)
	UpsertUser(ctx context.Context, tx *gorm.DB, userEntity *UserEntity) error
	CreateUserBalance(ctx context.Context, tx *gorm.DB, userBalanceEntity *UserBalanceMessage) error
	GetPlanByType(ctx context.Context, planType string) (*PlanEntity, error)
}
