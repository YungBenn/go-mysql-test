package main

import (
	"github.com/YungBenn/go-mysql-test/internal/app"
)

func main() {
	err := app.SetupApp()
	if err != nil {
		panic(err)
	}
}
