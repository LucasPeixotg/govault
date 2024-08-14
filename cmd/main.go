package main

import (
	"os"
	"os/exec"

	"github.com/LucasPeixotg/govault/internal/logger"
)

func main() {
	mainLogger := logger.New("[ MAIN ]   ", logger.ERROR)
	//cmdLogger := logger.New("[ CMD ]   ", logger.ERROR)

	if len(os.Args[1:]) == 0 {
		mainLogger.Error("command must be specified as a cli argument")
		os.Exit(1)
	}

	commandContent := os.Args[1:]

	mainLogger.Info("command to be executed: ", commandContent)

	var cmd *exec.Cmd
	unshareCommand := "unshare -uipf --mount-proc "
	cmd = exec.Command(unshareCommand, commandContent...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	// TODO:
	// chroot

	// command to isolate namespaces
	// unshare -uipf --mount-proc /bin/sh
	//

	mainLogger.Info("executing...")
	cmd.Run()
}
