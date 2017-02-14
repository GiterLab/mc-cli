package main

import (
	"errors"
	"fmt"
)

func cmd_del(cmd []string) error {
	if len(cmd) != 2 {
		return errors.New("param error")
	}
	err := MemcacheDel(cmd[1])
	if err != nil {
		return errors.New(fmt.Sprintf("[E] %s", err))
	}

	cprintln("OK", COLOR_GREEN)

	return nil
}
