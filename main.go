package main

import (
	"fmt"
	"github.com/thoj/go-ircevent"
	"strings"
)

func main() {
	fmt.Println("vim-go")
	con := irc.IRC("groot", "groot")
	err := con.Connect("server.com:6667")
	if err != nil {
		fmt.Println("Failed connecting")
		return
	}
	var roomName = "#channel"
	con.AddCallback("001", func(e *irc.Event) {
		con.Join(roomName)
	})
	con.AddCallback("PRIVMSG", func(e *irc.Event) {
		if strings.HasPrefix(e.Message(), "groot") {
			con.Privmsg(roomName, "I am groot")
			return
			if strings.Contains(e.Message(), "i love you") {
				con.Privmsg(roomName, "We are groot")
				return
			}
		}
	})
	con.Loop()
}
