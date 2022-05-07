package command

import (
	"Serpent/config"
	"fmt"
	"os"
)

func Args() {
	args := os.Args
	version := []string{"version", "-v", "-V"}
	switch args[1] {
	case version[0]:
		fmt.Println(config.Version)
	case version[1]:
		fmt.Println(config.Version)
	case version[2]:
		fmt.Println(config.Version)
	case "http":
		httpflag(args[2:])
	}
}
