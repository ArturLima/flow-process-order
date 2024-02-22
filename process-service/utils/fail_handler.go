package utils

import "log"

func FailWithError(m string, err error) {
	log.Printf("[ERROR] %s: %s", m, err)
}
