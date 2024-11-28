package cli

import (
	"flag"
)

func Launch() Config {
	config := Config{
		Address: "",
		Workdir: "",
	}
	flag.StringVar(&config.Address, "c", "", "commander address. Example: localhost:3000")
	flag.StringVar(&config.Workdir, "w", ".", "workdir. Example: ./aaa")
	flag.Parse()

	return config
}
