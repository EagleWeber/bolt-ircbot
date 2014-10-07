package boltircbot

import (
	"github.com/thoj/go-ircevent"
	"regexp"
)

func AddAction(ircproj *irc.Connection, hash string, response string) error {
	x := regexp.MustCompile(hash)
	ircproj.AddCallback("PRIVMSG", func(event *irc.Event) {
		matches := x.FindAllStringSubmatch(event.Message(), -1)
		if len(matches) > 0 {
			ircproj.Actionf(event.Arguments[0], response, event.Nick)
		}
	})

	return nil
}
