// revive:disable-next-line
package main

import (
	"fmt"
	"os"

	"github.com/Asac2142/go-cli-app/internal/cmd"
	"github.com/Asac2142/go-cli-app/internal/file"
	"github.com/Asac2142/go-cli-app/internal/task"
)

func main() {
	f := file.New[task.TContent]()
	err := cmd.HandleTrackerCLI(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
