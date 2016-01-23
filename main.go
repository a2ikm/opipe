package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	args := os.Args[1:]
	cmd := exec.Command("open", args...)
	out, _ := cmd.CombinedOutput()
	fmt.Println(string(out))
}
