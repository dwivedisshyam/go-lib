package slice

import (
	"fmt"
	"strings"
)

func Join(s any, delim string) string {
	return strings.Replace(strings.Trim(fmt.Sprint(s), "[]"), " ", delim, -1)
}
