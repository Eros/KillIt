package main

import (
	_ "os/exec"
	"os/exec"
)

func checkExe(path string) error {
	_, err := exec.LookPath(path)

	if err != nil {
		return err;
	}
	return err;
}
