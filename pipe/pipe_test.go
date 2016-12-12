package pipe

import (
	"fmt"
	"testing"
)

func TestPipe(t *testing.T) {
	PipeDemo()
}

func BenchmarkPipe(b *testing.B) {
	fmt.Println("BeanchmarkPipe")
	PipeDemo()
}
