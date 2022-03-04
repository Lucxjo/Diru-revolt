package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/5elenay/revoltgo"
	"github.com/lucxjo/diru-revolt/cfg"
)

func main() {

	config := cfg.GetConfig()

    // Init a new client.
    client := revoltgo.Client{
        Token: config.RevoltToken,
    }

    // Listen a on message event.
    client.OnMessage(func(m *revoltgo.Message) {
        if m.Content == "!ping" {
            sendMsg := &revoltgo.SendMessage{}
            sendMsg.SetContent("🏓 Pong!")

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
