package main

import (
	"strings"
)

var IgnoredKeys = map[string]bool{
	"DIRENV_CONFIG":   true,
	"DIRENV_BASH":     true,
	"DIRENV_IN_ENVRC": true,
	"COMP_WORDBREAKS": true, // Avoids segfaults in bash
	"PS1":             true, // PS1 should not be exported, fixes problem in bash
	"OLDPWD":          true,
	"PWD":             true,
	"SHELL":           true,
	"SHELLOPTS":       true,
	"SHLVL":           true,
	"_":               true,
}

type EnvDiff struct {
	Prev map[string]string `json:"p"`
	Next map[string]string `json:"n"`
}

func NewEnvDiff() *EnvDiff {
	return &EnvDiff{make(map[string]string), make(map[string]string)}
}

func BuildEnvDiff(e1, e2 Env) *EnvDiff {
	diff := NewEnvDiff()

	in := func(key string, e Env) bool {
		_, ok := e[key]
		return ok
	}

	for key := range e1 {
		if IgnoredEnv(key) {
			continue
		}
		if e2[key] != e1[key] || !in(key, e2) {
			diff.Prev[key] = e1[key]
		}
	}

	for key := range e2 {
		if IgnoredEnv(key) {
			continue
		}
		if e2[key] != e1[key] || !in(key, e1) {
			diff.Next[key] = e2[key]
		}
	}

	return diff
}

func (diff *EnvDiff) ToShell(shell Shell) string {
	e := make(ShellExport)

	for key := range diff.Prev {
		_, ok := diff.Next[key]
		if !ok {
			e.Remove(key)
		}
	}

	for key, value := range diff.Next {
		e.Add(key, value)
	}

	return shell.Export(e)
}

func IgnoredEnv(key string) bool {
	if strings.HasPrefix(key, "__fish") {
		return true
	}
	if strings.HasPrefix(key, "BASH_FUNC_") {
		return true
	}
	_, found := IgnoredKeys[key]
	return found
}
