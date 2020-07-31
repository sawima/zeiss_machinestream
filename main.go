package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"
)

var port = flag.String("port", ":8090", "local http service port")

func main() {
	flag.Parse()
	go startHTTP()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	c := connectWS()
	defer c.Close()
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				logc := color.New(color.FgYellow)
				logc.Printf("Connection error:%s \n", err)
				_, ok := err.(*websocket.CloseError)
				if ok {
					logc.Println("Websocket Server close the connection,try to reconnect...")
					c = connectWS()
				}
			} else {
				log.Printf("receive data: %s", message)
				updateNewRecord(message)
			}
		}
	}()
	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")
			return
		}
	}
}
