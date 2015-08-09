package config

type Config struct {
	SourceFilename string                    `hcl:""`
	UpstreamStates map[string]*UpstreamState `hcl:"upstream"`
	Build          *Stage                    `hcl:"build"`
	Deploy         *Stage                    `hcl:"deploy"`
	Infrastructure *Stage                    `hcl:"infrastructure"`
}

type UpstreamState struct {
	SourceURL string `hcl:"source"`
}

type Stage struct {
	ConfigDir string            `hcl:"config_dir"`
	Variables map[string]string `hcl:"variables"`
}
