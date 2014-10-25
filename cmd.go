package main

import (
	"fmt"
	"os/exec"
)

func resizeCmd(src, dest string, size int) error {
	args := []string{src, "-resize", fmt.Sprintf("%dx%d", size, size), dest}
	cmd := exec.Command("convert", args...)

	return cmd.Run()
}
