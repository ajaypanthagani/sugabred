package types

type Snapshot struct {
	Timestamp string            `yaml:"timestamp"`
	Arch      string            `yaml:"architecture"`
	OSVersion string            `yaml:"macos_version"`
	Homebrew  []BrewPackage     `yaml:"homebrew_packages"`
	Casks     []BrewCask        `yaml:"homebrew_casks"`
	EnvVars   map[string]string `yaml:"environment_variables"`
}
