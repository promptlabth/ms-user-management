package user

import context "context"

func (r *UserRepository) GetDetailUserById(ctx context.Context, firebaseId string) (*UserDetailInquirey, error) {
	result := UserDetailInquirey{}

	if err := r.db.Select(`
		users.*,
		plans.id AS plan_id,
		plans."planType" AS plan_type,
		plans."maxMessages" AS max_messages,
		plans.product_id AS product_id,
		user_balance_messages.balance_message AS balance_message
	`).
		Table("users").
		Joins("LEFT JOIN plans ON users.plan_id = plans.id").
		Joins("LEFT JOIN user_balance_messages ON users.firebase_id = user_balance_messages.firebase_id").
		Where("users.firebase_id = ?", firebaseId).Find(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
