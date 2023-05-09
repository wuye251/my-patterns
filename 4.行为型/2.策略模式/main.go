package main

import "fmt"

type Install interface {
	Add()
}

type Strategy struct {
	install Install
}

func (str *Strategy) Install(install Install) {
	fmt.Println("基本配置生成")

	install.Add()
}

type Tag struct{}

func (t Tag) Add() {
	fmt.Println("tag add()")
}

type Trend struct{}

func (trend Trend) Add() {
	fmt.Println("trend add()")
}

func main() {
	str := Strategy{}

	tag := Tag{}
	str.Install(tag)


	trend := Trend{}
	str.Install(trend)
}
