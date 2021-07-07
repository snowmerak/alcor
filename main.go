package main

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/webview/webview"
)

//go:embed ui/public
var web embed.FS

func main() {
	port, err := GetFreePort()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(port)

	go func() {
		public, err := fs.Sub(web, "ui/public")
		if err != nil {
			log.Fatal(err)
		}

		http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.FS(public))))

		http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		})

		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatal(err)
		}
	}()

	const debug = true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetSize(600, 800, webview.HintFixed)
	w.Navigate("http://127.0.0.1:" + port + "/public/")
	w.Run()
}

func GetFreePort() (string, error) {
	conn, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	sp := strings.Split(conn.Addr().String(), ":")
	if len(sp) < 2 {
		return "", errors.New("not exist port")
	}
	return sp[1], nil
}
