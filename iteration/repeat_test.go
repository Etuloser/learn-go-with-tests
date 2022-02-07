package iteration

import (
	"fmt"
	"testing"
)

const repeatConut = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatConut; i++ {
		repeated += character
	}
	return repeated
}
func ExampleRepeat() {
	repeated := Repeat("a")
	fmt.Println(repeated)
	// Output: aaaaa
}
func TestReapeat(t *testing.T) {
	repeat := Repeat("a")
	expected := "aaaaa"

	if repeat != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeat)
	}
}

// 基准测试
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
