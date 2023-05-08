package cmd

import (
    "github.com/Redeltaz/gofm/pkg/tui"
	"github.com/Redeltaz/gofm/pkg/controller"
)

func Root() {
    controller := controller.InitController()
    controller.SetCurrentContent()

    tui := tui.CreateTUI(controller)
    
    if err := tui.Start(); err != nil {
        panic(err)
    }
}
