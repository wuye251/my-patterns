package once_test

import (
	"my-singleton/once"
	"testing"
)

// command: `go test -v -timeout 30s -run ^TestMain$ my-singleton/once`
func TestMain(t *testing.T) {
	for i := 0; i < 30; i++ {
		go once.GetInstance()
	}
}
