package service

import (
	"testing"
)

func TestService_SendVerificationCode(t *testing.T) {

	//var (
	//	ctx           = context.Background()
	//	generatedCode = "1234"
	//	phone         = "123123123"
	//	code          = domain.Code{
	//		Code:  generatedCode,
	//		Phone: phone,
	//	}
	//	sms = &domain.Sms{
	//		Phone:   phone,
	//		Message: generatedCode,
	//	}
	//)
	//
	//// Setup mocks.
	//stubCtrl := gomock.NewController(t)
	//defer stubCtrl.Finish()
	//
	//// Mock repositories
	//smsServiceStub := mocks.NewMockSmsService(stubCtrl)
	//redisRepositoryStub := mocks.NewMockCodeRedisRepository(stubCtrl)
	//
	//// mock
	//smsServiceStub.EXPECT().
	//	SendSms(ctx, sms).
	//	Return(nil).
	//	AnyTimes()
	//
	//redisRepositoryStub.EXPECT().
	//	Set(ctx, generatedCode, code, domain.RedisCodeDuration).
	//	Return(nil).
	//	AnyTimes()
	//
	//// service
	//service := newBasicService(
	//	smsServiceStub,
	//	redisRepositoryStub,
	//)
	//
	//// test cases
	//type arguments struct {
	//	phone string
	//}
	//
	//tests := []struct {
	//	name        string
	//	arguments   arguments
	//	expectError bool
	//}{
	//	{
	//		name: "Success: sms send successfully",
	//		arguments: arguments{
	//			phone: phone,
	//		},
	//		expectError: false,
	//	},
	//	{
	//		name: "Fail: no phone number",
	//		arguments: arguments{
	//			phone: "",
	//		},
	//		expectError: true,
	//	},
	//}
	//
	//for _, test := range tests {
	//	t.Run(test.name, func(t *testing.T) {
	//		args := test.arguments
	//
	//		err := service.SendVerificationCode(ctx, args.phone)
	//		if !test.expectError {
	//			if err != nil {
	//				t.Errorf("unexpected error: %s", err)
	//			}
	//		} else {
	//			if err == nil {
	//				t.Error("expected error but got nothing")
	//			}
	//		}
	//	})
	//}
}

func TestService_ValidateCode(t *testing.T) {

}

func TestService_SendCode(t *testing.T) {

}
