package main

type Shell interface {
	Export(e ShellExport) string
}

type ShellExport map[string]*string

func (e ShellExport) Add(key, value string) {
	e[key] = &value
}

func (e ShellExport) Remove(key string) {
	e[key] = nil
}
