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
    CreateGrid() *tview.Grid
}

func (t TUI) Init(name string, root tview.Primitive) (*TUI, error){
    t.Name = name
    t.App = tview.NewApplication()

    err := t.App.SetRoot(root, true).SetFocus(root).Run()
    if err != nil {
        return nil, err
    }

    return &t, nil
}

func (t TUI) CreateGrid() *tview.Grid {
    return tview.NewGrid().
            SetRows(1, 2).
            SetColumns(2,3).
            SetBorders(true)
}
