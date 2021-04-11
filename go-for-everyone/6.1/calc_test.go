package calc

import (
	"testing"
	"fmt"
)

func TestSum(t *testing.T) {
	if sum(1, 2) != 3{
		t.Fatal("sum(1,2) should be 3, but doesn't match")
	}
}

func TestSumFatal(t *testing.T) {
	if sum(1, 3) != 3{
		t.Fatal("sum(1,2) should be 3, but doesn't match")
	}
}

func ExampleHello() {
    fmt.Println("Hell")
   // Output: Hello
}