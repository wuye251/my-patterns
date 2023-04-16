package myCor_test

import (
	myCor "my-cor"
	"testing"
)

func TestMyCor(t *testing.T) {
	base := myCor.BaseMiddleWare{}
	response := myCor.ResponseMiddleware{}
	base.NextM = &response
	response.NextM = &myCor.StatisticMiddleware{}

	base.Handle("测试")
}
