package magic8

import (
	"testing"
)

func findInList(got string, list []string) bool {
	for _, v := range list {
		if got == v {
			return true
		}
	}
	return false
}

func TestRandomAnswer(t *testing.T) {
	got := RandomAnswer()

	if !findInList(got, responses) {
		t.Errorf("Got: %q not in list", got)
	}
}
