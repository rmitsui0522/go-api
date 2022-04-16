package main

import (
	"os"
)

const defaultPort = "3000"

func port() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":" + defaultPort
}
