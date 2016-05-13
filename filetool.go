package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	dirname := flag.String("d", ".", "file server dir")
	hostport := flag.String("-s", ":5321", "file server dir")
	flag.Parse()

	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir(*dirname))))
	fmt.Println(http.ListenAndServe(*hostport, nil))
}
