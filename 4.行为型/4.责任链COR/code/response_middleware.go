package myCor

import "fmt"

type ResponseMiddleware struct {
	NextM MiddleWare
}

func (response *ResponseMiddleware) Next(middleware MiddleWare) {
	response.NextM = middleware
}

func (response ResponseMiddleware) Handle(msg string) {
	fmt.Printf("response middleware msg:%s", msg)
	if response.NextM != nil {
		response.NextM.Handle(msg)
	}
}
