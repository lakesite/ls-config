package config

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

var (
	cwd_arg = flag.String("cwd", "", "set cwd")
)

func init() {
	flag.Parse()
	if *cwd_arg != "" {
		if err := os.Chdir(*cwd_arg); err != nil {
			fmt.Println("Chdir error:", err)
		}
	}
}

func TestGetenv(t *testing.T) {
	// Test when environment is set
	if err := os.Setenv("LS-CONFIG_TEST", "set"); err != nil {
		t.Errorf("Error on os.Setenv: %s\n", err)
	}
	if Getenv("LS-CONFIG_TEST", "invalid") != "set" {
		t.Error("Getenv failed to retrieve environment variable.\n")
	}

	// Test fallback, and unset LS-CONFIG_TEST
	if err := os.Unsetenv("LS-CONFIG_TEST"); err != nil {
		t.Errorf("Error on os.Unsetenv: %s\n", err)
	}

	if Getenv("LS-CONFIG_TEST", "valid") != "valid" {
		t.Error("Getenv failed to retrieve fallback variable.\n")
	}
}

func TestGetWorkingDirectory(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Errorf("os.Getwd() returned error: %s\n", err)
	}
	if GetWorkingDirectory() != cwd {
		t.Error("GetWorkingDirectory() does not match os.Getwd()\n")
	}
}