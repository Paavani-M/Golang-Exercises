package mymath

import (
	"fmt"
	"testing"
)

func TestingCentered(t *testing.T) {
	type qna struct {
		data   []int
		answer float64
	}

	a := []qna{
		qna{[]int{3, 7, 8, 1, 2}, 4.2},
		qna{[]int{7, 1, 9, 2, 10}, 5.8},
		qna{[]int{5, 10, 11, 12, 19}, 11.4},
	}

	for _, v := range a {
		b := CenteredAvg(v.data)
		if b != v.answer {
			t.Error("Expected", v.answer, "got", b)
		}
	}

}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{5, 1, 3, 2, 4}))
	// Output:
	// 3
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{5, 19, 12, 34, 14})
	}
}

// Start with this code. Get the code ready to BET on (add benchmarks, examples, tests). Write a
// table test. Add documentation to the code. Run the following in this order:
// ● tests
// ● benchmarks
// ● coverage
// ● coverage shown in web browser
// ● examples shown in documentation in a web browser
