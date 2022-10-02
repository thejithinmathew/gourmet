package utils

import (
	"fmt"
	"regexp"
)

func ValidateName(name string) error {
	match, err := regexp.MatchString("^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$", name)
	if err != nil {
		return fmt.Errorf("error parsing request")
	}
	if !match {
		return fmt.Errorf("request name not valid. provide a proper name")
	}
	return nil
}
