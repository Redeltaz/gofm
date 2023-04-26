package layout

import (
    "github.com/rivo/tview"
)

type TUI struct {
    Name string
    App *tview.Application
}

type TUISetup interface {
    Init() *tview.Application
}

func (l TUI) Init(name string) *TUI{
    t := TUI{Name: name}
    t.App = tview.NewApplication()

    return &t
}
