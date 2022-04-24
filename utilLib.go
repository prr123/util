package utilLib
//
//
// library that parses text for chars
//
// author prr
// created 12/4/2022
//
// copyright 2022 prr
//

import (
	"os"
	"fmt"
)

func IsAlpha(let byte)(res bool) {
// function that test whether byte is alpha
    res = false
    if (let >= 'a' && let <= 'z') || (let >= 'A' && let <= 'Z') { res = true}
    return res
}

func IsAlphaNumeric(let byte)(res bool) {
// function that test whether byte is aphanumeric
    res = false
    tbool := (let >= 'a' && let <= 'z') || (let >= 'A' && let <= 'Z')
    if tbool || (let >= '0' && let <= '9') { res = true }
    return res
}

func IsNumeric(let byte)(res bool) {
// function that test whether byte is aphanumeric
    res = false
    if (let >= '0') && (let <= '9') { res = true }
    return res
}

func IsWsp(let byte)(res bool) {
	res = false
	if let ==' ' { res = true}
	return res
}

func FatErr(fs string, msg string, err error) {
// function that displays a console error message and exits program
	if err != nil {
		fmt.Printf("error %s:: %s!%v\n", fs, msg, err)
	} else {
		fmt.Printf("error %s:: %s!\n", fs, msg)
	}
	os.Exit(2)
}

