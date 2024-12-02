package aocutils

import (
	"regexp"
	"strconv"
)

// int operations

func GetAllIntegersFromString(text string) (vals []int) {
	re := regexp.MustCompile("[0-9]+")
	strings := re.FindAllString(text, -1)
	for _, stringint := range strings {
		val, _ := strconv.Atoi(stringint)
		vals = append(vals, val)
	}
	return vals
}
