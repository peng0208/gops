package gops_client

import (
	"os/exec"
	"fmt"
)

type CronTask int

func (c *CronTask) ExecShell(arg string, reply *interface{}) error {
	bin := "/bin/bash"
	run := fmt.Sprintf("source /etc/profile && %s", arg)
	cmd := exec.Command(bin, "-c", run)
	out, err := cmd.CombinedOutput()
	*reply = string(out)
	return err
}