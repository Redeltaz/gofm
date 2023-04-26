package cmd

import (
	"bytes"
	"fmt"
	"os/exec"

    "github.com/Redeltaz/gofm/pkg/layout"
)

func Root() {
    tui := layout.TUI{}

    firstGrid := tui.CreateGrid()
    tui.Init("tui", firstGrid)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
    //grid.AddItem(main, 0, 0, 1, 1, 0, 0, false)


	//if err := tview.NewApplication().SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		//panic(err)
	//}
}

func command() {
    cmd := exec.Command("ls")
    buff := bytes.Buffer{}
    cmd.Stdout = &buff
    err := cmd.Run()

    fmt.Print(buff.String(), err)
}
