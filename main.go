package main

import (
	"fmt"
	"os"
	"path/filepath"
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
    specifies the number for timer.
COMMANDS
    specifies the commands execute after timer.
    If no commands are specified, nml notifies by printing message
    "MESSAGE FROM NML ARRIVE!! (<NUMBER> <UNIT> left)" to STDOUT.`, programName, programName)
}

func goMain(args []string) int {
	fmt.Println(helpMessage(args[0]))
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
