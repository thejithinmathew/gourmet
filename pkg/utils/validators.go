package utils

import (
	"fmt"
	"strings"
)

func ValidateName(name string) error {
	if strings.HasPrefix(name, "-") {
		return fmt.Errorf("invalid name provided")
	}
	return nil
}
