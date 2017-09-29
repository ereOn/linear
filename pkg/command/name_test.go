package command

import (
	"runtime"
	"testing"
)

func TestGetCommandNameRegexp(t *testing.T) {
	r := getCommandNameRegexp("prefix")
	v := "prefix-hello-world"

	if runtime.GOOS == "windows" {
		v = "prefix-hello-world.exe"
	}

	matches := r.FindStringSubmatch(v)

	if len(matches) == 0 {
		t.Error("no matches")
	}

	value := matches[1]
	expected := "hello-world"

	if value != expected {
		t.Errorf("expected `%s`, got `%s`", expected, value)
	}
}
