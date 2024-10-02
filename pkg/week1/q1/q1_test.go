package q1

import (
	"github.com/dshmmd/quera-snapp-bootcamp-2024/internal/tester"
	"testing"
)

// TestSolve contains multiple test cases to validate the behavior of the Solve function.
func TestSolve(t *testing.T) {
	tester.RunTester(t, Solve, []tester.TestCase{
		{
			Name:     "Test Case 1",
			Input:    "start x=1 y=1\nhospital from start x=-1 y=+2",
			Expected: "hospital x=0 y=3",
			WantErr:  false,
		},
		{
			Name:     "Test Case 2",
			Input:    "start x=1 y=1\noffice from hospital x=+4 y=-1\nhospital from start x=-1 y=+2",
			Expected: "office x=4 y=2\nhospital x=0 y=3",
			WantErr:  false,
		},
	})
}
