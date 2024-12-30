package conf

import "os"

func New() Config {
	isDebugMode := os.Getenv("TASKHOP_DEBUG") == "true"

	return Config{
		Address: "",
		Workdir: "",
		Debug: isDebugMode,
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
