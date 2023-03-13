package gtk4extra

import (
	"github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// ItemList displays a list of items in a tabular fashion
type ItemList struct {
	*gtk.TreeView
	store  *gtk.ListStore
	cnames []string
	ctypes []glib.Type
}

// NewItemList create a new list of items.
func NewItemList() *ItemList {
	treeView := gtk.NewTreeView()

	l := &ItemList{treeView, nil, []string{}, []glib.Type{}}

	return l
}

// AddColumn adds a column to the table
func (i *ItemList) AddColumn(name string, t glib.Type) {
	i.ctypes = append(i.ctypes, t)
	i.cnames = append(i.cnames, name)

	i.AppendColumn(createColumn(name, len(i.cnames)-1))
	listStore := gtk.NewListStore(i.ctypes)
	i.SetModel(listStore)
	i.store = listStore
}

// AddColumn adds a column to the table with custom renderer
func (i *ItemList) AddColumnWithRenderer(name string, t glib.Type,   r gtk.CellRendererer) {
	i.ctypes = append(i.ctypes, t)
	i.cnames = append(i.cnames, name)

	i.AppendColumn(createColumnWithRenderer(name, len(i.cnames)-1, r))
	listStore := gtk.NewListStore(i.ctypes)
	i.SetModel(listStore)
	i.store = listStore
}

// Add adds a new item to the list.
func (i *ItemList) Add(items ...any) {
	if len(items) > len(i.cnames) {
		panic("number of items > number of columns")
	}

	values := []glib.Value{}
	for _, i := range items {
		values = append(values, *glib.NewValue(i))
	}
	colIds := []int{}
	for i, _ := range i.cnames {
		colIds = append(colIds, i)
	}
	i.store.Set(i.store.Append(),
		colIds,
		values,
	)
}

func createColumn(title string, id int) *gtk.TreeViewColumn {
	cellRenderer := gtk.NewCellRendererText()
	column := gtk.NewTreeViewColumn()
	column.SetTitle(title)

	column.PackEnd(cellRenderer, false)
	column.AddAttribute(cellRenderer, "text", int(id))
	column.SetResizable(true)

	return column
}

func createColumnWithRenderer(title string, id int, cellRenderer gtk.CellRendererer) *gtk.TreeViewColumn {
	column := gtk.NewTreeViewColumn()
	column.SetTitle(title)

	column.PackEnd(cellRenderer, false)
	if _, ok := cellRenderer.(*gtk.CellRendererProgress); ok {
		column.AddAttribute(cellRenderer, "value", int(id))
	}
	if _, ok := cellRenderer.(*gtk.CellRendererText); ok {
		column.AddAttribute(cellRenderer, "text", int(id))
	}
	if _, ok := cellRenderer.(*gtk.CellRendererPixbuf); ok {
		column.AddAttribute(cellRenderer, "pixbuf", int(id))
	}
	if _, ok := cellRenderer.(*gtk.CellRendererToggle); ok {
		column.AddAttribute(cellRenderer, "active", int(id))
	}
	if _, ok := cellRenderer.(*gtk.CellRendererCombo); ok {
		column.AddAttribute(cellRenderer, "text", int(id))
	}

	column.SetResizable(true)

	return column
}
