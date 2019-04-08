package utils

import "strings"

func GenerUUId() string {

	return strings.Replace(NewV4().String(), "-", "", -1)
}


func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}