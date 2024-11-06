package html

import (
	"fmt"

	"github.com/alfiehiscox/submarines/pkg/board"
	"github.com/alfiehiscox/submarines/pkg/cell"
	. "maragu.dev/gomponents"
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

const (
	HTMX_SOURCE    = "https://unpkg.com/htmx.org@2.0.3"
	HTMX_INTEGRITY = "sha384-0895/pl2MU10Hqc6jd4RvrthNlDiE9U1tWmX7WRESftEDRosgxNsQG/Ze9YMRzHq"
)

func Index() Node {
	return page(H1(Class("text-xl"), Text("Battleships")))
}

func PlaceShips(board board.Board) Node {
	return page(
		Div(Class("w-1/3 h-screen flex flex-col items-center justify-center"),
			ShipGallery(0),
			Div(Class("w-4/5 grid grid-cols-10 gap-2"),
				Map(board, func(cell cell.Cell) Node {
					return Cell(cell.Occupied, "hover:bg-blue-500")
				}),
			),
		),
	)
}

func Repeat(n int, node Node) Node {
	group := make(Group, n)
	for i := range group {
		group[i] = node
	}
	return group
}

func ShipGallery(chosen int) Node {
	return Div(ID("ship-gallery"),
		Class("w-full h-16 flex justify-around items-center"),
		Ship(1, 5, chosen),
		Ship(2, 4, chosen),
		Ship(3, 4, chosen),
		Ship(4, 3, chosen),
		Ship(5, 2, chosen),
	)
}

func Ship(id, count, chosen int) Node {
	var style string
	if id == chosen {
		style = "bg-blue-500"
	} else {
		style = "group-hover:bg-blue-500"
	}

	return Div(
		htmx.Get(fmt.Sprintf("/ship-select/%d", id)),
		htmx.Trigger("click"),
		htmx.Swap("outerHTML"),
		htmx.Target("#ship-gallery"),
		Class("w-1/5 flex justify-center items-center group"),
		Repeat(count, Cell(false, style)),
	)
}

func Cell(occupied bool, style string) Node {
	if occupied {
		return Div(Class("rounded w-4 h-4 " + style))
	} else {
		return Div(
			Class("rounded border w-5 h-5 " + style),
		)
	}
}

func page(children ...Node) Node {
	return HTML5(HTML5Props{
		Title:       "battleships",
		Description: "battleships",
		Language:    "en",
		Head: []Node{
			Link(Rel("stylesheet"), Href("static/app.css")),
			Script(Src(HTMX_SOURCE), Integrity(HTMX_INTEGRITY), CrossOrigin("anonymous")),
		},
		Body: []Node{
			Div(
				Class("w-full h-screen flex justify-center items-center"),
				Group(children),
			),
		},
	})
}
