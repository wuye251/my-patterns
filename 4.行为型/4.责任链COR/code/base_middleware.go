package myCor

import "fmt"

type MiddleWare interface {
	Next(middleware MiddleWare)
	Handle(msg string)
}

type BaseMiddleWare struct {
	NextM MiddleWare
}

func (base *BaseMiddleWare) Next(middleware MiddleWare) {
	base.NextM = middleware
}

func (base BaseMiddleWare) Handle(msg string) {
	fmt.Printf("base middleware msg:%s", msg)
	if base.NextM != nil {
		base.NextM.Handle(msg)
	}
}
