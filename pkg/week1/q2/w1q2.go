package w1q2

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
)

func isEmoticon(r rune) bool {
	return r >= 0x1F600 && r <= 0x1F64F
}

func shiftCharacters(message []rune) []rune {
	var shiftedMessage []rune

	for _, char := range message {
		shiftAmount := rune(char) % 5

		if unicode.IsLetter(char) {
			var baseChar rune
			if unicode.IsLower(char) {
				baseChar = 'a'
			} else {
				baseChar = 'A'
			}
			shiftedChar := baseChar + (char-baseChar+shiftAmount)%26
			shiftedMessage = append(shiftedMessage, shiftedChar)

		} else if isEmoticon(char) {
			baseEmoji := rune(0x1F600)
			shiftedEmoji := baseEmoji + (char-baseEmoji+shiftAmount)%80
			shiftedMessage = append(shiftedMessage, shiftedEmoji)

		} else {
			shiftedMessage = append(shiftedMessage, char)
		}
	}

	return shiftedMessage
}

func replaceDigits(message []rune) []rune {
	var result []rune
	for _, char := range message {
		if unicode.IsDigit(char) {
			replacedValue := []rune(strconv.Itoa(int(char) * len(message)))
			result = append(result, '(')
			result = append(result, replacedValue...)
			result = append(result, ')')
		} else {
			result = append(result, char)
		}
	}
	return result
}

func replaceSpaces(message []rune) []rune {
	var result []rune
	var prevWordLength int
	var isInGroup bool

	for _, r := range message {
		if r == ' ' {
			for i := 0; i < prevWordLength; i++ {
				result = append(result, '_')
			}
			prevWordLength = 0
		} else {
			result = append(result, r)
			if r != '_' {
				if r == '(' {
					isInGroup = true
				} else if r == ')' {
					prevWordLength++
					isInGroup = false
				} else if !isInGroup {
					prevWordLength++
				}
			}
		}
	}
	return result
}

func reverseMessage(message []rune) []rune {
	var result []rune
	var digits []rune
	var inGroup bool

	for _, r := range message {
		if r == '(' {
			inGroup = true
			digits = []rune{}
		} else if r == ')' {
			reverse(digits)
			result = append(result, digits...)
			digits = []rune{}
			inGroup = false
		} else if inGroup {
			digits = append(digits, r)
		} else {
			result = append(result, r)
		}
	}

	reverse(result)

	return result
}

func reverse(result []rune) {
	left, right := 0, len(result)-1
	for left < right {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}
}

func Solve(src io.Reader) (answer string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	scanner := bufio.NewScanner(src)
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	for i := 0; i < n; i++ {
		scanner.Scan()
		messageRunes := []rune(scanner.Text())

		messageRunes = shiftCharacters(messageRunes)
		messageRunes = replaceDigits(messageRunes)
		messageRunes = replaceSpaces(messageRunes)
		messageRunes = reverseMessage(messageRunes)
		answer += fmt.Sprintf("%s\n", string(messageRunes))
	}
	return
}

func init() {
	// make sure "os" package is present for submission script
	_ = os.DevNull
}
