package main

import (
	"os/exec"
	"time"
)

func main() {
	command := exec.Command("sleep", "100")
	_ = command.Run()
	time.Sleep(10 * time.Second)
}
