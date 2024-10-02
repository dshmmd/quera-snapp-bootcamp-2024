package tester

import (
	"bytes"
	"github.com/dshmmd/quera-snapp-bootcamp-2024/internal/resolver"
	"testing"
)

type TestCase struct {
	Name     string
	Input    string
	Expected string
	WantErr  bool
}

func RunTester(t *testing.T, Solve resolver.Resolver, tests []TestCase) {
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			// Create a new reader from the test case input
			reader := bytes.NewReader([]byte(tc.Input))

			// Call the Solve function
			answer, err := Solve(reader)

			// Check if an error was expected and compare with the result
			if (err != nil) != tc.WantErr {
				t.Fatalf("Solve() error = %v, wantErr %v", err, tc.WantErr)
			}

			// check if last character of answer is NL, remove it
			if answer[len(answer)-1] == '\n' {
				answer = answer[:len(answer)-1]
			}

			// If no error is expected, check the output against the expected value
			if answer != tc.Expected {
				t.Errorf("Solve() = %v, want %v", answer, tc.Expected)
			}
		})
	}
}
