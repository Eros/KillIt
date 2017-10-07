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
	"time"
	"os"
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

func shutdownSequence(conf * config){
	done := make(chan struct{})
	c1 := new(commandList)
	wg := new(sync.WaitGroup)
	t := time.NewTimer(time.Duration(conf.ShutdownTimeout) * time.Millisecond)

	if conf.Commands != nil {
		go func() {
			select {
			case <-done:
			case <-t.C:
				log.Println("Timed out")
				c1.KillAll()
			}
			if conf.Shutdown {
				log.Println("Shutting down")
				if err := shutdownNow(); err != nil {
					log.Fatal("Error shutting down: ", err)
				} else {
					log.Println("Commands finished")
					os.Exit(0)
				}
			}
		}()
	}
}
