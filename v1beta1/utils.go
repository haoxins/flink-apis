package v1beta1

import (
	"errors"

	"github.com/samber/lo"
)

func IsArgsEqual(a []string, b []string) (bool, error) {
	if len(a) != len(b) {
		return false, nil
	}

	if len(a) == 0 {
		return true, nil
	}

	if len(a)%2 != 0 {
		return false, errors.New("The args must be an even number")
	}

	for i := 0; i < len(a); i += 2 {
		_, j, _ := lo.FindIndexOf(b, func(s string) bool {
			return s == a[i]
		})

		if j < 0 {
			return false, nil
		}
		if a[i] != b[j] {
			return false, nil
		}
		if a[i+1] != b[j+1] {
			return false, nil
		}
	}

	return true, nil
}
