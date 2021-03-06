package main

import (
	"fmt"
)

type shell struct{}

var Shell shell

type ShellExport map[string]*string

func (e ShellExport) Add(key, value string) {
	e[key] = &value
}

func (e ShellExport) Remove(key string) {
	e[key] = nil
}

const (
	ACK           = 6
	TAB           = 9
	LF            = 10
	CR            = 13
	US            = 31
	SPACE         = 32
	AMPERSTAND    = 38
	SINGLE_QUOTE  = 39
	PLUS          = 43
	NINE          = 57
	QUESTION      = 63
	LOWERCASE_Z   = 90
	OPEN_BRACKET  = 91
	BACKSLASH     = 92
	UNDERSCORE    = 95
	CLOSE_BRACKET = 93
	BACKTICK      = 96
	TILDA         = 126
	DEL           = 127
)

func (sh shell) Export(e ShellExport) (out string) {
	for key, value := range e {
		if value == nil {
			out += sh.unset(key)
		} else {
			out += sh.export(key, *value)
		}
	}
	return out
}

func (sh shell) export(key, value string) string {
	return "export " + key + "=" + sh.escape(value) + "\n"
}

func (sh shell) unset(key string) string {
	return "unset " + sh.escape(key) + ";"
}

func (sh shell) escape(str string) string {
	if str == "" {
		return "''"
	}
	in := []byte(str)
	out := ""
	i := 0
	l := len(in)
	escape := false

	hex := func(char byte) {
		escape = true
		out += fmt.Sprintf("\\x%02x", char)
	}

	backslash := func(char byte) {
		escape = true
		out += string([]byte{BACKSLASH, char})
	}

	escaped := func(str string) {
		escape = true
		out += str
	}

	quoted := func(char byte) {
		escape = true
		out += string([]byte{char})
	}

	literal := func(char byte) {
		out += string([]byte{char})
	}

	for i < l {
		char := in[i]
		switch {
		case char == ACK:
			hex(char)
		case char == TAB:
			escaped(`\t`)
		case char == LF:
			escaped(`\n`)
		case char == CR:
			escaped(`\r`)
		case char <= US:
			hex(char)
		case char <= AMPERSTAND:
			quoted(char)
		case char == SINGLE_QUOTE:
			backslash(char)
		case char <= PLUS:
			quoted(char)
		case char <= NINE:
			literal(char)
		case char <= QUESTION:
			quoted(char)
		case char <= LOWERCASE_Z:
			literal(char)
		case char == OPEN_BRACKET:
			quoted(char)
		case char == BACKSLASH:
			backslash(char)
		case char == UNDERSCORE:
			literal(char)
		case char <= CLOSE_BRACKET:
			quoted(char)
		case char <= BACKTICK:
			quoted(char)
		case char <= TILDA:
			quoted(char)
		case char == DEL:
			hex(char)
		default:
			hex(char)
		}
		i++
	}

	if escape {
		out = "$'" + out + "'"
	}

	return out
}
