package valiwork

import (
	// stdlib
	"os"
	"strconv"
)

var (
	DEBUG bool
)

// Initializes debug output.
// nolint
func init() {
	debug, found := os.LookupEnv("VALIWORK_DEBUG")
	if found {
		debugBool, err := strconv.ParseBool(debug)
		if err != nil {
			return
		}

		DEBUG = debugBool
	}
}
