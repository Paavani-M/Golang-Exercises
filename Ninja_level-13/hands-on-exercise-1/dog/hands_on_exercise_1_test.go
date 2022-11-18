package dog

import (
	"fmt"
	"testing"
)

func TestExam(t *testing.T) {
	a := Years(1)
	if a != 7 {
		t.Errorf("Expected 7 got %d", a)
	}

	b := YearsTwo(1)
	if a != 7 {
		t.Errorf("Expected 7 got %d", b)
	}
}

func ExampleYears() {
	fmt.Println(Years(1))
	// Output:
	// 7
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(1))
	// Output:
	// 7
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(1)
	}
}

func BenchmarkYears1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(1)
	}
}

// Start with this code. Get the code ready to BET on the code (add benchmarks, examples, tests).
// Run the following in this order:
// ● tests        - go test ./...
// ● benchmarks   - go test -bench . ./...
// ● coverage     - go test -cover ./...
// ● coverage shown in web browser   go test -coverprofile c.out ./... , go tool cover -html=c.out
// ● examples shown in documentation in a web browser - godoc -http :8989
