package greetings

import (
	"os"
)

var Name = getHostname()

func getHostname() string {
	name, err := os.Hostname()
	if err != nil {
		return "mock-name"
	}
	return name
}
