package twofer

import (
    "strings"
	"fmt"
)

func ShareWith(name string) string {
	if len(strings.TrimSpace(name)) == 0 {
        name = "you"
    }
    return fmt.Sprintf("One for %s, one for me.", name)
}
