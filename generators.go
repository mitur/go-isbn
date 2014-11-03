package isbn

import (
	"math/rand"
	"time"
)

var (
	rg *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func GenerateISBN10List(num int) []string {
	strLi := make([]string, num)
	for i := 0; i < num; i++ {
		strLi[i] = GenerateISBN10()
	}
	return strLi
}

func GenerateISBN10() string {
	bites := make([]byte, 10)
	var sum uint32 = 0
	var i uint32 = 0
	for i < 9 {
		r := rg.Uint32() % 10
		bites[i] = byte(r + 48)
		sum = sum + r*(i+1)
		i++
	}

	mod := sum % 11

	if mod == 10 {
		bites[9] = byte(88) // "X"
	} else {

		bites[9] = byte(mod + 48)
	}
	return string(bites)
}

// Generates a bunch of ISBN13
func GenerateISBN13List(num int) []string {
	strLst := make([]string, num)

	for i := 0; i < num; i++ {
		strLst[i] = generateISBN13()
	}

	return strLst
}

func GenerateISBN13() string {
	bites := make([]byte, 13)
	var sum uint32 = 0
	var i uint32 = 0
	for i < 12 {
		r := rg.Uint32() % 10
		bites[i] = byte(r + 48)
		if i%2 == 0 {
			sum = sum + r
		} else {
			sum = sum + r*3
		}
		i++
	}

	calcChk := (10 - (sum % 10)) % 10
	bites[12] = byte(calcChk + 48)
	return string(bites)
}
