package strev

import (
	"testing"
	"./strev"
)

type ReverseTest struct {
	in, out		string
}

var ReverseTests = []ReverseTest {
	ReverseTest{"abcd", "dcba"},
	ReverseTest{"HELLO-TW", "WT-OLLEH"},
}

func TestReverse(t *testing.T) {
	for _, r := range ReverseTests {
		exp := strev.Reverse(r.in)
		if r.out != exp {
			t.ErrorF("Reverse of %s expects %s, but got %s", r.in, r.out, exp)
		}
	}
}

func BenchmarkReverse(b *testing.B)  {
	s := "abcd"
	for i := 0; i < b.N; i++ {
		strev.Reverse(s)
	}
}
