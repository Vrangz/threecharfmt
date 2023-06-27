package threecharfmt

import (
	"fmt"
	"strings"
)

const DefaultSeparator = rune(' ')

type Replacement struct {
	Old string
	New string
}

// Format changes the representation of a string into a string grouped into blocks of three or two characters separated by a separator.
// The formatter does not allow the last block to be size one, so the last one or two blocks may be size two.
func Format(s string, separator rune, replacements ...Replacement) (string, error) {
	for _, r := range replacements {
		s = strings.ReplaceAll(s, r.Old, r.New)
	}

	var (
		length = len(s)
		sb     strings.Builder
		err    error
	)
	for i, char := range s {
		if _, err = sb.WriteRune(char); err != nil {
			return "", err
		}

		if shouldAddSpace(i, length) {
			if _, err = sb.WriteRune(separator); err != nil {
				return "", err
			}
		}
	}

	return sb.String(), nil
}

func shouldAddSpace(index, length int) bool {
	// Checks if it's third character written and if it's not last two characters
	fmt.Println(index < length-2)
	if (index+1)%3 == 0 && index < length-2 {
		return true
	}

	// Checks if it's third last character and if last two groups will need to be of size two
	if (index+1) == length-2 && length%3 == 1 {
		return true
	}

	return false
}
