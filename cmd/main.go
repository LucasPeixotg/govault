package main

import (
	"os"
	"os/exec"

	"github.com/LucasPeixotg/govault/internal/logger"
)

func main() {
	mainLogger := logger.New("[ MAIN ]   ", logger.ERROR)
	cmdLogger := logger.New("[ CMD ]   ", logger.ERROR)

	if len(os.Args[1:]) == 0 {
		mainLogger.Error("command must be specified as a cli argument")
		os.Exit(1)
	}

	commandContent := os.Args[1:]

	mainLogger.Info("command to be executed: ", commandContent)

	var cmd *exec.Cmd
	if len(commandContent) > 1 {
		cmd = exec.Command(commandContent[0], commandContent[1:]...)
	} else {
		cmd = exec.Command(commandContent[0], commandContent[1:]...)
	}
	cmd.Stdout = cmdLogger
	cmd.Stdin = os.Stdin

	//unix.Unshare()

	// TODO:
	// isolate namespaces (network, pids, uts, mount [chroot])

	mainLogger.Info("executing...")
	cmd.Run()
}
