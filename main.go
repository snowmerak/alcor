package main

import (
	"embed"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"strings"

	"alcor/client/book"
	clientSocket "alcor/client/socket"
	"alcor/client/wallet"
	"alcor/config"
	transactionSocket "alcor/transaction/socket"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/webview/webview"
)

//go:embed ui/public
var web embed.FS

func main() {
	isMaster := flag.Bool("master", false, "")
	useWeb := flag.Bool("web", false, "")
	flag.Parse()

	// port, err := GetFreePort()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(port)
	port := "9999"
	app := fiber.New()

	go func() {
		public, err := fs.Sub(web, "ui/public")
		if err != nil {
			log.Fatal(err)
		}

		app.Use("/public/", filesystem.New(filesystem.Config{Root: http.FS(public)}))

		if *isMaster {
			app.Get("/client/enroll", func(c *fiber.Ctx) error {
				if err := clientSocket.EnrollServer(c.Context()); err != nil {
					return err
				}
				return nil
			})

			app.Get("/transaction/enroll", func(c *fiber.Ctx) error {
				if err := transactionSocket.EnrollServer(c.Context()); err != nil {
					return err
				}
				return nil
			})
		}

		app.Get("/wallet", func(c *fiber.Ctx) error {
			accounts, _ := wallet.List()
			return c.JSON(accounts)
		})
		app.Get("/wallet/:id", func(c *fiber.Ctx) error {
			id := c.Params("id", "")
			hash, err := hex.DecodeString(id)
			if err != nil {
				return c.JSON(errors.New("invalid id"))
			}
			account, err := wallet.Get(hash)
			if err != nil {
				return c.JSON(err)
			}
			return c.JSON(account)
		})
		app.Put("/wallet", func(c *fiber.Ctx) error {
			if err := clientSocket.EnrollClient("ws://"+config.HigherAddress+":"+port+"/client/enroll", string(c.Body())); err != nil {
				return c.SendString(err.Error())
			}
			return c.SendString("Succeed")
		})

		app.Get("/client", func(c *fiber.Ctx) error {
			clients, _ := book.List()
			return c.JSON(clients)
		})

		app.Get("/config/higheraddress", func(c *fiber.Ctx) error {
			return c.SendString(config.HigherAddress)
		})
		app.Put("/config/higheraddress", func(c *fiber.Ctx) error {
			body := c.Body()
			if len(body) == 0 {
				return c.SendString("Failed")
			}
			config.HigherAddress = string(body)
			fmt.Println(config.HigherAddress)
			return c.SendString("Succeed")
		})
		app.Delete("config/higheraddress", func(c *fiber.Ctx) error {
			config.HigherAddress = "127.0.0.1"
			return c.SendString("Succeed")
		})

		log.Fatal(app.Listen("0.0.0.0:" + port))
	}()

	if *useWeb {
		const debug = true
		w := webview.New(debug)
		defer w.Destroy()
		w.SetSize(800, 640, webview.HintNone)
		w.Navigate("http://127.0.0.1:" + port + "/public/")
		w.Run()
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
