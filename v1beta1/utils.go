package v1beta1

import (
	"strings"

	"github.com/samber/lo"
)

func IsArgsEqual(first []string, second []string) bool {
	if len(first) != len(second) {
		return false
	}

	for len(first) > 0 {
		s := first[0]

		_, i, _ := lo.FindIndexOf(second, func(str string) bool {
			return s == str
		})

		if i < 0 {
			return false
		}

		// Try to treats the two elements as a key-value pair
		if strings.HasPrefix(s, "-") {
			if len(first) == 1 {
				// The last element is a key, not a key-value pair
				return true
			}

			if strings.HasPrefix(first[1], "-") {
				// The next element is a key, not a value
				first = first[1:]
				second = append(second[:i], second[i+1:]...)
				continue
			}

			if first[1] != second[i+1] {
				return false
			}

			first = first[2:]
			second = append(second[:i], second[i+2:]...)
		} else {
			first = first[1:]
			second = append(second[:i], second[i+1:]...)
		}
	}

	return true
}
