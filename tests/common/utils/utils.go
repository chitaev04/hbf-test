package utils

import (
	"context"
	"time"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LocalCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

// AssertGRPCError Проверка кода ошибки gRPC и вхождения ожидаемых подстрок в сообщение
func AssertGRPCError(sCtx provider.StepCtx, err error, expectedCode codes.Code, expectedSubstrings ...string) {
	sCtx.Require().Error(err)
	st, ok := status.FromError(err)
	sCtx.Require().True(ok)
	sCtx.Assert().Equal(expectedCode, st.Code())
	for _, s := range expectedSubstrings {
		sCtx.Assert().Contains(st.Message(), s)
	}
}
