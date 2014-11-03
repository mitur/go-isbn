package isbn

import (
	"testing"
)

var (
	isbn10Li = [...]string{
		"0471958697",
		"0 471 60695 2",
		"0-470-84525-2",
		"0-321-14653-0",
	}

	isbn13Li = [...]string{
		"9780470059029",
		"978 0 471 48648 0",
		"978-0596809485",
		"978-0-13-149505-0",
		"978-0-262-13472-9",
	}

	invalidStrings = [...]string{
		"2442b52a68991",
		"",
		"-",
		"978-0-262-134722-9",
		"978-0-13-149505-2",
	}
)

func TestISBN10(t *testing.T) {
	for _, val := range isbn10Li {
		err := Validate(val)
		if err != nil {
			t.Errorf("Caught error \"%v\" when validating isbn10 %s", err, val)
		}
	}
}

func TestISBN13(t *testing.T) {
	for _, val := range isbn13Li {
		err := Validate(val)
		if err != nil {
			t.Errorf("Caught error \"%v\" when validating isbn13 %s", err, val)
		}
	}
}

func TestInvalidStrings(t *testing.T) {
	for _, val := range invalidStrings {
		if IsValid(val) {
			t.Errorf("%v should not be a valid ISBN", val)
		}
	}
}

// Generates a bunch of ISBN10s
func TestGenerateISBN10(t *testing.T) {
	for _, val := range generateISBN10List(1000) {
		err := Validate(val)
		if err != nil {
			t.Errorf("Caught error \"%v\" when validating string [%s]", err, val)
		}
	}
}

func TestGenerateISBN13(t *testing.T) {
	for _, val := range generateISBN13List(1000) {
		err := Validate(val)
		if err != nil {
			t.Errorf("Caught error \"%v\" when validating string [%s]", err, val)
		}
	}
}
