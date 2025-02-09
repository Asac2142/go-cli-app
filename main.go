// revive:disable-next-line
package main

import (
	"fmt"
	"github/Asac2142/go-cli-app/internal/cmd"
	"github/Asac2142/go-cli-app/internal/file"
	"github/Asac2142/go-cli-app/internal/task"
)

func main() {
	f := file.New[task.TContent]()
	err := cmd.HandleTrackerCLI(f)
	if err != nil {
		printError(err)
	}
}

func printError(err error) {
	if err != nil {
		fmt.Printf("Issue found: %v\n", err)
	}
}
