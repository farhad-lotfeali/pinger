package main

import (
	"flag"
	"fmt"

	"github.com/farhad-lotfeali/pinger/pinger"
	"github.com/farhad-lotfeali/pinger/rest"
)

func main() {
	pingerCmd := flag.Bool("pinger", false, "run pinger in backgroun")
	restCmd := flag.Bool("rest", false, "run rest server http")
	flag.Parse()

	if *pingerCmd {
		pinger.RunPinger()
	} else if *restCmd {
		rest.RunHttpRest()
	} else {
		fmt.Println("you should use pinger or rest flag for run pinger")
	}
}
