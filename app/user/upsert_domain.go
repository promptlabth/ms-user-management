package user

type UpsertUserReqDomain struct {
	FirebaseId  string
	Name        string
	Email       *string
	ProfilePic  *string
	Platform    *string
	AccessToken *string
	StripeId    *string
}

type UpsertUserResponseDomain struct {
	UserDetail UserDetailRespDomain
	PlanDetail PlanDetailRespDomain
}

type UserDetailRespDomain struct {
	FirebaseId  string
	Name        string
	Email       *string
	ProfilePic  *string
	Platform    *string
	AccessToken *string
	StripeId    *string
	Balance     int64
}

type PlanDetailRespDomain struct {
	PlanType   string
	MaxMessage int64
}
