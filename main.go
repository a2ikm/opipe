package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lines := strings.Split(string(bytes.TrimRight(data, "\n")), "\n")

	args0 := os.Args[1:]

	args := append(args0, lines...)

	cmd := exec.Command("open", args...)
	out, _ := cmd.CombinedOutput()
	fmt.Println(string(out))
}
