package pipe

import (
	"bufio"
	"fmt"
	"os/exec"
)

// PipeDemo sample

func PipeDemo() {
	cmd0 := exec.Command("echo", "-n", "my first Command from golang.")

	stdout, err := cmd0.StdoutPipe()

	if err != nil {
		fmt.Printf("Error: can not obtain the stdout pipe from Command No.0: %s", err)
		return
	}

	if err := cmd0.Start(); err != nil {
		fmt.Printf("Error: the Command No.0 can cant be startup: %s\n", err)
		return
	}

	/*output0 := make([]byte, 30)
	n, err := stdout.Read(output0)

	if err != nil {
		fmt.Printf("Error: Can not read data from pip: %s\n", err)
		return
	}

	fmt.Printf("%s\n", output0[:n])*/

	/*var outputBufo bytes.Buffer

	for {
		tempOuput := make([]byte, 5)
		n, err := stdout.Read(tempOuput)

		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("Error: Can not read data from the pipe: %s \n", err)
				return
			}
		}

		if n > 0 {
			outputBufo.Write(tempOuput[:n])
		}
	}

	fmt.Printf("%s\n", outputBufo.String())
	*/

	outputBuf0 := bufio.NewReader(stdout)
	output0, isRemain, err := outputBuf0.ReadLine()

	if err != nil {
		fmt.Printf("Error: Can not read data from the pipe: %s\n", err)
		return
	}

	fmt.Printf("isRemain = %v\n", isRemain)
	fmt.Printf("%s\n", string(output0))

}
