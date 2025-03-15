package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
	"github.com/enuesaa/taskhop/conf"
)

func New(config *conf.Config) ICli {
	cli := Cli{
		config: config,
	}
	return &cli
}

type ICli interface {
	Launch() error
}

type Cli struct {
	config *conf.Config
}

func (c *Cli) Launch() error {
	flag.Usage = func ()  {
		fmt.Fprintf(os.Stderr, "taskhop\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n")
		flag.PrintDefaults()
	}

	c.parse()

	if c.config.VersionFlag {
		fmt.Printf("%s\n", c.config.Version)
		os.Exit(0)
	}
	if c.config.HelpFlag {
		flag.Usage()
		os.Exit(0)
	}
	return nil
}

func (c *Cli) parse() {
	flag.BoolVarP(&c.config.TransferFlag, "transfer", "t", false, "Transfer assets in workdir to the runner")
	flag.StringVar(&c.config.Workdir, "w", ".", "workdir. Example: ./aaa")
	flag.BoolVar(&c.config.VersionFlag, "version", false, "Print version")
	flag.BoolVarP(&c.config.HelpFlag, "help", "h", false, "Print help message")
	flag.Parse()
}
