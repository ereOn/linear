package main

import "regexp"

// This might be improved by reading PATHEXT.
var subcommandBinaryPattern = regexp.MustCompile(`^linear-(.*).(exe|com|bat|cmd)`)
