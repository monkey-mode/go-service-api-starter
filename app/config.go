package app

type (
	AppConfig struct {
		Configs Configs `mapstructure:"configs"`
	}

	Configs struct {
	}
)
