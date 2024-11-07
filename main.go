package main

import (
	"fmt"
	"strconv"
	"strings"
)


type PhoneNumber string

// Check if there is only one + symbol
func onlyOnePrefix(p string) (bool, int) {
	r := strings.Count(p, "+")
	if r == 1 {
		return true, 0
	} else {
		if r > 1 {
			return false, 1006
		} 
		return false, 1003
	}
}

// check if a string is made of digits
func onlyDigits(p string) (bool, int) {
	noprefix := strings.ReplaceAll(p, "+", "")
	nospaces := strings.ReplaceAll(noprefix, " ", "")
	_, err := strconv.Atoi(nospaces)
	if err == nil {
		return true, 0
	} else {
		return false, 1005
	}
}

// Check if it is a valid Swiss number
func isSuissePhone(p string) (bool, int) {
	trimmed := trimmedPhone(p)
	if !strings.HasPrefix(trimmed, "+41") {
		return false, 0
	} else {
		if len(trimmed) == 12 {
			return true, 0
		} else {
			if len(trimmed) > 12 {
				return true, 1001
			} else {
				return true, 1002
			}
		}
	}
}

// Check if it is a valid international number
func isInternationalPhone(p string) (bool, int) {
	trimmed := trimmedPhone(p)
	
	if strings.HasPrefix(trimmed, "+41") {
		return false, 0
	} else {
		if !strings.HasPrefix(trimmed, "+") {
			return false, 1003
		} else {
			if len(trimmed) < 11 {
				return true, 1002
			} else {
				if len(trimmed) > 15 {
					return true, 1001
				} else {
					return true, 0
				}
			}
		}
	} 
}

func trimmedPhone(p string) string {
	return strings.ReplaceAll(p, " ", "")
}


func NewPhoneNumber(p string) (string, int) {
	phoneNumber := trimmedPhone(p)
	isSwiss, err1 := isSuissePhone(phoneNumber)
	isInternational, err2 := isInternationalPhone(phoneNumber)
	containsOnlyOnePrefix, err3 := onlyOnePrefix(phoneNumber)
	containsOnlyDigits, err4 := onlyDigits(phoneNumber)

	if !containsOnlyOnePrefix {
		return "", err3
	} else {
		if !containsOnlyDigits {
			return "", err4
		} else {
			if isSwiss {
				return phoneNumber, err1
			} else {
				if isInternational {
					return phoneNumber, err2
				} else {
					return "", err2
				}
			}
		}
	}
}

func main() {
	number := "32328989555"
	fmt.Print("Number without prefix is valid : ")
	fmt.Println(onlyDigits(number))

	fmt.Print("Is Swiss phone ")
	fmt.Println(isSuissePhone(number))

	fmt.Print("Is International Phone ")
	fmt.Println(isInternationalPhone(number))

	fmt.Print(" Resultado, error : ")
	fmt.Println(NewPhoneNumber(number))

}
