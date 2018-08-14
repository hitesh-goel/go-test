package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("run default with no no's", func(t *testing.T) {
		repeat := Repeat("a", 0)
		expected := "aaaaa"

		if repeat != expected {
			t.Errorf("expected '%s' but got '%s'", expected, repeat)
		}
	})

	t.Run("run default with 6 no's", func(t *testing.T) {
		repeat := Repeat("a", 6)
		expected := "aaaaaa"

		if repeat != expected {
			t.Errorf("expected '%s' but got '%s'", expected, repeat)
		}
	})
}

func ExampleRepeat() {
	rep := Repeat("a", 6)
	fmt.Println(rep)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 8)
	}
}
