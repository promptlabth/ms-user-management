package user

type UserEntity struct {
	ID          int64   `gorm:"column:id;primaryKey"`
	FirebaseId  string  `gorm:"column:firebase_id;uniqueIndex"`
	Name        string  `gorm:"column:name"`
	Email       *string `gorm:"column:email"`
	ProfilePic  *string `gorm:"column:profilepic"`
	Platform    *string `gorm:"column:platform"`
	AccessToken *string `gorm:"column:access_token"`
	StripeId    *string `gorm:"column:stripe_id"`
	PlanId      int64   `gorm:"column:plan_id"`
}

func (UserEntity) TableName() string {
	return "users"
}

type PlanEntity struct {
	ID          int64   `gorm:"column:id;primaryKey"`
	PlanType    string  `gorm:"column:plan_type"`
	MaxMessages int64   `gorm:"column:max_messages"`
	ProductId   *string `gorm:"column:product_id"`
}

func (PlanEntity) TableName() string {
	return "plans"
}

type UserBalanceMessage struct {
	FirebaseId     string `gorm:"column:firebase_id;primaryKey"`
	BalanceMessage int64  `gorm:"column:balance_message"`
}

func (UserBalanceMessage) TableName() string {
	return "user_balance_messages"
}
