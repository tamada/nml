package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	flag "github.com/spf13/pflag"
)

// Command shows the executing command after time.
type Command interface {
	Execute() int
}

type printCommand struct {
	message string
}

func (pc *printCommand) Execute() int {
	fmt.Println(pc.message)
	return 0
}

type cliCommand struct {
	args []string
}

func (cc *cliCommand) Execute() int {
	cmd := cc.args[0]
	args := []string{}
	if len(cc.args) > 1 {
		args = cc.args[1:]
	}
	command := exec.Command(cmd, args...)
	return execImpl(command)
}

func execImpl(cmd *exec.Cmd) int {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("exec error: %s\n", err.Error())
		return 1
	}
	return cmd.ProcessState.ExitCode()
}

type options struct {
	time       *timer
	header     bool
	background bool
	help       bool
	command    Command
}

func applyNumber(opts *options, args []string) error {
	if len(args) <= 1 {
		return errors.New("<NUMBER> is mandatory.")
	}
	value, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return fmt.Errorf("%s: %w", args[1], err)
	}
	opts.time.number = value
	return nil
}

func applyGivenCommands(opts *options, args []string, index int) (*options, error) {
	if index >= len(args) {
		return nil, fmt.Errorf("no commands were given")
	}
	opts.command = &cliCommand{args: args[index:]}
	return opts, nil
}

func applyDefaultMessage(opts *options, args []string) (*options, error) {
	time := opts.time
	opts.command = &printCommand{message: fmt.Sprintf("MESSAGE FROM NML ARRIVE!! (%d %s left)", time.number, time.unit)}
	return opts, nil
}

func applyArguments(opts *options, args []string, original []string) (*options, error) {
	err := applyNumber(opts, args)
	if err != nil {
		return nil, err
	}
	for i, arg := range original {
		if arg == "--" {
			return applyGivenCommands(opts, original, i+1)
		}
	}
	return applyDefaultMessage(opts, args)
}

func validate(time *timer) error {
	availableUnits := []string{"nsec", "usec", "msec", "sec", "min", "hour"}
	value := strings.TrimSpace(strings.ToLower(time.unit))
	for _, available := range availableUnits {
		if value == available {
			time.unit = available
			return nil
		}
	}
	return fmt.Errorf("%s: unknown unit", time.unit)
}

func parseArguments(opts *options, args []string, original []string) (*options, error) {
	err := validate(opts.time)
	if err != nil {
		return nil, err
	}
	return applyArguments(opts, args, original)
}

func parseArgs(args []string) (*options, error) {
	opts := &options{time: &timer{}, header: false, help: false}
	flags := flag.NewFlagSet("nml", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args[0])) }
	flags.StringVarP(&opts.time.unit, "unit", "u", "min", "specifies the time unit. Default is min")
	flags.BoolVarP(&opts.header, "with-header", "H", false, "shows the header on notification")
	flags.BoolVarP(&opts.help, "help", "h", false, "prints this message")
	flags.Parse(args)
	if !opts.help {
		// pflags が "--" を解釈し，削除してしまうため，オリジナルの args も渡している．
		return parseArguments(opts, flags.Args(), args)
	}
	return opts, nil
}
