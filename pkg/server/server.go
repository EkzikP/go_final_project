package server

import (
	"net/http"
)

const webDir = "./web"

func StartServer(port string) error {
	http.Handle("/", http.FileServer(http.Dir(webDir)))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return err
	}
	return nil
}
