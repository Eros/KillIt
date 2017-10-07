package main

import (
	_ "log"
	_ "os"
	_ "os/exec"
	_ "strings"
	_ "sync"
	_ "time"

	"sync"
	"os/exec"
	"log"
)

type commandList struct {
	sync.Mutex
	commands []*exec.Cmd
}

func (c1 * commandList) Add(cmd *exec.Cmd){
	c1.Lock()
	c1.commands = append(c1.commands, cmd)
	c1.Unlock()
}

func (c1 * commandList) KillAll(){
	for _, c := range c1.commands{
		if c.Process != nil {
			_ = c.Process.Kill()
		}
	}
	c1.commands = nil
}

func parallelExecute(cmd *exec.Cmd, wg *sync.WaitGroup){
	err := cmd.Start()
	if err != nil {
		log.Printf("Error starting pc %s: %s", cmd.Path, err)
	}
	wg.Add(1)
	err = cmd.Wait()
	wg.Done()
	if err != nil {
		log.Printf("Error starting pc %s: %s", cmd.Path, err)
	}
}
