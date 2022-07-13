package architecture

import (
	"runtime"
	"testing"
)

func TestDependencies(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		t.Skip("Expected amd64 architecture, but got", runtime.GOARCH)
	}

	t.Fail()
}
