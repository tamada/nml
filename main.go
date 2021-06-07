package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func helpMessage(originalProgramName string) string {
	programName := filepath.Base(originalProgramName)
	return fmt.Sprintf(`%s [OPTIONS] <NUMBER> [-- <COMMANDS...>]
%s means 'notify me, later!'
OPTIONS
    -b, --background      runs nml on background.  Same as 'nml ... &'.
    -H, --with-header     shows the header on notification.
    -u, --unit <UNIT>     specifies the time unit. Default is min.
                          Available units are: nsec, msec, sec, min, and hour.
    -h, --help            prints this message.
NUMBER
    specifies the number for timer by the integer value.
COMMANDS
    specifies the commands execute after timer.
    If no commands are specified, nml notifies by printing message
    "MESSAGE FROM NML ARRIVE!! (<NUMBER> <UNIT> left)" to STDOUT.`, programName, programName)
}

func handle(opts *options) int {
	return opts.command.Execute()
}

func perform(opts *options) int {
	select {
	case <-time.After(opts.time.Duration()):
		return handle(opts)
	}
}

func goMain(args []string) int {
	opts, err := parseArgs(args)
	if err != nil {
		fmt.Printf("parsing args fail: %s\n", err.Error())
		fmt.Println(helpMessage(filepath.Base(args[0])))
		return 1
	}
	if opts.help {
		fmt.Println(helpMessage(filepath.Base(args[0])))
		return 0
	}
	return perform(opts)
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
