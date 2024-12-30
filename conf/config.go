package conf

func New() Config {
	return Config{
		Address: "",
		Workdir: "",
		Debug: false,
		Version: "0.0.4",
		VersionFlag: false,
	}
}

type Config struct {
	// commander address
	Address string

	// workdir
	Workdir string

	// debug
	Debug bool

	// version
	Version string
	VersionFlag bool
}
