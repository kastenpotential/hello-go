package main

import (
	"fmt"
	"os"

	"github.com/kastenpotential/hello-go/tplink/commands"
	tpoutput "github.com/kastenpotential/hello-go/tplink/outputs"
	_ "github.com/kastenpotential/mage-utils/gen/travis-ci"
)

func PrintHelp() {
	help := `Usage:
    -l	Query devices 
    -p	Poll data 
    -P	Poll data constantly`
	fmt.Println(help)
}

func GetOutput() (error, tpoutput.Output) {
	use_output := "brief"
	if len(os.Args) > 2 {
		use_output = os.Args[2]
	}
	return tpoutput.GetOutput(use_output)
}

func main() {
	var err error
	err, output := GetOutput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(os.Args) < 3 {
		err = commands.Query(output)
		if err != nil {
			os.Exit(1)
		}
		os.Exit(0)
	}
	switch os.Args[1] {
	case "-q":
		err = commands.Query(output)
	/*
		case "-p":
			os.Exit(tpcmds.PollDevices())
		case "-P":
			os.Exit(tpcmds.LoopPollDevices())
	*/
	default:
		PrintHelp()
		os.Exit(1)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
