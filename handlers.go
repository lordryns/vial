package main

import (
	"errors"
	"fmt"
	"os"
)

func getArg(index int) (string, error) {
	args := os.Args
	if len(args) > index {
		return args[index], nil
	}

	return "", errors.New(fmt.Sprintf("Not enough arguments to get index: %v", index))
}
