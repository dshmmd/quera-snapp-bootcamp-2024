package w1q1

import (
	"github.com/dshmmd/quera-snapp-bootcamp-2024/internal/tester"
	"testing"
)

// TestSolve contains multiple test cases to validate the behavior of the Solve function.
func TestSolve(t *testing.T) {
	tester.RunOverTestDirectory(t, Solve)
}
