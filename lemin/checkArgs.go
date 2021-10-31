package lemin

import (
	"os"
)

func CheckArgs(args []string) {
	if len(args) != 1 {
		os.Exit(0)
	}
}
