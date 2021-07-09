package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"errors"
	"flag"
	"io"
	"io/fs"
	"log"
	"net"
	"net/http"
	"strings"

	"alcor/client/book"
	clientSocket "alcor/client/socket"
	"alcor/client/wallet"
	"alcor/config"

	"github.com/webview/webview"
)

//go:embed ui/public
var web embed.FS

func main() {
	isMaster := flag.Bool("master", false, "")
	flag.Parse()

	// port, err := GetFreePort()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(port)
	port := "9999"

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

		http.HandleFunc("/wallet", func(rw http.ResponseWriter, r *http.Request) {
			switch strings.ToUpper(r.Method) {
			case "GET":
				rw.WriteHeader(202)
				accounts, _ := wallet.List()
				buffer := bytes.NewBuffer(nil)
				encoder := json.NewEncoder(buffer)
				encoder.Encode(accounts)
				rw.Write(buffer.Bytes())
			case "PUT":
				rw.WriteHeader(202)
				bs, err := io.ReadAll(r.Body)
				if err != nil {
					rw.Write(bs)
					return
				}
				clientSocket.EnrollClient("ws://"+config.HigherAddress+":"+port+"/client/enroll", string(bs))
			}
		})

		http.HandleFunc("/client", func(rw http.ResponseWriter, r *http.Request) {
			switch strings.ToUpper(r.Method) {
			case "GET":
				rw.WriteHeader(202)
				clients, _ := book.List()
				buffer := bytes.NewBuffer(nil)
				encoder := json.NewEncoder(buffer)
				encoder.Encode(clients)
				rw.Write(buffer.Bytes())
			}
		})

		http.HandleFunc("/config/higheraddress", func(rw http.ResponseWriter, r *http.Request) {
			switch strings.ToUpper(r.Method) {
			case "GET":
				rw.WriteHeader(202)
				rw.Write([]byte(config.HigherAddress))
			case "PUT":
				rw.WriteHeader(202)
			case "DELETE":
				rw.WriteHeader(202)
			case "PATCH":
				rw.WriteHeader(202)
			}
		})

		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatal(err)
		}
	}()

	const debug = true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetSize(800, 640, webview.HintNone)
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
