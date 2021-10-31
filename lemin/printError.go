package lemin

import (
	"fmt"
	"os"
)

func PrintError() {
	fmt.Println("ERROR: invalid data format")
	os.Exit(0)
}
