package main

type Subject interface {
	Do() string
}

type RealSubject struct{}

func (RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	real RealSubject
}

func (p Proxy) Do() string {
	var res string
	res += "pre:"
	res += p.real.Do()
	res += ":after"
	return res
}

func main() {
	var sub Subject
	sub = Proxy{RealSubject{}}
	println(sub.Do())
}

