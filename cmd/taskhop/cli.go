package main

import (
	"fmt"
	"net"
	"os"

	"github.com/enuesaa/taskhop/conf"
	flag "github.com/spf13/pflag"
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
	c.parse()

	if c.config.VersionFlag {
		fmt.Printf("%s\n", c.config.Version)
		os.Exit(0)
	}
	if c.config.HelpFlag {
		flag.Usage()
		os.Exit(0)
	}
	c.printBanner()

	return nil
}

func (c *Cli) parse() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "taskhop\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n")
		flag.PrintDefaults()
	}
	flag.StringVarP(&c.config.Workdir, "workdir", "w", ".", "workdir. Example: ./aaa")
	flag.BoolVar(&c.config.VersionFlag, "version", false, "Print version")
	flag.BoolVarP(&c.config.HelpFlag, "help", "h", false, "Print help message")
	flag.Parse()
}

func (c *Cli) printBanner() {
	addr, err := c.getLocalIpAddress()
	if err != nil {
		addr = "localhosr"
	}
	fmt.Printf("running!\n")
	fmt.Printf("┌─────────────────────────────────────────────────────────────────\n")
	fmt.Printf("│ To launch the agent:\n")
	fmt.Printf("│   taskhop-agent -c %s:3000\n", addr)
	fmt.Printf("└─────────────────────────────────────────────────────────────────\n")
}

// see https://stackoverflow.com/questions/23558425
func (c *Cli) getLocalIpAddress() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	addr := conn.LocalAddr().(*net.UDPAddr)

	return addr.IP.To4().String(), nil
}
