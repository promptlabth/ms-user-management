package user

type UserSerivce struct {
	userRepository userRepository
}

func NewService(userRepository userRepository) *UserSerivce {
	return &UserSerivce{
		userRepository: userRepository,
	}
}
