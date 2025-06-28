package collectors

//go:generate mockgen -source=shell.go -destination=mocks/shell.go

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/ajaypanthagani/sugabred/commands"
	"github.com/ajaypanthagani/sugabred/types"
)

type ShellCollector interface {
	CollectShell() (*types.ShellSnapshot, error)
}

type shellCollector struct {
	shellCommander commands.ShellCommander
	fileCommander  commands.FileCommander
}

func NewShellCollector(shellCommander commands.ShellCommander, fileCommander commands.FileCommander) ShellCollector {
	return &shellCollector{
		shellCommander: shellCommander,
		fileCommander:  fileCommander,
	}
}

var shellConfigFiles = []string{
	".zshrc", ".zprofile", ".zshenv", ".bashrc", ".bash_profile", ".profile",
}

func (s *shellCollector) CollectShell() (*types.ShellSnapshot, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("could not retrieve current shelluser: %w", err)
	}

	cmdArgs := []string{".", "-read", "/Users/" + usr.Username, "UserShell"}

	defaultShell, err := s.shellCommander.RunCommand("dscl", cmdArgs...)
	if err != nil {
		return nil, fmt.Errorf("could not read default shell: %w", err)
	}

	defaultShell = parseShellFromDSCL(defaultShell)
	activeShell := os.Getenv("SHELL")

	aliasesRaw, _ := s.shellCommander.RunCommand(activeShell, "-i", "-c", "alias")
	aliases := parseAliases(aliasesRaw)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("could not retrieve homedir: %w", err)
	}

	configMap := map[string]string{}
	for _, configFile := range shellConfigFiles {
		fullPath := filepath.Join(homeDir, configFile)
		if content, err := s.fileCommander.ReadFile(fullPath); err == nil {
			configMap[configFile] = content
		}
	}

	return &types.ShellSnapshot{
		DefaultShell: strings.TrimSpace(defaultShell),
		ActiveShell:  strings.TrimSpace(activeShell),
		ConfigFiles:  configMap,
		Aliases:      aliases,
	}, nil
}

func parseShellFromDSCL(output string) string {
	for _, line := range strings.Split(output, "\n") {
		if strings.HasPrefix(line, "UserShell:") {
			return strings.TrimSpace(strings.TrimPrefix(line, "UserShell:"))
		}
	}

	return ""
}

func parseAliases(output string) map[string]string {
	aliases := map[string]string{}
	for _, line := range strings.Split(output, "\n") {
		parts := strings.SplitN(line, "=", 2)

		if len(parts) == 2 {
			name := strings.TrimSpace(strings.TrimPrefix(parts[0], "alias "))
			val := strings.Trim(parts[1], "'")
			aliases[name] = val
		}
	}

	return aliases
}
