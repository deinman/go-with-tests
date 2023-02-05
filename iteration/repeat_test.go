package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("0 times produces empty string", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""
		assertCorrectMessage(t, repeated, expected)
	})
	t.Run("repeat a 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})
	t.Run("character count matches expected times", func(t *testing.T) {
		const count = 5
		repeated := Repeat("a", count)
		found := strings.Count(repeated, "a")
		if count != found {
			t.Errorf("expected %d but found %d", count, found)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 5)
	fmt.Println(result)
	// Output: aaaaa
}

func assertCorrectMessage(t testing.TB, expected, repeated string) {
	t.Helper()
	if expected != repeated {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
