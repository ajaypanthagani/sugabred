package types

type BrewPackageInfo struct {
	Formulae []BrewFormula `json:"formulae"`
}

type BrewFormula struct {
	Name     string       `json:"name"`
	Versions BrewVersions `json:"versions"`
}

type BrewVersions struct {
	Stable string `json:"stable"`
}

type BrewCaskInfo struct {
	Casks []BrewCaskFormula `json:"casks"`
}

type BrewCaskFormula struct {
	Token   string `json:"token"`
	Version string `json:"version"`
}

type BrewPackage struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type BrewCask struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}
