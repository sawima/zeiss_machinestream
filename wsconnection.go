package main

import (
	"log"
	"net/url"

	"github.com/fatih/color"

	"github.com/gorilla/websocket"
)

func connectWS() *websocket.Conn {
	u := url.URL{Scheme: "ws", Host: "machinestream.herokuapp.com", Path: "/ws"}

	logc := color.New(color.FgGreen)
	logc.Printf("connecting to %s \n", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	return c
}
