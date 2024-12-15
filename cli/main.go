package cli

func New() ICli {
	cli := Cli{
		Address: "",
		Workdir: ".",
	}
	return &cli
}
