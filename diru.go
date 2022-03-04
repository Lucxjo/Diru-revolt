package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/5elenay/revoltgo"
	"github.com/lucxjo/diru-revolt/cfg"
)

var dev bool

func initFlags() {
	flag.BoolVar(&dev, "d", false, "Run in development mode")

	flag.Parse()
}

func main() {

	initFlags()
	var config cfg.DiruConfig

	if dev {
		config = cfg.GetConfig("_diru")
	} else {
		config = cfg.GetConfig("diru")
	}

    // Init a new client.
    client := revoltgo.Client{
        Token: config.Revolt.Token,
    }

    // Listen a on message event.
    client.OnMessage(func(m *revoltgo.Message) {
		if m.AuthorId == config.Revolt.Uid {
			return
		}
		if m.Content == "!ping" {
			sendMsg := &revoltgo.SendMessage{}
			sendMsg.SetContent("!pong")
			m.Reply(true, sendMsg)
		}
    })

    // Start the client.
    client.Start()

    // Wait for close.
    sc := make(chan os.Signal, 1)

    signal.Notify(
        sc,
        syscall.SIGINT,
        syscall.SIGTERM,
        os.Interrupt,
    )
    <-sc

    // Destroy client.
    client.Destroy()
}
