package main

import (
	"fmt"

	"github.com/thatisuday/commando"
)

func main() {
	//configure commando
	commando.
		SetExecutableName("golist").
		SetVersion("1.0.0").
		SetDescription("this tool looks for nested data in directory")

	//configuring root command
	commando.
		Register(nil).
		AddArgument("dir", "local dir path", "./").                      //"default ./"
		AddFlag("level,l", "level of depth of travel", commando.Int, 1). // default 1
		AddFlag("size", "display size of each dir", commando.Bool, nil). //default false
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Println("printing options for root command ...\n\n")

			for k, v := range args {
				fmt.Printf("args -> %v: %v(%T)\n", k, v.Value, v.Value)
			}

			for k, v := range flags {
				fmt.Printf("args -> %v: %v(%T)\n", k, v.Value, v.Value)
			}

		})

	//configuring the info command
	commando.
		Register("info").
		SetShortDescription("displays detailed information of a directory").
		SetDescription("This command displays more information about the contents of the directory like size, permission and ownership, etc.").
		AddArgument("dir", "local directory path", "./").                  // default `./`
		AddFlag("level,l", "level of depth to travel", commando.Int, nil). // required
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Printf("Printing options of the `info` command...\n\n")

			// print arguments
			for k, v := range args {
				fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			}

			// print flags
			for k, v := range flags {
				fmt.Printf("flag -> %v: %v(%T)\n", k, v.Value, v.Value)
			}
		})

	commando.Parse(nil)

}
