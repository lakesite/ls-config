package main

import (
	"fmt"

	"github.com/lakesite/ls-config"
)

func main() {
	address := config.Getenv("APP_HOST", "127.0.0.1") + ":" + config.Getenv("APP_PORT", "7999")
	cwd := config.GetWorkingDirectory()
	fmt.Printf("Address: %s\n", address)
	fmt.Printf("CWD: %s\n", cwd)
}
