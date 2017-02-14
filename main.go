package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/context"
)

var VERSION = "v0.0.1"

func main() {
	host := flag.String("host", "", "address of memcached server")
	version := flag.Bool("v", false, "show version")
	flag.Parse()

	if *version {
		cprintln(VERSION, COLOR_NONE)
		os.Exit(0)
	}

	if *host == "" {
		*host = "127.0.0.1:11211"
	}

	MemcacheInit(*host)

	// handle commands
	ctx, done := context.WithCancel(context.Background())
	go func() {
		cprintln("==============================", COLOR_GREEN)
		cprintln("memcache client "+VERSION, COLOR_GREEN)
		cprintln("==============================", COLOR_GREEN)
		running := true
		reader := bufio.NewReader(os.Stdin)
	Shell:
		for running {
			cprint("MC>> ", COLOR_GREEN)
			data, _, _ := reader.ReadLine()
			data_lower := strings.ToLower(string(data))
			cmd := strings.Split(data_lower, " ")
			if len(cmd) == 0 || cmd[0] == "" {
				continue
			}

			switch cmd[0] {
			case "bye":
				fallthrough
			case "exit":
				fallthrough
			case "e":
				fallthrough
			case "quit":
				fallthrough
			case "q":
				fmt.Println("  \033[32mQuit\033[0m")
				break Shell

			case "get": // get key
				err := cmd_get(cmd)
				if err != nil {
					cprintln(err.Error(), COLOR_RED)
					continue
				}

			case "getmore": // getmore key
				err := cmd_getmore(cmd)
				if err != nil {
					cprintln(err.Error(), COLOR_RED)
					continue
				}

			case "set": // set key value, set key value 60
				err := cmd_set(cmd)
				if err != nil {
					cprintln(err.Error(), COLOR_RED)
					continue
				}

			case "del": // del key
				err := cmd_del(cmd)
				if err != nil {
					cprintln(err.Error(), COLOR_RED)
					continue
				}

				// 命令行列表
			case "list":
				fallthrough
			case "l":
				fallthrough
			case "h":
				cprintln("------------------------", COLOR_GREEN)
				cprintln("  \033[31mset\033[32m: set key value, set key value expiration_time\033[0m", COLOR_NONE)
				cprintln("  \033[31mget\033[32m: get key\033[0m", COLOR_NONE)
				cprintln("  \033[31mgetmore\033[32m: getmore key\033[0m", COLOR_NONE)
				cprintln("  \033[31mdel\033[32m: del key\033[0m", COLOR_NONE)

				cprintln("  \033[31mlist(l)\033[32m: list commands\033[0m", COLOR_NONE)
				cprintln("  \033[31mquit(q)\033[32m: quit this app\033[0m", COLOR_NONE)
				cprintln("  \033[31mexit(e)\033[32m: quit this app\033[0m", COLOR_NONE)

			default:
				cprintln("Unknown command", COLOR_RED)
			}
			fmt.Println("")
		}
		done()
	}()

	<-ctx.Done()
}
