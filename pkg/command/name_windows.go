package command

import (
	"fmt"
	"regexp"
)

func getCommandNameRegexp(prefix string) *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf(`^%s-([a-z]*).(exe|com|bat|cmd)`, regexp.QuoteMeta(prefix)))
}
