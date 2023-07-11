package cli

import (
	"fmt"
	"gtc/api"
	"gtc/utils"
	"os"
	"os/user"
	"strings"

	"github.com/logrusorgru/aurora/v4"
)

func GetHomeDir() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	return currentUser.HomeDir, nil
}

func GetFileName() (string, error) {
	homedir, err := GetHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	return homedir + "/desio-thoughts.csv", nil
}

func Run() {
	if len(os.Args) < 2 {
		utils.PrintHelp()
	}

	filepath, err := GetFileName()
	if err != nil {
		fmt.Println(err)
	}

	// TODO: add a delete flag
	switch os.Args[1] {
	case "-h", "-help": // print help
		utils.PrintHelp()

	case "-l", "-list": // list all thoughts
		err := api.ListAllThoughts(filepath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	case "-a", "-add": // add a thought
		// TODO: abstract this into a function

		thought := strings.Join(os.Args[2:], " ")

		err := api.AddThought(filepath, thought)
		if err != nil {
			fmt.Println(aurora.Red(err))
			os.Exit(1)
		}

	default: // add a thought
		err := api.ListAllThoughts(filepath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}
}
