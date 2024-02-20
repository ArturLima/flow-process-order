package main

import "process-service/config"

func main() {
	config.NewProvider().Initialize()
}
