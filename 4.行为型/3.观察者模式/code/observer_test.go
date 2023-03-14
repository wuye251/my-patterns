package my_observer_test

import (
	my_observer "my-observer"
	"testing"
)

// go test -v -timeout 30s -run ^TestPublish$ my-observer
func TestPublish(t *testing.T) {
	obs11 := &my_observer.Observer1{Name: "obs11"}
	obs12 := &my_observer.Observer1{Name: "obs12"}

	obs21 := &my_observer.Observer2{Attr: "obs21"}
	obs22 := &my_observer.Observer2{Attr: "obs22"}

	pub := my_observer.NewPublisher()
	pub.Register(obs11)
	pub.Register(obs12)
	pub.Register(obs21)
	pub.Register(obs22)

	pub.Publish("new product is arriving")
	// Output:
	// obs1 obs11 update message is new product is arriving
	// obs1 obs12 update message is new product is arriving
	// obs1 obs21 update  messageg is new product is arriving
	// obs1 obs22 update  messageg is new product is arriving
}
