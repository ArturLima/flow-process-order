package main

import (
	"receive-service/config"
)

func main() {
	config.NewProvider().Initialize()
}
