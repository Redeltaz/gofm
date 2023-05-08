package tui

import (
	"github.com/Redeltaz/gofm/pkg/controller"
	"github.com/rivo/tview"
)

type TUI struct {
	App  *tview.Application
	Grid *tview.Grid

	Navigator *tview.List
	Preview   *tview.TextView
	Info      *tview.TextView
}

func CreateTUI(c *controller.Controller) *TUI {
	t := &TUI{}
	t.App = tview.NewApplication()
	t.App.SetInputCapture(t.InputHandler)

	t.Navigator = tview.NewList()

	t.Preview = tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText("preview")
	t.Info = tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText("Info")

	t.UpdateNavigator(c)

	t.Grid = tview.NewGrid().
		SetRows(0, 5).
		SetColumns(0, 0).
		SetBorders(true).
		AddItem(t.Navigator, 0, 0, 1, 1, 0, 0, false).
		AddItem(t.Preview, 0, 1, 2, 1, 0, 0, false).
		AddItem(t.Info, 1, 0, 1, 1, 0, 0, false)

	return t
}

func (t *TUI) Start() error {
	return t.App.SetRoot(t.Grid, true).SetFocus(t.Navigator).Run()
}

func (t *TUI) UpdateNavigator(c *controller.Controller) {
	t.Navigator.Clear()

	for _, file := range c.CurrentContent {
		if file.Permissions.IsDir() {
			file.Name = file.Name + "/"
		}
		t.Navigator.AddItem(file.Name, "", '1', nil)
	}
}
