package connection

import (
	"os"
)

const defaultPort = "3000"

func Port() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":" + defaultPort
}
