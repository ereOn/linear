package command

import (
	"runtime"
	"testing"
)

func TestGetCommandNameMismatch(t *testing.T) {
	v := "prefix-hello_world"

	if runtime.GOOS == "windows" {
		v = "prefix-hello-world.exe"
	}

	value := GetCommandName("prefix", v)
	expected := ""

	if value != expected {
		t.Errorf("expected `%s`, got `%s`", expected, value)
	}
}

func TestGetCommandName(t *testing.T) {
	v := "prefix-hello-world"

	if runtime.GOOS == "windows" {
		v = "prefix-hello-world.exe"
	}

	value := GetCommandName("prefix", v)
	expected := "hello-world"

	if value != expected {
		t.Errorf("expected `%s`, got `%s`", expected, value)
	}
}
