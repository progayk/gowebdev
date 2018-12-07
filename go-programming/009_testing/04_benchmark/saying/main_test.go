package saying_test

import (
	"fmt"
	"github.com/progayk/gowebdev/go-programming/009_testing/04_benchmark/saying"
	"testing"
)

// Run:
// go test -bench .

func TestGreet(t *testing.T) {
	s := saying.Greet("Mayk")
	if s != "Welcome, Mayk" {
		t.Error("got", s, "want", "Welcome, Mayk")
	}
}

func ExampleGreet() {
	fmt.Println(saying.Greet("Mayk"))
	// Output:
	// Welcome, Mayk
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		saying.Greet("Mayk")
	}
}
