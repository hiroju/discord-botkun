package usecase

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

type (
	CommandParser struct {
		command_name   string
		command_string string
		flag_set       *flag.FlagSet
		args           map[string][]string
	}
)

const (
	COMMAND_PREFIX = "/"
)

func ParseCommand(commandName string, commandString string) (map[string][]string, error) {
	if !strings.HasPrefix(commandString, COMMAND_PREFIX+commandName) {
		return nil, errors.New("this is not expect command")
	}

	res := strings.SplitAfter(commandString, COMMAND_PREFIX+commandName)
	if len(res) < 2 {
		return nil, nil
	}

	args := map[string][]string{}
	res = strings.SplitAfter(strings.TrimSpace(res[1]), "-")
	for i := 0; i < len(res); i++ {
		if res[i] == "-" {
			continue
		}
		options := strings.SplitAfter(res[i], " ")
		optionName := ""
		for argCount := 0; argCount < len(options); argCount++ {
			fmt.Fprintln(os.Stdout, options)
			if options[argCount] == " " || options[argCount] == "-" {
				continue
			}
			if optionName == "" {
				optionName = strings.TrimSpace(options[argCount])
			}
			if _, ok := args[optionName]; !ok {
				args[optionName] = []string{}
				continue
			}
			args[optionName] = append(args[optionName], strings.TrimSpace(options[argCount]))
		}
	}
	return args, nil
}
