package utils

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
)

func PrintHelp() {
	fmt.Println("")
	fmt.Println(aurora.Green("Welcome to your thoughts!"))
	fmt.Println("")
	fmt.Println(aurora.Cyan("-a, -add: Add a thought: t -a <thought>"))
	fmt.Println(aurora.Cyan("-l, -list: List all thoughts: t -l"))
	fmt.Println(aurora.Cyan("-d, -delete: Delete a thought: t -d <number>"))

	fmt.Println("")
	os.Exit(0)
}
