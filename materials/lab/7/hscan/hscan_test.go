// Optional Todo

package hscan

import (
	"testing"
)

func TestGuessSingle(t *testing.T) {
	got := GuessSingle("2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae", "/home/cabox/workspace/course-materials/materials/lab/7/main/wordlist.txt")
	want := "foo"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

}
