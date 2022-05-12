package command

import (
	"Serpent/config"
	"os"
)

func Args() {
	args := os.Args
	version := []string{"version", "-v", "-V"}
	help := []string{"help", "-h", "--help"}
	switch args[1] {
	case version[0]:
		config.Version()
	case version[1]:
		config.Version()
	case version[2]:
		config.Version()
	case "http":
		httpflag(args[2:])
	//case "blast":
	//blastflag(args[2:])
	case "sniff":
		sniffflag(args[2:])
	case help[0]:
		config.Help()
	case help[1]:
		config.Help()
	case help[2]:
		config.Help()
	}
}
