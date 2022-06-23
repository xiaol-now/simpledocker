package container

type Option func(*ProcessConfig)

type ProcessConfig struct {
	Name string
	Envs []string
	TTY  bool
}

func SetContainerName(name string) Option {
	return func(config *ProcessConfig) {
		config.Name = name
	}
}

func SetEnvs(envs []string) Option {
	return func(config *ProcessConfig) {
		config.Envs = envs
	}
}

func SetTTY(tty bool) Option {
	return func(config *ProcessConfig) {
		config.TTY = tty
	}
}
