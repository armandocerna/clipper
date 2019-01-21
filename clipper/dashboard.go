package clipper

import (
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

var CurrentClipboard Clipboard

var clipTemplate string = `
<h1>Clipboard History</h1>
<ul>
    {{range $i, $a := .}}
            <li>{{.Message}}	Age: {{minutesSince .Date}}</li>
	{{end}}
</ul>
</ul>
`

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
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
}

func (c Clip) New(s string) *Clip {
	c = Clip{s, time.Now()}
	return &c
}

func (cb *Clipboard) Append(c *Clip) {
	*cb = append(*cb, *c)
}

func ReadClipboard() {
	var lastCopy string
	var c Clip
	for {
		current, err := clipboard.ReadAll()
		if err != nil {
			log.Fatalf("Error reading clipboard %v", err)
		}
		if current != lastCopy {
			CurrentClipboard.Append(c.New(current))
			lastCopy = current
		}
		time.Sleep(time.Second * 5)
	}
}
