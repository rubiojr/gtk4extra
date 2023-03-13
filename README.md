# Reusable Go GTK4 components

## ItemList

A list of items in a tabular fashion.

```Go
package main

import (
	"os"

	"github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/rubiojr/gtk4extra"
	"tailscale.com/net/interfaces"
)

func main() {
	app := gtk.NewApplication("simplelist", gio.ApplicationFlagsNone)
	app.ConnectActivate(func() { activate(app) })

	if code := app.Run(os.Args); code > 0 {
		os.Exit(int(code))
	}
}

func activate(app *gtk.Application) {
	win := gtk.NewApplicationWindow(app)
	win.SetTitle("Simple List")
	win.SetDefaultSize(600, 300)

	list := gtk4extra.NewItemList()
	win.SetChild(&list.Widget)

	list.AddColumn("Name", glib.TypeString)
	list.AddColumn("Description", glib.TypeString)
  
  list.Add("foobar", "this is a foobar")

	win.Show()
}

```
