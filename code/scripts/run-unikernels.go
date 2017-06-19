package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const NUM_RUNS = 10

func main() {

	fmt.Printf("Starting %d unikernel instances...\n", NUM_RUNS)
	// for i := 0; i <= NUM_RUNS; i++ {
	// 	runUnik(i)
	// }

	time.Sleep(2 * time.Second)

	fmt.Println("Cleaning up Unik instances...")
	for i := 0; i <= NUM_RUNS; i++ {
		cleanupUnik(i)
	}

}

func runUnik(i int) {

	cmdName := "time"
	args := []string{
		"unik",
		"run",
		"--instanceName",
		fmt.Sprintf("argo%s", strconv.Itoa(i)),
		"--imageName",
		"argo",
	}

	cmd := exec.Command(cmdName, args...)
	// out, err := exec.Command(cmd, args...).Output()
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	os.Exit(1)
	// }
	// fmt.Println(string(out))

	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("docker build out | %s\n", scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		os.Exit(1)
	}
}

func cleanupUnik(i int) {
	cmd := "unik"
	args := []string{
		"delete-instance",
		"--force",
		"--instance",
		fmt.Sprintf("argo%s", strconv.Itoa(i)),
	}
	err := exec.Command(cmd, args...).Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
