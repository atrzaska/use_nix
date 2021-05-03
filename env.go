package main

import (
	"os"
	"strings"
)

type Env map[string]string

func GetEnv() Env {
	env := make(Env)

	for _, kv := range os.Environ() {
		kv2 := strings.SplitN(kv, "=", 2)

		key := kv2[0]
		value := kv2[1]

		env[key] = value
	}

	return env
}

func (env Env) ToShell() string {
	e := make(ShellExport)

	for key, value := range env {
		e.Add(key, value)
	}

	return Shell.Export(e)
}

func (env Env) Diff(other Env) *EnvDiff {
	return BuildEnvDiff(env, other)
}
