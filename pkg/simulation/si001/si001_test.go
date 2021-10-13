package si001

import (
	"testing"
)

func Test_si001(t *testing.T) {
	err := Run()
	if err != nil {
		t.Fatal(err)
	}
}
