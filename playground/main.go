package main

import (
	"log"
	"net"
	"os"
	"os/exec"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	c, err := net.Dial("unix", os.Args[2])
	if err != nil {
		return err
	}
	f, err := c.(*net.UnixConn).File()
	if err != nil {
		return err
	}

	args := []string{
		"--env=WHOAMI=me",
		"--env=WASI_FD=3", // 3 corresponds to first specified --fd (0, 1 and 2 are taken)
		"--fd=4",          // 4 corresponds to second fd in ExtraFiles.
		//	fmt.Sprintf("--mapdir=/:container"),
		os.Args[1],
	}

	f.Sync()

	cmd := exec.Command("wasmtime", args...)
	log.Printf("running: %v\n", cmd.Args)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = []*os.File{os.NewFile(1, ""), f}

	return cmd.Run()
}
