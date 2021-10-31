package lemin

import (
	"strings"
)

func ReadyFile(file []byte) []string {

	str := string(file)
	temp := strings.Split(str, "\n")
	return temp
}
