package integers

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 2)
	want := "aa"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 1)
	}
}

func ExampleRepeat() {
	result := Repeat("r", 5)
	fmt.Println(result)
	// Output: rrrrr
}
