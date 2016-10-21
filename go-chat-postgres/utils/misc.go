package utils

import (
	"fmt"
	"log"
)

func Must(err error, msg string) {
	if err != nil {
		log.Fatal(fmt.Sprintf("%s: %s", err.Error()))
	}
}

func BuildByteString(str string) string {
	bs := []byte(str)
	var byteStr string
	for _, b := range bs {
		byteStr = fmt.Sprintf("%s%d ", byteStr, b)
	}

	return byteStr
}

func BuildStringFromByte(b string) string {

	return ""
}
