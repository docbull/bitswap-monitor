package widget

import (
	"github.com/docbull/bitswap-monitor/conn"
	"github.com/rivo/tview"
)

<<<<<<< HEAD
// Widget
=======
>>>>>>> 7390f0e229051a223c6c7f485611ed82848ca8d4
type Widget struct {
	View   *tview.TextView
	Client *conn.Client
}

<<<<<<< HEAD
// NewWidget create widget
=======
>>>>>>> 7390f0e229051a223c6c7f485611ed82848ca8d4
func NewWidget() Widget {
	view := tview.NewTextView()
	view.SetBackgroundColor()
	view.SetBorder(true)
	view.SetBorderColor()
	view.SetDynamicColors(true)
	view.SetTextColor()
	view.SetTitleColor()
	view.SetWrap(false)
	view.SetScrollable(true)

	return Widget{
		View: view,
		// Client:
	}
}

func NewTviewGrid() *tview.Grid {
	grid := tview.NewGrid()
	grid.SetBackgroundColor()
	grid.SetColumns()
	grid.SetRows()
	grid.SetBorder(false)

	return grid
}

func NewTviewApp(grid *tview.Grid) *tview.Application {
	pages := tview.NewPages()
	pages.AddPage("grid", grid, true, true)
	pages.Box.SetBackgroundColor()

	tviewApp := tview.NewApplication()
	tviewApp.SetRoot(pages, true)

	return tviewApp
}
