package cmd

import (
	"bytes"
	"fmt"
	"os/exec"

    "github.com/Redeltaz/gofm/pkg/tui"
)

func Root() {
    tui := tui.CreateTUI()
    
    if err := tui.Start(); err != nil {
        panic(err)
    }
}

func command() {
    cmd := exec.Command("ls")
    buff := bytes.Buffer{}
    cmd.Stdout = &buff
    err := cmd.Run()

    fmt.Print(buff.String(), err)
}
