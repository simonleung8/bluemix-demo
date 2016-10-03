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
