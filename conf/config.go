package conf

func New() *Config {
	return &Config{
		Address:      "",
		Workdir:      "",
		TransferFlag: false,
		Version:      "0.1.0",
		VersionFlag:  false,
	}
}

type Config struct {
	// commander address
	Address string

	// workdir
	Workdir string

	// transfer
	TransferFlag bool

	// version
	Version     string
	VersionFlag bool

	// help
	HelpFlag bool
}
