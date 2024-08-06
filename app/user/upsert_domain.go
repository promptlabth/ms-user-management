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
