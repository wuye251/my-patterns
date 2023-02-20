package my_builder_test

import (
	my_builder "my-builder"
	"testing"
)

// command `go test -v -timeout 30s -run ^TestBuilder$ my-builder`
func TestBuilder(t *testing.T) {
	thinPerson := my_builder.NewPersonInstance("thin")
	direct := my_builder.NewDirector()
	direct.SetBuilder(thinPerson)
	direct.Build()

	fatPerson := my_builder.NewPersonInstance("fat")
	direct.SetBuilder(fatPerson)
	direct.Build()

}
