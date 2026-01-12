//go:build unix

package cmd

import (
	"io/fs"
	"syscall"

	"chezmoi.io/chezmoi/internal/chezmoi"
)

const defaultEditor = "vi"

var defaultInterpreters = map[string]chezmoi.Interpreter{
	"yaml": {
		Command: "ansible-playbook",
		Args:    []string{"-i", "localhost,", "-c", "local"},
	},
	"yml": {
		Command: "ansible-playbook",
		Args:    []string{"-i", "localhost,", "-c", "local"},
	},
}

func fileInfoUID(info fs.FileInfo) int {
	return int(info.Sys().(*syscall.Stat_t).Uid) //nolint:forcetypeassert
}

func windowsVersion() (map[string]any, error) {
	return nil, nil
}
