package main

import (
	"embed"
	"errors"
	"flag"
	"io/fs"
	"log"
	"net"
	"net/http"
	"strings"

	clientSocket "alcor/client/socket"

	"github.com/webview/webview"
)

//go:embed ui/public
var web embed.FS

func main() {
	isMaster := flag.Bool("master", false, "")
	clientEnrollTest := flag.Bool("client_enroll", false, "")
	flag.Parse()

	// port, err := GetFreePort()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(port)
	port := "9999"

	if !(*clientEnrollTest) {
		go func() {
			public, err := fs.Sub(web, "ui/public")
			if err != nil {
				log.Fatal(err)
			}

			http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.FS(public))))

			if *isMaster {
				http.HandleFunc("/client/enroll", func(rw http.ResponseWriter, r *http.Request) {
					if err := clientSocket.EnrollServer(rw, r); err != nil {
						log.Panicln(err)
					}
				})
			}

			if err := http.ListenAndServe(":"+port, nil); err != nil {
				log.Fatal(err)
			}
		}()

		const debug = true
		w := webview.New(debug)
		defer w.Destroy()
		w.SetSize(1200, 860, webview.HintFixed)
		w.Navigate("http://127.0.0.1:" + port + "/public/")
		w.Run()
	}

	if *clientEnrollTest {
		log.Println(clientSocket.EnrollClient("ws://127.0.0.1:"+port+"/client/enroll", "kym985478"))
	}
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
