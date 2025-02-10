// revive:disable-next-line
package main

import (
	"fmt"
	"github/Asac2142/go-cli-app/internal/cmd"
	"github/Asac2142/go-cli-app/internal/file"
	"github/Asac2142/go-cli-app/internal/task"
	"os"
)

func main() {
	f := file.New[task.TContent]()
	err := cmd.HandleTrackerCLI(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
