package tui

import (
	"github.com/rivo/tview"
	"github.com/Redeltaz/gofm/pkg/controller"
)

type TUI struct {
    App *tview.Application
    Grid *tview.Grid

    Navigator *tview.TextView
    Preview *tview.TextView
    Info *tview.TextView
}

type TUISetup interface {
    Start() error
}

func CreateTUI(content []controller.File) *TUI{
    t := &TUI{}
    t.App = tview.NewApplication()

    //t.Navigator = tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(content)
    t.Preview = tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText("preview")
    t.Info = tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText("Info")

    t.Grid = tview.NewGrid().
            SetRows(0, 5).
            SetColumns(0, 0).
            SetBorders(true).
            AddItem(t.Navigator, 0, 0, 2, 1, 0, 0, false).
            AddItem(t.Preview, 0, 1, 2, 1, 0, 0, false).
            AddItem(t.Info, 1, 0, 1, 1, 0, 0, false)

    return t
}

func (t *TUI) Start() error {
    return t.App.SetRoot(t.Grid, true).Run()
}
