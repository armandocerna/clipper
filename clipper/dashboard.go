package clipper

import (
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/atotto/clipboard"
)

type Clipboard []Clip

type Clip struct {
	Message string
	Date    time.Time
}

var upgrader = websocket.Upgrader{}
var CurrentClipboard Clipboard

var clipTemplate string = `
<h1>Clipboard History</h1>
<ul>
    {{range $i, $a := .}}
            <li>{{.Message}}	Age: {{minutesSince .Date}}</li>
	{{end}}
</ul>
`

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}

		tmpl, err := template.New("").Funcs(template.FuncMap{
			"minutesSince": func(t time.Time) string {
				return time.Since(t).String()
			},
		}).Parse(clipTemplate)
		if err != nil {
			log.Fatalf("error parsing template: %v", err)
		}

		err = tmpl.Execute(w, CurrentClipboard)
		if err != nil {
			log.Fatalf("error templating clipboard history: %v", err)
		}
		ReadClipboard(ws)
	}
}

func (c Clip) New(s string) *Clip {
	c = Clip{s, time.Now()}
	return &c
}

func (cb *Clipboard) Append(c *Clip) {
	*cb = append(*cb, *c)
}

func ReadClipboard(ws *websocket.Conn) {
	var lastCopy string
	var c Clip
	defer ws.Close()
	for {
		current, err := clipboard.ReadAll()
		if err != nil {
			log.Fatalf("error reading clipboard %v", err)
		}
		if current != lastCopy {
			CurrentClipboard.Append(c.New(current))
			lastCopy = current
			if err := ws.WriteMessage(websocket.TextMessage, []byte("message!")); err != nil {
				err := ws.Close()
				if err != nil {
					log.Printf("error closing: %v", err)
				}
				break
			}
		}
		time.Sleep(time.Second * 5)
	}
}
