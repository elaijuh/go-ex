package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	log.SetFlags(0)
	pid := os.Getpid()
	log.Println("Current pid:", pid)

	cmd := exec.Command("sleep", "1")
	procErr := make(chan error)

	// start a new process in goroutine
	go func() {
		procErr <- cmd.Run()
	}()

	// wait a little bit for the process to start
	time.Sleep(100 * time.Millisecond)
	// (p *Process) Wait() (*ProcessState, error)
	log.Printf("Child pid: %d", cmd.Process.Pid)

	// sleep for a little bit until the child process finished
	nwait := 2 * time.Second
	time.Sleep(nwait)

	select {
	default:
		log.Println("Process is running")
	case e := <-procErr:
		if e != nil {
			log.Printf("Process exited: %s", e)
		} else {
			log.Println("Process exited without error")
		}
	}

}
