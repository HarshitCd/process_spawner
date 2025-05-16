package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func processSpawner(n int) []*exec.Cmd {
	processes := make([]*exec.Cmd, n)

	for i := 0; i < n; i++ {
		cmd := exec.Command("sleep", "1000")
		err := cmd.Start()
		if err != nil {
			fmt.Printf("failed to spawn %d process\n", i+1)
		}
		processes[i] = cmd
	}

	return processes
}

func processStopper(processes []*exec.Cmd) {
	for i, cmd := range processes {
		if cmd.Process == nil {
			continue
		}

		err := cmd.Process.Kill()
		if err != nil {
			fmt.Printf("failed to kill %d process\n", i+1)
		}
	}
}

func main() {
	var (
		processCount  int
		cycleCount    int
		cycleDuration int
		err           error
	)

	args := os.Args
	if len(args) == 4 {
		processCount, err = strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("error in the process count input")
			os.Exit(1)
		}

		cycleCount, err = strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("error in the cycle count input")
			os.Exit(1)
		}

		cycleDuration, err = strconv.Atoi(args[3])
		if err != nil {
			fmt.Println("error in the cycle count input")
			os.Exit(1)
		}
	} else {
		fmt.Println("input parameter count is not as expected")
		os.Exit(1)
	}

	for i := 0; i < cycleCount; i++ {
		fmt.Printf("cycle %d\n", i+1)
		processes := processSpawner(processCount)

		fmt.Printf("sleeping for %d\n", cycleDuration)
		time.Sleep(time.Second * time.Duration(cycleDuration))

		processStopper(processes)
	}
}
