package cmd

import (
	"bytes"
	"fmt"
	"os/exec"

    "github.com/rivo/tview"
)

func Root() {
	//newPrimitive := func(text string) tview.Primitive {
		//return tview.NewTextView().
			//SetTextAlign(tview.AlignCenter).
			//SetText(text)
	//}
	//main := newPrimitive("Main content")

	grid := tview.NewGrid().
		SetRows(0).
		SetColumns(0).
		SetBorders(true)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
    //grid.AddItem(main, 0, 0, 1, 1, 0, 0, false)


	if err := tview.NewApplication().SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
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
