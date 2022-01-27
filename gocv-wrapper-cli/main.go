package main

import (
	"bufio"
	"fmt"
	"os"
	"skkim-01/gocv-wrapper/src/ifaces"
	"strings"
)

func main() {
	_shell()
}

func _shell() {
	var cmd string = ""
	var bExit bool = false

	api := ifaces.IFace{}

	for {
		if bExit {
			break
		}

		fmt.Print("gocv-wrapper-cli: > ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			cmd = scanner.Text()
		}
		slCmd := strings.Split(cmd, " ")
		switch slCmd[0] {
		case "list":
			results := api.GetAvaliCamIdxs()
			fmt.Println(results)

		case "start":
			api.StartCam(0)

		case "stop":
			api.Stop()

		case "exit":
			bExit = true

		case "":
			break

		default:
			fmt.Println("Not Support Command:", slCmd[0])
		}
	}
}
