package lib

import (
	"testing"
)

func TestStr_replace(t *testing.T) {
	r := str_replace([]string{"a", "bb"}, []string{"11", "22"}, "haha,bbc")
	if r != "h11h11,22c" {
		t.Error(r)
	}
}
