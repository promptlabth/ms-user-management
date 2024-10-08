package user

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

type CreateServiceSuite struct {
	suite.Suite
	ctrl *gomock.Controller

	userRepository *MockuserRepository

	svc *UserSerivce
}

func (t *CreateServiceSuite) SetupTest() {
	t.ctrl = gomock.NewController(t.T())
	t.userRepository = NewMockuserRepository(t.ctrl)
	t.svc = NewService(t.userRepository)
}

func (t *CreateServiceSuite) TearDownTest() {
	t.ctrl.Finish()
}

func TestCreateService(t *testing.T) {
	suite.Run(t, new(CreateServiceSuite))
}

func (t *CreateServiceSuite) Test_GetPlanFailed_ReturnError() {
	// Arrange
	msg := UpsertUserReqDomain{}
	t.userRepository.EXPECT().GetPlanByType(gomock.Any(), freePlanType).Return(nil, errors.New("error get plan failed"))

	// Act
	_, err := t.svc.CreateUser(context.Background(), msg)

	// Assert
	t.Error(err)
	t.EqualError(err, "error get plan failed")
}

func (t *CreateServiceSuite) Test_CreateUserFailed_ReturnError() {
	msg := UpsertUserReqDomain{}
	t.userRepository.EXPECT().GetPlanByType(gomock.Any(), freePlanType).Return(&PlanEntity{}, nil)

	t.userRepository.EXPECT().Transactional(gomock.Any()).DoAndReturn(
		func(fn func(tx *gorm.DB) error) error {
			t.userRepository.EXPECT().UpsertUser(
				gomock.Any(),
				gomock.Any(),
				gomock.Any(),
			).Return(errors.New("error create user"))
			return fn(nil)
		},
	)

	// Act
	_, err := t.svc.CreateUser(context.Background(), msg)

	// Assert
	t.Error(err)
	t.EqualError(err, "error create user")
}

func (t *CreateServiceSuite) Test_CreateUserBalanceIsFailed_ReturnError() {
	// Arrange
	msg := UpsertUserReqDomain{
		FirebaseId:  "firebase",
		Name:        "Name",
		Email:       TypeToPrt("prompt.lab@gmail.com"),
		ProfilePic:  TypeToPrt("www.img.url/profile"),
		Platform:    TypeToPrt("facebook"),
		AccessToken: TypeToPrt("accessToken"),
	}
	t.userRepository.EXPECT().GetPlanByType(gomock.Any(), freePlanType).Return(&PlanEntity{
		ID:       1,
		PlanType: "Free",
	}, nil)

	t.userRepository.EXPECT().Transactional(gomock.Any()).DoAndReturn(
		func(fn func(tx *gorm.DB) error) error {
			t.userRepository.EXPECT().UpsertUser(
				gomock.Any(),
				gomock.Any(),
				gomock.Eq(&UserEntity{
					FirebaseId:  "firebase",
					Name:        "Name",
					Email:       TypeToPrt("prompt.lab@gmail.com"),
					ProfilePic:  TypeToPrt("www.img.url/profile"),
					Platform:    TypeToPrt("facebook"),
					AccessToken: TypeToPrt("accessToken"),
					StripeId:    nil,
					PlanId:      1,
				}),
			).Return(nil)

			t.userRepository.EXPECT().CreateUserBalance(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(errors.New("error balance failed"))
			return fn(nil)
		},
	)

	// Act
	_, err := t.svc.CreateUser(context.Background(), msg)

	// Assert
	t.Error(err)
	t.EqualError(err, "error balance failed")
}

func (t *CreateServiceSuite) Test_CreateUserSuccessflow() {
	// Arrange
	msg := UpsertUserReqDomain{
		FirebaseId:  "firebase",
		Name:        "Name",
		Email:       TypeToPrt("prompt.lab@gmail.com"),
		ProfilePic:  TypeToPrt("www.img.url/profile"),
		Platform:    TypeToPrt("facebook"),
		AccessToken: TypeToPrt("accessToken"),
	}
	t.userRepository.EXPECT().GetPlanByType(gomock.Any(), freePlanType).Return(&PlanEntity{
		ID:       1,
		PlanType: "Free",
	}, nil)

	t.userRepository.EXPECT().Transactional(gomock.Any()).DoAndReturn(
		func(fn func(tx *gorm.DB) error) error {
			t.userRepository.EXPECT().UpsertUser(
				gomock.Any(),
				gomock.Any(),
				gomock.Eq(&UserEntity{
					FirebaseId:  "firebase",
					Name:        "Name",
					Email:       TypeToPrt("prompt.lab@gmail.com"),
					ProfilePic:  TypeToPrt("www.img.url/profile"),
					Platform:    TypeToPrt("facebook"),
					AccessToken: TypeToPrt("accessToken"),
					StripeId:    nil,
					PlanId:      1,
				}),
			).Return(nil)

			t.userRepository.EXPECT().CreateUserBalance(
				gomock.Any(),
				gomock.Any(),
				&UserBalanceMessage{
					FirebaseId:     "firebase",
					BalanceMessage: 0,
				},
			).
				Return(nil)
			return fn(nil)
		},
	)

	// t.userRepository.EXPECT().GetUserByFirebaseId(gomock.Any(), "firebase").Return(&UserEntity{
	// 	ID:          1,
	// 	FirebaseId:  "firebase",
	// 	PlanId:      1,
	// 	Name:        "Name",
	// 	Email:       TypeToPrt("prompt.lab@gmail.com"),
	// 	ProfilePic:  TypeToPrt("www.img.url/profile"),
	// 	Platform:    TypeToPrt("facebook"),
	// 	AccessToken: TypeToPrt("accessToken"),
	// }, nil)

	t.userRepository.EXPECT().GetPlanById(gomock.Any(), int64(1)).Return(&PlanEntity{
		PlanType:    "Free",
		MaxMessages: 60,
	}, nil)

	t.userRepository.EXPECT().GetBalanceMessage(gomock.Any(), "firebase").Return(
		&UserBalanceMessage{
			FirebaseId:     "firebase",
			BalanceMessage: 0,
		}, nil,
	)
	// Act
	res, err := t.svc.CreateUser(context.Background(), msg)

	// Assert
	t.NoError(err)
	t.Equal(&UpsertUserResponseDomain{
		UserDetail: UserDetailRespDomain{
			FirebaseId:  "firebase",
			Name:        "Name",
			Email:       TypeToPrt("prompt.lab@gmail.com"),
			ProfilePic:  TypeToPrt("www.img.url/profile"),
			Platform:    TypeToPrt("facebook"),
			AccessToken: TypeToPrt("accessToken"),
			Balance:     0,
		},
		PlanDetail: PlanDetailRespDomain{
			PlanType:   "Free",
			MaxMessage: 60,
		},
	}, res)
}
