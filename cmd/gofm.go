package cmd

import (
	//"fmt"

	//"github.com/Redeltaz/gofm/pkg/tui"
	"fmt"

	"github.com/Redeltaz/gofm/pkg/controller"
)

func Root(c *controller.Controller) {
    //tui := tui.CreateTUI()
    
    //if err := tui.Start(); err != nil {
        //panic(err)
    //}
    //fmt.Println(c.CWD)
    c.GetDirContent()
    fmt.Println(c.CurrentContent)
}
