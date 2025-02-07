// Package main - main package
package main

import (
	"github/Asac2142/go-cli-app/cmd"
	"github/Asac2142/go-cli-app/file"
	"github/Asac2142/go-cli-app/task"
)

func main() {
	f := file.New[task.TContent]()
	cmd.HandleTrackerCLI(f)
}
