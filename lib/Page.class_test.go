package lib

import (
	"testing"
)

func TestStrReplace(t *testing.T) {
	r := strReplace([]string{"a", "bb"}, []string{"11", "22"}, "haha,bbc")
	if r != "h11h11,22c" {
		t.Error(r)
	}
}
