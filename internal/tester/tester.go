package tester

import (
	"bufio"
	"fmt"
	"github.com/dshmmd/quera-snapp-bootcamp-2024/internal/resolver"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

const (
	InputFormat  = "input%d.txt"
	OutputFormat = "output%d.txt"
)

type TestCase struct {
	Name     string
	Input    io.Reader
	Expected string
	WantErr  bool
}

func RunTester(t *testing.T, Solve resolver.Resolver, tests []TestCase) {
	log.Println("SLAAMAMAMMMAMMAMMAMMMMAMAMMA")
	log.Println(len(tests))
	for _, tc := range tests {
		log.Println(tc.Input)
		t.Run(tc.Name, func(t *testing.T) {
			// Call the Solve function
			answer, err := Solve(tc.Input)

			// Check if an error was expected and compare with the result
			if (err != nil) != tc.WantErr {
				t.Fatalf("Solve() error = %v, wantErr %v", err, tc.WantErr)
			}

			// check if last character of answer is NL, remove it
			if answer[len(answer)-1] == '\n' {
				answer = answer[:len(answer)-1]
			}

			if tc.Expected[len(answer)-1] == '\n' {
				tc.Expected = tc.Expected[:len(answer)-1]
			}

			// If no error is expected, check the output against the expected value
			if answer != tc.Expected {
				t.Errorf("Solve() = %v, want %v", answer, tc.Expected)
			}
		})
	}
}

func RunOverTestDirectory(t *testing.T, Solve resolver.Resolver) {
	_, testCallerPath, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("Failed to get the test file path")
	}
	testDir := filepath.Join(testCallerPath, "../testcases")

	testCases := make([]TestCase, 0)
	err := filepath.Walk(testDir, func(path string, info os.FileInfo, err error) error {
		if path == testDir {
			return nil
		}

		if info.IsDir() {
			t.Fatalf("found directory (all files should be .txt): %s", path)
		}

		var num int
		if _, err := fmt.Sscanf(info.Name(), InputFormat, &num); err == nil {
			expectedOutputFile := filepath.Join(testDir, fmt.Sprintf(OutputFormat, num))
			if _, err := os.Stat(expectedOutputFile); err != nil {
				t.Fatalf("failed to find expected output for %s (where looking for: %s)", info.Name(), expectedOutputFile)
			}

			inputFile, err := os.Open(path)
			if err != nil {
				t.Fatalf("failed to open file %s for reading: %s", path, err)
			}

			outputFile, err := os.Open(expectedOutputFile)
			if err != nil {
				t.Fatalf("failed to open file %s for reading: %s", expectedOutputFile, err)
			}

			expectedOutputString, err := io.ReadAll(outputFile)
			if err != nil {
				t.Fatalf("failed to read expected output file %s: %s", expectedOutputFile, err)
			}

			testCases = append(testCases, TestCase{
				Name:     fmt.Sprintf("TestCase #%d", num),
				WantErr:  false,
				Input:    bufio.NewReader(inputFile),
				Expected: string(expectedOutputString),
			})
		} else if _, err := fmt.Sscanf(info.Name(), OutputFormat, &num); err == nil {
			expectedInputFile := filepath.Join(testDir, fmt.Sprintf(InputFormat, num))
			if _, err := os.Stat(expectedInputFile); err != nil {
				t.Fatalf("failed to find expected input for %s (where looking for: %s)", info.Name(), expectedInputFile)
			}
		} else {
			t.Fatalf("file \"%s\" does not follow format \"input*.txt\" nor \"output*.txt\"", path)
		}
		return nil
	})

	if err != nil {
		t.Fatal(fmt.Errorf("failed to walk over test directory (%s): %w", testDir, err))
	}
	RunTester(t, Solve, testCases)
}
