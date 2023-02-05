package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat a 5 times", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func ExampleRepeat() {
	result := Repeat("a")
	fmt.Println(result)
	// Output: aaaaa
}

func assertCorrectMessage(t testing.TB, expected, repeated string) {
	t.Helper()
	if expected != repeated {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
