package utils

import "strconv"

func SToIP(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}


func BoolToYN(b bool) string {
	if b == true {
		return "Y"
	}
	return "N"
}