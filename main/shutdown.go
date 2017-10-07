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
