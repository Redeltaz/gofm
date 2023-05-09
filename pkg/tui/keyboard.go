package tui

import (
	"fmt"

	"github.com/Redeltaz/gofm/pkg/controller"
	"github.com/gdamore/tcell/v2"
)

var navKeys = []rune{'h', 'j', 'k', 'l'}

func (tui *TUI) NavInputHandler(event *tcell.EventKey, c *controller.Controller) *tcell.EventKey {
    switch event.Key(){
    case tcell.KeyRight:
        fmt.Println("right")
    case tcell.KeyLeft:
        fmt.Println("left")
    case tcell.KeyRune:
        for _, key := range navKeys {
            if event.Rune() == key {
                tui.NavUpdateFocus(key, c)
            }
        }
    }
    return event
}

func (tui *TUI) NavUpdateFocus(pressedKey rune, c *controller.Controller) {
    currentItemIndex := tui.Navigator.GetCurrentItem()

    switch pressedKey {
    case 'k':
            tui.Navigator.SetCurrentItem(currentItemIndex - 1)
    case 'j':
        // Navigator.GetItemCount return 1 more item idk why so I let it like that
        if currentItemIndex == tui.Navigator.GetItemCount() - 1 {
            tui.Navigator.SetCurrentItem(0)
        } else {
            tui.Navigator.SetCurrentItem(currentItemIndex + 1)
        }
    }
    tui.UpdateInfo(c)
}
