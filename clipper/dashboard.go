package clipper

import (
	"fmt"
	"github.com/atotto/clipboard"
	"log"
	"net/http"
	"time"
)

type Clipboard []Clip

type Clip struct {
	Message string
	Date time.Time
}

var CurrentClipboard Clipboard


func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is where the clipboard goes"))
}

func (c Clip) New(s string) Clip {
	return Clip{s,time.Now()}
}

func (cb Clipboard) Append(c Clip) Clipboard {
	cb = append(cb,c)
	return cb
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
		}
		fmt.Println("Current clipboard line")
		fmt.Println(current)
		time.Sleep(time.Second * 30)
		fmt.Println("Clipboard History")
		fmt.Println(CurrentClipboard)
	}
}
