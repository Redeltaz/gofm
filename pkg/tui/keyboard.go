package tui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

var (
    NavKeys
)

var NavKeys = []tcell.Key{tcell.KeyUp, tcell.KeyDown}

func (tui *TUI) InputHandler(event *tcell.EventKey) *tcell.EventKey {
    tui.App.Stop()
    switch event.Key(){
    case tcell.KeyRight:
        fmt.Println("right")
    case tcell.KeyLeft:
        fmt.Println("left")
    case 256:
        fmt.Println("letter")
    }
    return event
}
