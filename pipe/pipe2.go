package pipe

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
)

func Pipe2() {

	cmd1 := exec.Command("ls", "al")

	stdout1, err := cmd1.StdoutPipe()
	if err != nil {
		fmt.Printf("Error: Can not obtain the stdout pipe for Command: %s\n", err)
		return
	}

	if err := cmd1.Start(); err != nil {
		fmt.Printf("Error: The command can not be startup: %s \n", err)
		return
	}

	cmd2 := exec.Command("grep", "*.go")
	var outputBuf2 bytes.Buffer
	cmd2.Stdout = &outputBuf2
	stdin2, err := cmd2.StdinPipe()

	if err != nil {
		fmt.Printf("Error: Can not obtain the stdin pip for command: %s\n", err)
		return
	}

	outputBuf1 := bufio.NewReader(stdout1)
	outputBuf1.WriteTo(stdin2)

	if err := cmd2.Start(); err != nil {
		fmt.Printf("Error: the command can not be startup: %s\n", err)
		return
	}

	err = stdin2.Close()
	if err != nil {
		fmt.Printf("Error: Can not close the stdio pipe : %s\n", err)
		return
	}

	if err := cmd2.Wait(); err != nil {
		fmt.Printf("Error: Cannot not wait for the command: %s \n", err)
		return
	}

	fmt.Println("the result : ", outputBuf2)

}
