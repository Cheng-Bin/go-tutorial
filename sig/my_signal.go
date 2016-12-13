package signal

import (
	"os/exec"
	"fmt"
	"strconv"
	"os"
	"syscall"
	"bytes"
	"io"
)

func MySignal() {

    cmds := []*exec.Cmd{
        exec.Command("ps", "aux"),
        exec.Command("grep", "mysignal.go"),
        exec.Command("grep", "-v", "grep"),
        exec.Command("awk", "{print $2}")
    }

    output, err := runCmds(cmds)
    if err != nil {
        fmt.Printf("Command Execution Error: %s\n", err)
        return
    }

    pid := strconv.Atoi(output)
    proc, err := os.FindProcess(pid)

    err = proc.Signal(syscall.SIGINT)



}


func runCmds(cmds []*exec.Cmd) ([]string, error) {

    var outputbuffer bytes.Buffer
    var stdout io.ReadCloser
    var stdin io.WriteCloser

    for _, cmd := range cmds {

        if stdout != nil {
            outputbuffer := bufio.NewReader(stdout)
            stdin, err := cmd.StdinPipe()
            if err != nil {

            }
            outputbuffer.WriteTo(stdin)
        }

        stdout, err := cmd.StdoutPipe()
        if err != nil {
            Failed(err, "Get StdoutPipe failed.")
        }

        if err := cmd.Start(); err != nil {
            Failed(err, "Start failed.")
        }


    }

    return nil
}

func Failed(err error, msg error) {
    fmt.Printf("error: %s[%s]", msg, err)
    return
}