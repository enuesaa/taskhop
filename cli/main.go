package cli

func New() ICli {
	cli := Cli{
		Address: "",
		Workdir: ".",		
	}
	return &cli
}

type ICli interface {
	Launch() error
	IsCommander() bool
	GetAddress() string
	GetWorkdir() string
}
