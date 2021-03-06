package main

import (
	"runtime"
	"testing"
)

func TestExecute(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip()
	}

	if exit := execute("true"); exit != 0 {
		t.Fatal(exit)
	}
	if exit := execute("false"); exit != 1 {
		t.Fatal(exit)
	}
	if exit := execute("sh", "-c", "exit 42"); exit != 42 {
		t.Fatal(exit)
	}
}
