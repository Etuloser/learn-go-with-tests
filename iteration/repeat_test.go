package iteration

import (
	"fmt"
	"testing"
)

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

func ExampleRepeat() {
	repeated := Repeat("a")
	fmt.Println(repeated)
	// Output: aaaaa
}
