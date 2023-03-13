package main

import (
	"net"
	"net/netip"
	"os"
	"regexp"

	"github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/rubiojr/gtk4extra"
	"tailscale.com/net/interfaces"
)

func main() {
	app := gtk.NewApplication("com.github.rubiojr.gtk4extra.addrlist", gio.ApplicationFlagsNone)
	app.ConnectActivate(func() { activate(app) })

	if code := app.Run(os.Args); code > 0 {
		os.Exit(int(code))
	}
}

func activate(app *gtk.Application) {
	win := gtk.NewApplicationWindow(app)
	win.SetTitle("Simple IP Address List")
	win.SetDefaultSize(600, 300)

	list := gtk4extra.NewItemList()
	win.SetChild(&list.Widget)

	list.AddColumn("Name", glib.TypeString)
	list.AddColumn("IPAddr", glib.TypeString)

	l, err := addrList()
	if err != nil {
		panic(err)
	}

	for k, v := range l {
		list.Add(k, v)
	}

	win.Show()
}

func addrList() (map[string]string, error) {
	l, err := interfaces.GetList()
	if err != nil {
		return nil, err
	}
	ipMap := map[string]string{}
	l.ForeachInterface(func(i interfaces.Interface, p []netip.Prefix) {
		addrs, err := i.Addrs()
		if err == nil {
			for _, addr := range addrs {
				if matched, _ := regexp.MatchString(`^en|wl|tail`, i.Name); matched {
					if ipv4Addr := addr.(*net.IPNet).IP.To4(); ipv4Addr == nil {
						continue
					}
					ipMap[i.Name] = addr.String()
				}
			}
		}
	})

	return ipMap, nil
}
