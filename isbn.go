package isbn

import (
	"fmt"
)

// " " = 32
// "-" = 45
// 48-57 is the range of 0-9

func Validate(isbn string) error {
	nums := make([]int, 13)
	var i = 0
	for index, val := range isbn {
		ival := int(val)
		// Bail out if i is out of bounds
		if i > 12 {
			return fmt.Errorf("Invalid ISBN \"%s\", must contain 10 or 13 digits\n", isbn)
		} else if ival < 58 && 47 < ival { // numeric
			nums[i] = ival - 48
			i++
		} else if index == 9 && ival == 88 || ival == 120 { //ISBN allows checkdigit X(10), 88="X" 120="x"
			nums[i] = 10
			i++
			break
		} else if ival == 32 || ival == 45 { // seperator, (" " or "-"), ignore
			// dont do anything, simple ignore
		} else {
			return fmt.Errorf("%v is not valid in an ISBN string", string(val))
		}
	}

	if i == 10 {
		return validate10(nums[:10])
	} else if i == 13 {
		return validate13(nums)
	} else {
		return fmt.Errorf("Invalid ISBN \"%s\", must contain 10 or 13 digits\n", isbn)
	}

}

func IsValid(isbn string) bool {
	err := Validate(isbn)
	if err != nil {
		return false
	} else {
		return true
	}

}

// The check digit for ISBN-10 is calculated by multiplying
// each digit by its position (i.e., 1 x 1st digit, 2 x 2nd
// digit, etc.), summing these products together and taking
// modulo 11 of the result (with 'X' being used if the result
// is 10).
func validate10(isbn []int) error {
	var sum = 0
	for i := 0; i < len(isbn)-1; i++ {
		sum = sum + isbn[i]*(i+1)
	}

	if sum%11 == isbn[9] {
		return nil
	} else {
		return fmt.Errorf("Invalid checksum")
	}

}

// The check digit for ISBN-13 is calculated by multiplying
// each digit alternately by 1 or 3 (i.e., 1 x 1st digit,
// 3 x 2nd digit, 1 x 3rd digit, 3 x 4th digit, etc.), summing
// these products together, taking modulo 10 of the result
// and subtracting this value from 10, and then taking the
// modulo 10 of the result again to produce a single digit.
func validate13(isbn []int) error {
	var sum = 0
	for i := 0; i < len(isbn)-1; i++ {
		if i%2 == 0 {
			sum = sum + isbn[i]
		} else {
			sum = sum + isbn[i]*3
		}
	}

	calcChk := (10 - (sum % 10)) % 10
	actualChk := isbn[len(isbn)-1]

	if calcChk == actualChk {
		return nil
	} else {
		return fmt.Errorf("Invalid checksum %d (should be %d) for ISBN %v",
			actualChk, calcChk, isbn)
	}
}
