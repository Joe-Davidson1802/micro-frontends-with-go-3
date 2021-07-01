package main

import (
	"hello/app"
	"os"
)

func main() {
	app.Run(os.Getenv("ENV"))
}
