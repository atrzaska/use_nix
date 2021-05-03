package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func execCommand() string {
	var stdoutBuffer bytes.Buffer
	cmd := exec.Command("nix-shell", "--show-trace", "--run", "\"env\"")
	cmd.Stdout = &stdoutBuffer
	err := cmd.Run()
	check(err)

	return stdoutBuffer.String()
}

func main() {
	dataContent := execCommand()
	lines := strings.Split(dataContent, "\n")
	newEnv := make(Env)

	for _, line := range lines {
		kv := strings.SplitN(line, "=", 2)

		if len(kv) != 2 {
			continue
		}

		key := kv[0]
		value := kv[1]

		newEnv[key] = value
	}

	script := GetEnv().Diff(newEnv).ToShell(Bash)
	fmt.Println(script)
}
