package main

import "order-status-log-service/config"

func main() {
	config.NewWorker().Initialize()
}
