package conf

func New() Config {
	return Config{
		Address: "",
		Workdir: "",
		Version: "0.0.4",
		VersionFlag: false,
	}
}

type Config struct {
	// commander address
	Address string

	// workdir
	Workdir string

	// version
	Version string
	VersionFlag bool
}
