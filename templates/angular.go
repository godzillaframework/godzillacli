package templates

import (
	"log"
	"os"
	"os/exec"
)

type angularTemplate struct{}

func (a *angularTemplate) build(location string) {

}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func pipe(cmd *exec.Cmd) *exec.Cmd {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}
