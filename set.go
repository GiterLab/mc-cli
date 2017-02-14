package main

import (
	"errors"
	"fmt"
	"strconv"
)

func cmd_set(cmd []string) error {
	if len(cmd) != 3 && len(cmd) != 4 {
		return errors.New("param error")
	}
	if len(cmd) == 3 {
		err := MemcacheSet(cmd[1], []byte(cmd[2]))
		if err != nil {
			return errors.New(fmt.Sprintf("[E] %s", err))
		}
	}
	if len(cmd) == 4 {
		exp, err := strconv.Atoi(cmd[3])
		if err != nil {
			return errors.New(fmt.Sprintf("[E] %s", err))
		}
		err = MemcacheSetByExpired(cmd[1], []byte(cmd[2]), int32(exp))
		if err != nil {
			return errors.New(fmt.Sprintf("[E] %s", err))
		}
	}

	cprintln("OK", COLOR_GREEN)

	return nil
}
