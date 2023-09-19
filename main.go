package main

import (
	"polen/delivery"
)

// @title           Laundry APps
// @version         1.0

// @BasePath  /api/v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

// @schemes http
func main() {
	delivery.NewServer().Run()
}
