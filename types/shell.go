package types

type ShellSnapshot struct {
	DefaultShell string            `yaml:"default_shell"`
	ActiveShell  string            `yaml:"active_shell"`
	ConfigFiles  map[string]string `yaml:"config_files"` // filename -> content
	Aliases      map[string]string `yaml:"aliases"`
}
