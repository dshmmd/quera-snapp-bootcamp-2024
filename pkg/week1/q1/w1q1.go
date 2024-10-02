package w1q1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

// Location represents each input line, which gives us relative coordinates of a location.
type Location struct {
	Name         string
	Reference    string
	X, Y         int
	IsCalculated bool
}

// getMatches gets a string and a regex expression, then compiles the string and returns matches.
// It panics if the string does not match the given pattern.
func getMatches(str string, regularExpression string) []string {
	regexCompiler := regexp.MustCompile(regularExpression)
	matches := regexCompiler.FindStringSubmatch(str)
	return matches
}

// calculateCoordinates is a recursive function that calculates absolute coordinates of a location.
func calculateCoordinates(locationMap map[string]*Location, loc *Location) {
	if loc.IsCalculated {
		return
	}

	if loc.Name == "start" {
		loc.IsCalculated = true
		return
	}

	referenceLocation, exists := locationMap[loc.Reference]
	if !exists {
		return
	}
	calculateCoordinates(locationMap, referenceLocation)

	loc.X += referenceLocation.X
	loc.Y += referenceLocation.Y
	loc.IsCalculated = true
}

// extractCoordinates calculates coordinates of each location and prints them in the same order as input.
func extractCoordinates(locationMap map[string]*Location, order []string) string {
	for _, name := range order {
		loc := locationMap[name]
		if !loc.IsCalculated {
			calculateCoordinates(locationMap, loc)
		}
	}

	var answer string
	for _, name := range order {
		if name != "start" {
			loc := locationMap[name]
			answer += fmt.Sprintf("%s x=%d y=%d\n", loc.Name, loc.X, loc.Y)
		}
	}
	return answer
}

func Solve(src io.Reader) (answer string, err error) {
	scanner := bufio.NewScanner(src)

	defer func() {
		r := recover()
		if r != nil {
			err = r.(error)
		}
	}()

	locationMap := make(map[string]*Location)
	var order []string

	scanner.Scan()
	startCoordination := scanner.Text()
	startMatches := getMatches(startCoordination, "start x=(-?\\d+) y=(-?\\d+)")
	xStart, err := strconv.Atoi(startMatches[1])
	if err != nil {
		return
	}
	yStart, err := strconv.Atoi(startMatches[2])
	if err != nil {
		return
	}

	locationMap["start"] = &Location{Name: "start", X: xStart, Y: yStart, IsCalculated: true}
	order = append(order, "start")

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		newMatches := getMatches(line, "([a-z0-9]+) from ([a-z0-9]+) x=([-+]?\\d+) y=([-+]?\\d+)")

		newLocationName := newMatches[1]
		newLocationFrom := newMatches[2]

		xNewPoint, err := strconv.Atoi(newMatches[3])
		if err != nil {
			return "", err
		}
		yNewPoint, err := strconv.Atoi(newMatches[4])
		if err != nil {
			return "", err
		}

		locationMap[newLocationName] = &Location{
			Name:         newLocationName,
			Reference:    newLocationFrom,
			X:            xNewPoint,
			Y:            yNewPoint,
			IsCalculated: false,
		}
		order = append(order, newLocationName)
	}

	return extractCoordinates(locationMap, order), nil
}

func init() {
	// make sure "os" package is present for submission script
	_ = os.DevNull
}
