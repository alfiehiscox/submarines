package html

import (
	"github.com/alfiehiscox/submarines/pkg/board"
	"github.com/alfiehiscox/submarines/pkg/cell"
	. "maragu.dev/gomponents"
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
		Div(Class("w-1/3 grid grid-cols-10 gap-2"),
			Map(board, func(cell cell.Cell) Node {
				return Cell(cell.Occupied)
			}),
		),
	)
}

func Cell(occupied bool) Node {
	if occupied {
		return Div(Class("rounded bg-red w-4 h-4 hover:bg-blue-500"))
	} else {
		return Div(
			Class("rounded bg-white border w-5 h-5 hover:bg-blue-500"),
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
			Script(Href(HTMX_SOURCE), Integrity(HTMX_INTEGRITY), CrossOrigin("anonymous")),
		},
		Body: []Node{
			Div(
				Class("w-full h-screen flex justify-center items-center"),
				Group(children),
			),
		},
	})
}
