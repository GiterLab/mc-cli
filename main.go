package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/context"
)

var VERSION = "v1.0.0"

func main() {
	host := flag.String("host", "", "address of memcached server")
	version := flag.Bool("v", false, "show version")
	flag.Parse()

	if *version {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	if *host == "" {
		fmt.Println("[E] User shoud specifies the address of the memcached server")
		os.Exit(0)
	}

	MemcacheInit(*host)

	// 命令行处理
	ctx, done := context.WithCancel(context.Background())
	go func() {
		fmt.Println("\033[32m==============================\033[0m")
		fmt.Println("\033[32mWelcome to memcache client    \033[0m")
		fmt.Println("\033[32m==============================\033[0m")
		running := true
		reader := bufio.NewReader(os.Stdin)
		for running {
			fmt.Printf("\033[32mMC>> \033[0m")
			data, _, _ := reader.ReadLine()
			data_lower := strings.ToLower(string(data))
			cmd := strings.Split(data_lower, " ")
			if len(cmd) == 0 || cmd[0] == "" {
				continue
			}

			switch cmd[0] {
			case "exit":
				fallthrough
			case "quit":
				fallthrough
			case "q":
				fmt.Println("  \033[32mQuit\033[0m")
				os.Exit(0)

			case "get":
				if len(cmd) != 2 {
					fmt.Println("  \033[32mparam error\033[0m")
					continue
				}
				b, err := MemcacheGet(cmd[0])
				if err != nil {
					fmt.Println("  \033[31m" + fmt.Sprintf("[E] %s", err) + "\033[0m")
					continue
				}
				fmt.Println(string(b))

				// 命令行列表
			case "list":
				fallthrough
			case "l":
				fmt.Println("  \033[32m------------------------\033[0m")
				fmt.Println("  \033[31mlist(l)\033[32m: list commands\033[0m")
				fmt.Println("  \033[31mquit(q)\033[32m: quit this app\033[0m")
				fmt.Println("  \033[31mat\033[32m: test command\033[0m")
				fmt.Println("  \033[31mactive(a)\033[32m: actice a device\033[0m")
				fmt.Println("  \033[31mstatus(s)\033[32m: check the status of device, offline or online\033[0m")
				fmt.Println("  \033[31mtemperature(t/r)\033[32m: query the current temperature of device\033[0m")
				fmt.Println("  \033[31mswitchon(son)\033[32m: turn on the switch of device, maybe a light\033[0m")
				fmt.Println("  \033[31mswitchoff(soff)\033[32m: turn off the switch of device, maybe a light\033[0m")

			default:
				fmt.Println("  \033[32mUnknown command\033[0m")
			}
			fmt.Println("")
		}
		done()
	}()

	<-ctx.Done()
}
