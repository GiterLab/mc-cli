package main

import (
	"errors"
	"fmt"
)

func cmd_getmore(cmd []string) error {
	if len(cmd) != 2 {
		return errors.New("param error")
	}
	item, err := MemcacheGetMore(cmd[1])
	if err != nil {
		return errors.New(fmt.Sprintf("[E] %s", err))
	}

	cprintln(fmt.Sprintf("Key: %s", item.Key), COLOR_GREEN)
	cprintln(fmt.Sprintf("Value: %s", string(item.Value)), COLOR_GREEN)
	cprintln(fmt.Sprintf("Exp: %d", item.Expiration), COLOR_GREEN)
	cprintln(fmt.Sprintf("Flag: %b", item.Flags), COLOR_GREEN)

	return nil
}
