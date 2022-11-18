package word

import "fmt"
import "testing"
import "desktop/Ninja_level-13/hands-on-exercise-2/quote"

func TestingCount(t *testing.T){
	s:="Hello i am ready to rock"
	c:=Count(s)
	if c!=24{
		t.Errorf("Expected 24 got %d",c)
	}
}

func TestingUseCount(t *testing.T){
	s:="Hello hi Hello hi"
    m:=UseCount(s)
	for k,v :=range m{
		switch k{
		case "Hello":
			if v!=2{
				t.Errorf("Expected 2 got %d",v)
			}
		case "hi":
			if v!=2{
				t.Errorf("Expected 2 got %d",v)
			}
		}
	}
}


func ExampleCount(){
	fmt.Println(Count("Hello i am ready to rock"))
	// Output:
	// 6
}

func BenchmarkCount(b *testing.B){
	for i:=0;i<b.N;i++{
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B){
	for i:=0;i<b.N;i++{
		UseCount(quote.SunAlso)
	}
}

// Start with this code. Get the code ready to BET on (add benchmarks, examples, tests) however
// do not write an example for the func that returns a map; and only write a test for it as an extra
// challenge. Add documentation to the code. Run the following in this order:
// ● tests
// ● benchmarks
// ● coverage
// ● coverage shown in web browser
// ● examples shown in documentation in a web browser
