package net

import (
	"fmt"
	"net/http"
	"time"

	"github.com/wlbr/shorty/gotils"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	gotils.LogInfo("Receiving request.")

	keys, ok := r.URL.Query()["id"]

	if ok && len(keys) >= 1 {
		id := keys[0]
		gotils.LogInfo("Got ip %s from request url.", id)
	} else {

	}
	var err error
	if err != nil {

	}
	fmt.Fprintf(w, "<html>start<br>\n")
	for i := 0; i < 5; i++ {
		fmt.Fprintf(w, "%d<br>\n", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprintf(w, "end\n</html")

}
