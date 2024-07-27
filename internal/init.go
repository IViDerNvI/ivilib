package internal

import "os"

func init() {
	os.Setenv("TIME_ZONE", "UTC")
}
