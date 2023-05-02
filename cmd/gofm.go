package cmd

import (
    "fmt"

    "github.com/Redeltaz/gofm/pkg/tui"
	"github.com/Redeltaz/gofm/pkg/controller"
)

func Root() {
    controller := controller.InitController()
    controller.GetDirContent()

    tui := tui.CreateTUI(controller.CurrentContent)
    
    if err := tui.Start(); err != nil {
        panic(err)
    }
    fmt.Println(controller.CurrentContent)
}
