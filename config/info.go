package config

import (
	"fmt"
)

var VersionInfo = "Serpent 1.05"

func Version() {
	fmt.Println(VersionInfo)
}
func Help() {
	//fmt.Fprintln(ss, "aaaa")
	fmt.Println(`
   _____                            _
  / ____|                          | |
 | (___   ___ _ __ _ __   ___ _ __ | |_
  \___ \ / _ | '__| '_ \ / _ | '_ \| __|
  ____) |  __| |  | |_) |  __| | | | |_
 |_____/ \___|_|  | .__/ \___|_| |_|\__|
                  | |
                  |_|
Options:

  version
        View version information
  http
        Use HTTP mode
  sniff
        Sniff port
  help
        View help information
`)
}
