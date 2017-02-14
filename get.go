package main

import (
	"errors"
	"fmt"
)

func cmd_get(cmd []string) error {
	if len(cmd) != 2 {
		return errors.New("param error")
	}
	b, err := MemcacheGet(cmd[1])
	if err != nil {
		return errors.New(fmt.Sprintf("[E] %s", err))
	}

	cprintln(string(b), COLOR_GREEN)

	return nil
}
