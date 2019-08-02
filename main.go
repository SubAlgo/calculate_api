package main // import "github.com/subalgo/mod"

import (
	"github.com/subalgo/mod/internal/app/cal"
	"log"
	"net/http"

	//"github.com/subalgo/mod/internal//cal"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/cal/", http.StripPrefix("/cal", cal.Handler()))

	log.Println("staring server on :8000")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}
}