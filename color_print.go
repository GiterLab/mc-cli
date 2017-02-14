package main

import (
	"fmt"
)

const (
	COLOR_NONE    = ""
	COLOR_RESTORE = "\033[0m"
	COLOR_RED     = "\033[31m"
	COLOR_GREEN   = "\033[32m"
)

func cprint(str string, color string) {
	if color != COLOR_NONE {
		fmt.Print(color + str + COLOR_RESTORE)
		return
	}
	fmt.Println(str)
}

func cprintln(str string, color string) {
	if color != COLOR_NONE {
		fmt.Println(color + str + COLOR_RESTORE)
		return
	}
	fmt.Println(str)
}
