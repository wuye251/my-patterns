package myCor

import "fmt"

type StatisticMiddleware struct {
	NextM MiddleWare
}

func (static *StatisticMiddleware) Next(middleware MiddleWare) {
	static.NextM = middleware
}

func (static StatisticMiddleware) Handle(msg string) {
	fmt.Printf("static middleware msg:%s", msg)
	if static.NextM != nil {
		static.NextM.Handle(msg)
	}
}
