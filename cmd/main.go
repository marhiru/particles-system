package main

import (
	"fmt"
	"os"
	"os/exec"
	particles "particles-system/ascii"
	"runtime"
	"time"
)

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}

	if runtime.GOOS != "windows" {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	width, height := 20, 10
	coffeeSystem := particles.NewCoffeeSystem(width, height)

	coffeeSystem.Start()

	for {
		clearScreen()
		coffeeSystem.Update()

		frame := coffeeSystem.Print()
		for _, row := range frame {
			fmt.Println(string(row))
		}

		time.Sleep(500 * time.Millisecond)
	}
}
