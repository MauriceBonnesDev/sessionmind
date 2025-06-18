package main

import (
	"fmt"
	"os"
)
import "github.com/mauricebonnesdev/sessionmind/types"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No command provided")
		return
	}

	switch os.Args[1] {
	case string(types.CmdTask):
		fmt.Println("Command = task")
		handleTaskSubCommand()
	case string(types.CmdSession):
		fmt.Println("Command = session")
		handleSessionSubCommand()
	default:
		fmt.Println("Command not found")
	}
}

func handleTaskSubCommand() {
	if len(os.Args) < 3 {
		fmt.Println("No subcommand provided")
		return
	}

	switch os.Args[2] {
	case string(types.TaskAdd):
		fmt.Println("SubCommand = add")
	case string(types.TaskDelete):
		fmt.Println("SubCommand = delete")
	case string(types.TaskEdit):
		fmt.Println("SubCommand = edit")
	default:
		fmt.Println("SubCommand not found")
	}
}

func handleSessionSubCommand() {
	if len(os.Args) < 3 {
		fmt.Println("No subcommand provided")
		return
	}

	switch os.Args[2] {
	case string(types.SessionStart):
		fmt.Println("SubCommand = start")
	case string(types.SessionPause):
		fmt.Println("SubCommand = pause")
	case string(types.SessionStop):
		fmt.Println("SubCommand = stop")
	default:
		fmt.Println("SubCommand not found")
	}
}
