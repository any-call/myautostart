//go:build !darwin
// +build !darwin

package myautostart

import (
	"strconv"
	"strings"
)

func quote(args []string) string {
	for i, v := range args {
		args[i] = strconv.Quote(v)
	}

	return strings.Join(args, " ")
}
