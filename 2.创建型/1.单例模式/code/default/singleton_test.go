package singleton_test

import (
	singleton "my-singleton/default"
	"testing"
)

// command: `go test -v -timeout 30s -run ^TestMain$ my-singleton/default`
func TestMain(t *testing.T) {
	for i := 0; i < 30; i++ {
		go singleton.GetInstance()
	}
}
