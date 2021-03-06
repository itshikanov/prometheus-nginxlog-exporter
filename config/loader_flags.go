package config

// LoadConfigFromFlags fills a configuration object (passed as parameter) with
// values from command-line flags.
func LoadConfigFromFlags(config *Config, flags *StartupFlags) error {
	config.Listen = ListenConfig{
		Port:    flags.ListenPort,
		Address: "0.0.0.0",
	}
	config.Namespaces = []NamespaceConfig{
		{
			Format:      flags.Format,
			SourceFiles: flags.Filenames,
			Name:        flags.Namespace,
			HistogramStep: flags.HistogramStep,
		},
	}

	return nil
}
